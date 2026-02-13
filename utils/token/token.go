package utils_token

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// var refreshTokens = make(map[string]int)
// var mu sync.Mutex

// func StoreRefreshToken(token string, id_user int) {
// 	mu.Lock()
// 	refreshTokens[token] = id_user
// 	mu.Unlock()
// 	log.Printf("Stored refresh token for id user %d", id_user)
// }

// func IsRefreshTokenValid(token string) (int, bool) {
// 	mu.Lock()
// 	id_user, ok := refreshTokens[token]
// 	mu.Unlock()

// 	if !ok {
// 		return 0, false
// 	}

// 	return id_user, true
// }

// func InvalidateRefreshToken(token string) {
// 	mu.Lock()
// 	delete(refreshTokens, token)
// 	mu.Unlock()
// 	log.Printf("invalidated refresh token : %s", token)
// }

func GenerateTokens(id_user int) (accessToken string, refreshToken string, accessExp time.Time, refreshExp time.Time, err error) {

	accessExpiration, accErr := strconv.Atoi(os.Getenv("ACCESS_TOKEN_LIFESPAN"))

	if accErr != nil {
		return "", "", time.Time{}, time.Time{}, accErr
	}

	refreshExpiration, refErr := strconv.Atoi(os.Getenv("REFRESH_TOKEN_LIFESPAN"))
	if refErr != nil {
		return "", "", time.Time{}, time.Time{}, refErr
	}

	accessClaims := &AccessClaims{
		Id_User: id_user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(accessExpiration) * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	refreshClaims := &RefreshClaims{
		Id_User: id_user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(refreshExpiration) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	//create access tokens
	accessTokenWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessToken, accessTokenErr := accessTokenWithClaims.SignedString([]byte(os.Getenv("API_ACCESS_SECRET")))

	if accessTokenErr != nil {
		log.Println("[utils][token][generateTokens] access token creation error", err)
	}

	//create refresh tokens
	refreshTokenWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshToken, refreshTokenErr := refreshTokenWithClaims.SignedString([]byte(os.Getenv("API_REFRESH_SECRET")))

	if refreshTokenErr != nil {
		log.Println("[utils][token][generateTokens] access token creation error", err)
	}

	// StoreRefreshToken(refreshToken, id_user)
	return accessToken, refreshToken, accessClaims.ExpiresAt.Time, refreshClaims.ExpiresAt.Time, nil
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")

	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")

	if bearerToken == "" {
		return ""
	}

	parts := strings.Split(bearerToken, " ")

	if len(parts) == 2 && strings.ToLower(parts[0]) == "bearer" {
		return parts[1]
	}

	return ""
}

func ValidateAccessToken(tokenString string) (*AccessClaims, error) {
	claims := &AccessClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])

		}
		return []byte(os.Getenv("API_ACCESS_SECRET")), nil
	})
	if err != nil {

		if errors.Is(err, jwt.ErrTokenExpired) {
			log.Println("[utils][token][ValidateAccessToken] access token expired ", err.Error())
			return nil, err
		}

		log.Println("[utils][token][ValidateAccessToken] token parse error : ", err.Error())
		return nil, err
	}
	if !token.Valid {
		log.Println("invalid access token signature or claims")
		return nil, err
	}
	return claims, nil
}

func ValidateRefreshToken(tokenString string) (*RefreshClaims, error) {
	claims := &RefreshClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])

		}
		return []byte(os.Getenv("API_REFRESH_SECRET")), nil
	})
	if err != nil {

		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, err
		}

		log.Println("[utils][token][ValidateAccessToken] token parse error : ", err.Error())
		return nil, err
	}
	if !token.Valid {
		log.Println("invalid access token signature or claims")
		return nil, err
	}

	//verify that token is stored
	// storedUserId, isStored := IsRefreshTokenValid(tokenString)

	// if !isStored || storedUserId != claims.Id_User {
	// 	log.Printf("[utils][token][ValidateRefreshToken] Server-side check failed for token: %s", tokenString)
	// 	return nil, errors.New("invalid or revoked refresh token")
	// }

	return claims, nil
}
