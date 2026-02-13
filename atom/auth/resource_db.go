package atom_auth

import (
	"database/sql"
	"errors"
	"log"
	"optimach_service/config/database"
)

func GetUserDataDB(inputData UserCredentialsRequestModel) (UserDataListResponseModel, bool, error) {

	query := `
	SELECT
		"id_user",
		"username",
		"password",
		"have_filled_form",
		"is_admin",
		"is_active"
	FROM "public"."tbl_users"
	WHERE "username" = $1
	`

	var data UserDataListResponseModel
	err := database.DB.QueryRow(query, inputData.Username).Scan(
		&data.Id_User,
		&data.Username,
		&data.Password,
		&data.Have_Filled_Form,
		&data.Is_Admin,
		&data.Is_Active)

	if err != nil {

		if err == sql.ErrNoRows {
			log.Println("[atom][auth][resource_db] no user found", err.Error())
			return data, true, errors.New("no user found")
		}

		log.Println("[atom][auth][resource_db] error on queryRow", err.Error())
		return data, false, err
	}

	return data, true, nil
}

func PostCreateUserDB(inputData UserCredentialsRequestModel) (bool, error) {

	query := `
	INSERT INTO 
		"public"."tbl_users"("username", "password")
	VALUES ($1, $2)
	`

	_, err := database.DB.Exec(query, inputData.Username, inputData.Password)

	if err != nil {
		log.Println("[atom][auth][resource_db][PostCreateUserDB] query execution failed", err.Error())
		return false, err
	}

	return true, nil
}

func PutUpdateUserPasswordDB(inputData UserCredentialsRequestModel) (bool, error) {

	query := `
	UPDATE 
		"public"."tbl_users"
	SET 
		password = $1
	WHERE username = $2
	`
	_, err := database.DB.Exec(query, inputData.Password, inputData.Username)

	if err != nil {
		log.Println("[atom][auth][resource_db][putUpdateUserPassword] query execution failed", err.Error())
		return false, err
	}

	return true, nil

}

func CheckCreatedUserDB(inputData UsernameInputRequestModel) (bool, error) {

	query := `
		SELECT EXISTS (
			SELECT 1
		FROM 
			"public"."tbl_users"
		WHERE 
			"username" = $1
		)
	`
	var exists bool

	err := database.DB.QueryRow(query, &inputData.Username).Scan(&exists)

	if err != nil {
		log.Println("[atom][auth][resource_db][PostCreateUserDB] query execution failed", err.Error())
		return false, err
	}

	return exists, nil
}

// func IsRefreshTokenStored(token string) (int, bool) {
// 	log.Println("[atom][auth][resource_db][IsRefreshTokenStored] checking refresh token ")

// 	return utils_token.IsRefreshTokenValid(token)
// }

// func StoreNewRefreshToken(token string, id_user int) {
// 	log.Println("[atom][auth][resource_db][IsRefreshTokenStored] storing new refresh token")

// 	utils_token.StoreRefreshToken(token, id_user)
// }
