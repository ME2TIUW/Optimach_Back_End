package atom_user

import (
	"log"
	"optimach_service/config/database"
)

func GetAllUserListDB() ([]UserResponseList, bool, error) {

	query := `
	SELECT
		"id_user",
		"is_admin",
		"username",
		"password",
		"height_cm",
		"weight_kg",
		"bmi",
		"gender",
		"dob",
		"age",
		"created_date",
		"created_by",
		"updated_date",
		"updated_by",
		"deleted_date",
		"deleted_by",
		"is_active"
	FROM "public"."tbl_users"
	`

	rows, err := database.DB.Query(query)
	if err != nil {
		log.Println("[atom][user][resource_db][GetAllUserListDB] error in executing query : ", err.Error())

		return nil, false, err
	}
	defer rows.Close()

	var datas []UserResponseList

	for rows.Next() {
		var data UserResponseList

		err := rows.Scan(
			&data.Id_User,
			&data.Is_Admin,
			&data.Username,
			&data.Height_Cm,
			&data.Weight_Kg,
			&data.BMI,
			&data.Gender,
			&data.DOB,
			&data.Age,
			&data.Created_Date,
			&data.Created_By,
			&data.Updated_Date,
			&data.Updated_By,
			&data.Deleted_Date,
			&data.Deleted_By,
			&data.Is_Active,
		)

		if err != nil {
			log.Println("[atom][user][resource_db][GetAllUserListDB] error while scanning : ", err.Error())
			return nil, false, err
		}

		datas = append(datas, data)
	}

	return datas, true, nil

}

func GetAllActiveUserListDB() ([]UserResponseList, bool, error) {

	query := `
	SELECT
		"id_user",
		"is_admin",
		"username",
		"password",
		"height_cm",
		"weight_kg",
		"have_filled_form",
		"bmi",
		"gender",
		"dob",
		"age",
		"created_date",
		"created_by",
		"updated_date",
		"updated_by",
		"deleted_date",
		"deleted_by",
		"is_active"
	FROM "public"."tbl_users"
	WHERE "is_active" = 1
	`

	rows, err := database.DB.Query(query)
	if err != nil {
		log.Println("[atom][user][resource_db][GetAllUserListDB] error in executing query : ", err.Error())

		return nil, false, err
	}
	defer rows.Close()

	var datas []UserResponseList

	for rows.Next() {
		var data UserResponseList

		err := rows.Scan(
			&data.Id_User,
			&data.Is_Admin,
			&data.Username,
			&data.Height_Cm,
			&data.Weight_Kg,
			&data.Have_Filled_Form,
			&data.BMI,
			&data.Gender,
			&data.DOB,
			&data.Age,
			&data.Created_Date,
			&data.Created_By,
			&data.Updated_Date,
			&data.Updated_By,
			&data.Deleted_Date,
			&data.Deleted_By,
			&data.Is_Active,
		)

		if err != nil {
			log.Println("[atom][user][resource_db][GetAllActiveUserListDB] error while scanning : ", err.Error())
			return nil, false, err
		}

		datas = append(datas, data)
	}

	return datas, true, nil
}

func PostCreateUserDB(inputData PostCreateUserRequestList) (bool, error) {

	query := `
	INSERT INTO 
		"public"."tbl_users"("username", "password")
	VALUES 
		($1, $2)
	`

	_, err := database.DB.Exec(query, inputData)

	if err != nil {
		log.Println("[atom][user][resource_db][PostCreateUserDB] error while executing query: ", err.Error())
	}

	return true, nil
}

func PutUpdateUserDB(inputData UpdateUserRequestList) (bool, error) {

	query := `
	UPDATE 
		"public"."tbl_users"
	SET 
		"height_cm" = $1,
		"weight_kg" = $2,
		"have_filled_form" = $3,
		"bmi" = $4,
		"gender" = $5,
		"dob" = $6,
		"age" = $7,
		"updated_date"=CURRENT_TIMESTAMP,
		"updated_by"=$8
	WHERE "id_user"=$9
	`

	_, err := database.DB.Exec(query,
		inputData.Height_Cm,
		inputData.Weight_Kg,
		inputData.Have_Filled_Form,
		inputData.BMI,
		inputData.Gender,
		inputData.DOB,
		inputData.Age,
		inputData.Updated_By,
		inputData.Id_User,
	)

	if err != nil {
		log.Println("[atom][user][resource_db][PutUpdateUserDB] error while updating query: ", err.Error())
		return false, err
	}

	return true, nil
}

func PutDeleteUserDB(inputData DeleteUserRequestList) (bool, error) {

	query := `
	UPDATE "public"."tbl_users"
	 	SET "is_active" = 0
	WHERE "id_user" = $1
	`

	_, err := database.DB.Exec(query, inputData.Id_User)

	if err != nil {
		log.Println("[atom][user][resource_db][putDeleteUserDB] error while executing query: ", err.Error())
		return false, err
	}

	return true, nil
}

func GetDetailUser(userId int) (UserResponseList, bool, error) {

	query := `
	SELECT 
		"id_user",
		"is_admin",
		"username",
		"height_cm",
		"weight_kg",
		"bmi",
		"gender",
		"dob",
		"age",
		"created_date",
		"created_by",
		"updated_date",
		"updated_by",
		"deleted_date",
		"deleted_by",
		"is_active"
	FROM "public"."tbl_users"
	WHERE "id_user" = $1
	`

	var data UserResponseList
	err := database.DB.QueryRow(query, userId).Scan(
		&data.Id_User,
		&data.Is_Admin,
		&data.Username,
		&data.Height_Cm,
		&data.Weight_Kg,
		&data.BMI,
		&data.Gender,
		&data.DOB,
		&data.Age,
		&data.Created_Date,
		&data.Created_By,
		&data.Updated_Date,
		&data.Updated_By,
		&data.Deleted_Date,
		&data.Deleted_By,
		&data.Is_Active,
	)

	if err != nil {
		log.Println("[atom][user][resource_db][getDetailUserDB] error while executing query: ", err.Error())
		return data, false, err
	}

	return data, true, nil
}
