package atom_food_log

import (
	"database/sql"
	"errors"
	"log"
	"optimach_service/config/database"
)

func GetFoodLogListDB() ([]FoodLogResponseModel, bool, error) {

	query := `
	SELECT
  "id_food_log",
  "id_user",
  "food_name",
  "protein_grams",
  "carbohydrate_grams",
  "fat_grams",
  "weight_grams",
  "created_date",
  "is_active",
  "calories",
	"occasion"
FROM
  "public"."tbl_food_log"
	`

	rows, err := database.DB.Query(query)

	if err != nil {
		log.Println("[atom][food_log][resource_db][GetFoodLogListDB] error on query ", err.Error())
		return nil, false, err
	}
	defer rows.Close()

	var datas []FoodLogResponseModel

	for rows.Next() {
		var data FoodLogResponseModel
		err = rows.Scan(
			&data.Id_Food_Log,
			&data.Id_User,
			&data.Food_Name,
			&data.Protein_Grams,
			&data.Carbohydrate_Grams,
			&data.Fat_Grams,
			&data.Weight_Grams,
			&data.Created_Date,
			&data.Is_Active,
			&data.Calories,
			&data.Occasion,
		)

		if err != nil {
			log.Println("[[atom][food_log][resource_db][GetFoodLogListDB] error scanning data ", err.Error())
			return nil, false, err
		}
		datas = append(datas, data)
	}

	return datas, true, nil
}

func GetActiveFoodLogListDB() ([]FoodLogResponseModel, bool, error) {

	query := `
	SELECT
  "id_food_log",
  "id_user",
  "food_name",
  "protein_grams",
  "carbohydrate_grams",
  "fat_grams",
  "weight_grams",
  "created_date",
  "is_active",
  "calories",
	"occasion"
FROM
  "public"."tbl_food_log"
WHERE "is_active" = 1
	`

	rows, err := database.DB.Query(query)

	if err != nil {
		log.Println("[atom][food_log][resource_db][GetActiveFoodLogListDB] error on query ", err.Error())
		return nil, false, err
	}

	defer rows.Close()
	var datas []FoodLogResponseModel

	for rows.Next() {
		var data FoodLogResponseModel
		err = rows.Scan(
			&data.Id_Food_Log,
			&data.Id_User,
			&data.Food_Name,
			&data.Protein_Grams,
			&data.Carbohydrate_Grams,
			&data.Fat_Grams,
			&data.Weight_Grams,
			&data.Created_Date,
			&data.Is_Active,
			&data.Calories,
			&data.Occasion,
		)

		if err != nil {
			log.Println("[atom][food_log][resource_db][GetActiveFoodLogListDB] error scanning data ", err.Error())
			return nil, false, err
		}
		datas = append(datas, data)
	}

	return datas, true, nil
}

func PostCreateFoodLogDB(inputData CreateFoodLogRequestModel) (bool, error) {

	query := `
	INSERT INTO 
	"public"."tbl_food_log"
	(
	"id_user",
	"food_name",
	"protein_grams",
	"carbohydrate_grams",
	"fat_grams",
	"weight_grams",
	"calories",
	"occasion"
	)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	_, err := database.DB.Exec(query,
		&inputData.Id_User,
		&inputData.Food_Name,
		&inputData.Protein_Grams,
		&inputData.Carbohydrate_Grams,
		&inputData.Fat_Grams,
		&inputData.Weight_Grams,
		&inputData.Calories,
		&inputData.Occasion,
	)

	if err != nil {
		log.Println("[atom][food_log][resource_db][PostCreateFoodLogDB] error executing insert food log query ", err.Error())
		return false, err
	}

	return true, nil
}

func PutUpdateFoodLogListDB(inputData UpdateFoodLogRequestModel) (bool, error) {

	query :=
		`UPDATE
		"public"."tbl_food_log"
	SET
		"food_name" = $1,
		"protein_grams" = $2,
		"carbohydrate_grams" = $3,
		"fat_grams" = $4,
		"weight_grams" = $5,
		"calories" = $6,
		"occasion" = $7
	WHERE "id_food_log" = $8
	`

	_, err := database.DB.Exec(query,
		inputData.Food_Name,
		inputData.Protein_Grams,
		inputData.Carbohydrate_Grams,
		inputData.Fat_Grams,
		inputData.Weight_Grams,
		inputData.Calories,
		inputData.Occasion,
		inputData.Id_Food_Log,
	)

	if err != nil {
		log.Println("[atom][food_log][resource_db][PutUpdateFoodListDB] error executing update food log query", err.Error())
		return false, err
	}

	return true, nil
}

func PutDeleteFoodLogListDB(idFood int) (bool, error) {

	query := `
	DELETE 
	FROM 
		"public"."tbl_food_log"
  WHERE 
		"id_food_log" = $1
	`

	_, err := database.DB.Exec(query, idFood)

	if err != nil {
		log.Println("[atom][food_log][resource_db][PutDeleteFoodLogListDB] error executing delete food log query", err.Error())

		return false, err
	}

	return true, nil
}

func GetDetailFoodLogByIdUserDB(idUser int, createdDate string, timezone string) ([]FoodLogResponseModel, bool, error) {

	query := `
	SELECT
  	"id_food_log",
  	"id_user",
  	"food_name",
  	"protein_grams",
  	"carbohydrate_grams",
  	"fat_grams",
  	"weight_grams",
  	"created_date",
  	"is_active",
  	"calories",
		"occasion"
	FROM
  	"public"."tbl_food_log"
	WHERE "id_user" = $1 
	AND 
	("created_date" AT TIME ZONE 'UTC' AT TIME ZONE $3)::DATE = $2::DATE AND 
  "is_active" = 1
	`

	date_only := createdDate[:10]

	if len(createdDate) < 10 {
		return nil, false, errors.New("invalid date format")
	}

	if timezone == "" {
		timezone = "Asia/Jakarta"
	}

	rows, err := database.DB.Query(query, idUser, date_only, timezone)

	if err != nil {

		if err == sql.ErrNoRows {
			log.Printf("[atom][food_log][resource_db][GetDetailFoodLogByIdUser] calorie diary from id_user %d not found", idUser)
			return nil, false, sql.ErrNoRows
		}

		log.Println("[atom][food_log][resource_db] error : ", err.Error())
		return nil, false, err
	}

	var datas []FoodLogResponseModel
	for rows.Next() {
		var data FoodLogResponseModel
		err = rows.Scan(
			&data.Id_Food_Log,
			&data.Id_User,
			&data.Food_Name,
			&data.Protein_Grams,
			&data.Carbohydrate_Grams,
			&data.Fat_Grams,
			&data.Weight_Grams,
			&data.Created_Date,
			&data.Is_Active,
			&data.Calories,
			&data.Occasion,
		)
		if err != nil {
			log.Println("[atom][food_log][resource_db][GetDetailFoodLogByIdUser] scan error : ", err.Error())
			return nil, false, err
		}

		datas = append(datas, data)
	}

	return datas, true, nil
}
