package atom_activity_log

import (
	"database/sql"
	"log"
	"optimach_service/config/database"
)

func GetAllActivityLogListDB() ([]ActivityLogListResponseModel, bool, error) {

	query := `
	SELECT
  "id_activities",
	"id_user",
	"id_diary",
	"length_minutes",
	"calories_burned",
	"logged_at",
	"is_active",
	"activities"
FROM
  "public"."tbl_activities_log"
	`

	rows, err := database.DB.Query(query)

	if err != nil {
		log.Println("[atom][activity_log][resource_db][GetAllActivityLogListDB] error on query ", err.Error())
		return nil, false, err
	}
	defer rows.Close()

	var datas []ActivityLogListResponseModel

	for rows.Next() {
		var data ActivityLogListResponseModel
		err = rows.Scan(
			&data.Id_Activities,
			&data.Id_User,
			&data.Id_Diary,
			&data.Length_Minutes,
			&data.Calories_Burned,
			&data.Logged_At,
			&data.Is_Active,
			&data.Activities,
		)

		if err != nil {
			log.Println("[[atom][activity_log][resource_db][GetAllActivityLogListDB] error scanning data ", err.Error())
			return nil, false, err
		}
		datas = append(datas, data)
	}

	return datas, true, nil
}

func GetAllActiveActivityLogListDB() ([]ActivityLogListResponseModel, bool, error) {

	query := `
	SELECT
  	"id_activities",
		"id_user",
		"id_diary",
		"length_minutes",
		"calories_burned",
		"logged_at",
		"is_active",
		"activities"
	FROM
  	"public"."tbl_activities_log"
	WHERE "is_active" = 1
	`

	rows, err := database.DB.Query(query)

	if err != nil {
		log.Println("[atom][activity_log][resource_db][GetAllActiveActivityLogListDB] error on query ", err.Error())
		return nil, false, err
	}
	defer rows.Close()

	var datas []ActivityLogListResponseModel

	for rows.Next() {
		var data ActivityLogListResponseModel
		err = rows.Scan(
			&data.Id_Activities,
			&data.Id_User,
			&data.Id_Diary,
			&data.Length_Minutes,
			&data.Calories_Burned,
			&data.Logged_At,
			&data.Is_Active,
			&data.Activities,
		)

		if err != nil {
			log.Println("[[atom][activity_log][resource_db][GetAllActiveActivityLogListDB] error scanning data ", err.Error())
			return nil, false, err
		}
		datas = append(datas, data)
	}

	return datas, true, nil
}

func PostCreateActivityLogDB(inputData CreateActivityLogListRequestModel) (bool, error) {

	query := `
	INSERT INTO 
	"public"."tbl_activities_log"
	(
		id_user,
		id_diary,
		length_minutes,
		calories_burned,
		activities
	)
	VALUES ($1, $2, $3, $4, $5)
	`

	_, err := database.DB.Exec(query,
		&inputData.Id_User,
		&inputData.Id_Diary,
		&inputData.Length_Minutes,
		&inputData.Calories_Burned,
		&inputData.Activities,
	)

	if err != nil {
		log.Println("[atom][activity_log][resource_db][PostCreateActivityLogDB] error executing insert activity log query ", err.Error())
		return false, err
	}

	return true, nil
}

func PutUpdateActivityLogDB(inputData UpdateActivityLogListRequestModel) (bool, error) {

	query :=
		`UPDATE
		"public"."tbl_activities_log"
	SET
		"id_user" = $1,
		"id_diary" = $2,
		"length_minutes" = $3,
		"calories_burned" = $4,
		"activities" = $5
	WHERE "id_activities" = $6
	`

	_, err := database.DB.Exec(query,
		&inputData.Id_User,
		&inputData.Id_Diary,
		&inputData.Length_Minutes,
		&inputData.Calories_Burned,
		&inputData.Activities,
		&inputData.Id_Activities,
	)

	if err != nil {
		log.Println("[atom][activity_log][resource_db][PutUpdateActivityLogDB] error executing update activity log query", err.Error())
		return false, err
	}

	return true, nil
}

func PutDeleteActivityLogListDB(idActivity int) (bool, error) {

	query := `
	UPDATE
	"public"."tbl_activities_log"
	SET
	"is_active" = 0
	WHERE "id_activities" = $1
	`

	_, err := database.DB.Exec(query, idActivity)

	if err != nil {
		log.Println("[atom][activity_log][resource_db][PutDeleteActivityLogDB] error executing delete activity log query", err.Error())

		return false, err
	}

	return true, nil
}

func GetDetailActivityLogByIdUserDB(idUser int) ([]ActivityLogListResponseModel, bool, error) {

	query := `
	SELECT
  	"id_activities",
		"id_user",
		"id_diary",
		"length_minutes",
		"calories_burned",
		"logged_at",
		"is_active",
		"activities"
	FROM
  	"public"."tbl_activities_log"
	WHERE "id_user" = $1 AND "is_active" = 1
	`

	rows, err := database.DB.Query(query, idUser)

	if err != nil {

		if err == sql.ErrNoRows {
			log.Printf("[atom][activity_log][resource_db][GetDetailActivityLogByIdUserDB] activity log from id_user %d not found", idUser)
			return nil, false, sql.ErrNoRows
		}

		log.Println("[atom][activity_log][resource_db] error : ", err.Error())
		return nil, false, err
	}

	var datas []ActivityLogListResponseModel
	for rows.Next() {
		var data ActivityLogListResponseModel
		err = rows.Scan(
			&data.Id_Activities,
			&data.Id_User,
			&data.Id_Diary,
			&data.Length_Minutes,
			&data.Calories_Burned,
			&data.Logged_At,
			&data.Is_Active,
			&data.Activities,
		)
		if err != nil {
			log.Println("[atom][activity_log][resource_db][GetDetailActivityLogByIdUserDB] scan error : ", err.Error())
			return nil, false, err
		}

		datas = append(datas, data)
	}

	return datas, true, nil
}
