package atom_auth

import (
	"errors"
	"fmt"
	"log"
	utils_bcrypt "optimach_service/utils/bcrypt"
	utils_token "optimach_service/utils/token"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var ErrAuthenticationFailed = errors.New("authentication failed")
var ErrValidationFailed = errors.New("token validation failed")

func GetAllUserListUseCase(inputData UserCredentialsRequestModel) (UserDataListResponseModel, string, string, time.Time, time.Time, bool, error, error) {

	data, status, err := GetUserDataDB(inputData)

	if !status {
		log.Println("[atom][auth][resource][GetAllUserListUseCase] status = false, err = ", err.Error())

		return data, "", "", time.Time{}, time.Time{}, status, ErrAuthenticationFailed, nil
	}

	if err != nil {
		log.Println("[atom][auth][resource][GetAllUserListUseCase] error = ", err.Error())
		return data, "", "", time.Time{}, time.Time{}, false, err, nil
	}

	errPassword := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(inputData.Password))

	if errPassword != nil {
		log.Println("[atom][auth][resource][errPassword] password didn't match", errPassword.Error())
		return data, "", "", time.Time{}, time.Time{}, true, err, errPassword
	}

	accessToken, refreshToken, accessExp, refreshExp, err := utils_token.GenerateTokens(data.Id_User)

	if err != nil {
		log.Println("[atom][auth][resource][GetAllUserListUseCase]", err)
		return data, "", "", time.Time{}, time.Time{}, false, err, nil
	}

	return data, accessToken, refreshToken, accessExp, refreshExp, true, nil, nil
}

func PostCreateUserUseCase(inputData UserCredentialsRequestModel) (bool, error) {

	var inputtedUsername UsernameInputRequestModel

	inputtedUsername.Username = inputData.Username

	exists, err := CheckCreatedUserDB(inputtedUsername)

	if err != nil {
		log.Println("[atom][auth][resource][PostCreateUserUseCase] error in checking created user DB", err.Error())
		return false, err
	}

	if exists {
		duplicateErr := errors.New("username has already been taken")
		log.Println("[atom][auth][resource][PostCreateUserUseCase] username has already been taken, please choose a different username", duplicateErr.Error())

		return false, duplicateErr
	}

	hashedPassword, hashedErr := utils_bcrypt.EncryptString(inputData.Password)

	if hashedErr != nil {
		log.Println("[atom][auth][resource][PostCreateUserUseCase] error hashing password", hashedErr.Error())
		return false, hashedErr
	}

	inputData.Password = hashedPassword
	status, err := PostCreateUserDB(inputData)

	if !status {
		log.Println("[atom][auth][resource][PostCreateUserUseCase] status error = ", status)
		return status, nil
	}

	if err != nil {
		log.Println("[atom][auth][resource][PostCreateUserUseCase] error = ", err.Error())
		return true, err
	}

	return true, nil
}

func PutUpdatePasswordUseCase(inputData UserCredentialsRequestModel) (bool, error) {
	var inputtedUsername UsernameInputRequestModel

	inputtedUsername.Username = inputData.Username

	exists, err := CheckCreatedUserDB(inputtedUsername)

	if err != nil {
		log.Println("[atom][auth][resource][PostCreateUserUseCase] error in checking created user DB", err.Error())
		return false, err
	}

	if !exists {
		notExistErr := errors.New("username does not exist")
		log.Println("[atom][auth][resource][PostCreateUserUseCase] username does not exist", notExistErr.Error())

		return false, notExistErr
	}

	hashedPassword, hashedErr := utils_bcrypt.EncryptString(inputData.Password)

	if hashedErr != nil {
		log.Println("[atom][auth][resource][PostCreateUserUseCase] error hashing password", hashedErr.Error())
		return false, hashedErr
	}

	inputData.Password = hashedPassword
	status, err := PutUpdateUserPasswordDB(inputData)

	if !status {
		log.Println("[atom][auth][resource][putUpdateUserPasswordUseCase] status error = ", status)
		return status, nil
	}

	if err != nil {
		log.Println("[atom][auth][resource][putUpdateUserPasswordUseCase] error = ", err.Error())
		return true, err
	}

	return true, nil
}

func RefreshAccessTokenUseCase(refreshTokenString string) (newAccessToken string, newAccessExp time.Time, err error) {
	log.Println("[atom][auth][resource_db][IsRefreshTokenStored] attempting to refresh token")

	refreshClaims, validateErr := utils_token.ValidateRefreshToken(refreshTokenString)

	if validateErr != nil {
		log.Println("[atom][auth][resource_db][IsRefreshTokenStored] error in validating refresh token")

		return "", time.Time{}, fmt.Errorf("%w: %w", ErrValidationFailed, validateErr)
	}

	id_user := refreshClaims.Id_User

	newAccessToken, _, newAccessExp, _, generateTokenErr := utils_token.GenerateTokens(id_user)

	if generateTokenErr != nil {
		log.Println("[atom][auth][resource_db][IsRefreshTokenStored] error in generating access token")

		return "", time.Time{}, err
	}

	return newAccessToken, newAccessExp, nil

}
