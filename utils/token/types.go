package utils_token

import 
"github.com/golang-jwt/jwt/v5"

type AccessClaims struct{
	Id_User int `json:"id_user"`
	jwt.RegisteredClaims
}

type RefreshClaims struct {
	Id_User int `json:"id_user"`
	jwt.RegisteredClaims
}