package atom_masterdata_food

import (
	"database/sql"
	"errors"
	"log"
	"optimach_service/config/database"
)

func GetAllFoodListDB() ([]FoodDataResponse, bool, error) {

	query := `
  SELECT
    "food_id",
		"food_description",
    "food_name",
    "food_type",
    "brand_name",
    "calories",
    "serving_g",
    "carbohydrate_g",
    "protein_g",
    "fat_g",
    "is_active"
  FROM "public"."tbl_master_food"
  ORDER BY "food_name" ASC
  `

	rows, err := database.DB.Query(query)
	if err != nil {
		log.Println("[atom][masterdata_food][resource_db][GetAllFoodListDB] error in executing query : ", err.Error())
		return nil, false, err
	}
	defer rows.Close()

	var datas []FoodDataResponse

	for rows.Next() {
		var data FoodDataResponse

		err := rows.Scan(
			&data.Food_ID,
			&data.Food_Description,
			&data.Food_Name,
			&data.Food_Type,
			&data.Brand_Name,
			&data.Calories,
			&data.Serving_G,
			&data.Carbohydrate_G,
			&data.Protein_G,
			&data.Fat_G,
			&data.Is_Active,
		)

		if err != nil {
			log.Println("[atom][masterdata_food][resource_db][GetAllFoodListDB] error while scanning : ", err.Error())
			return nil, false, err
		}

		datas = append(datas, data)
	}

	return datas, true, nil

}

func GetAllActiveFoodListDB() ([]FoodDataResponse, bool, error) {

	query := `
  SELECT
    "food_id",
		"food_description",
    "food_name",
    "food_type",
    "brand_name",
    "calories",
    "serving_g",
    "carbohydrate_g",
    "protein_g",
    "fat_g",
    "is_active"
  FROM "public"."tbl_master_food"
  WHERE "is_active" = 1
  ORDER BY "food_name" ASC
  `

	rows, err := database.DB.Query(query)
	if err != nil {
		log.Println("[atom][masterdata_food][resource_db][GetAllActiveFoodListDB] error in executing query : ", err.Error())
		return nil, false, err
	}
	defer rows.Close()

	var datas []FoodDataResponse

	for rows.Next() {
		var data FoodDataResponse

		err := rows.Scan(
			&data.Food_ID,
			&data.Food_Description,
			&data.Food_Name,
			&data.Food_Type,
			&data.Brand_Name,
			&data.Calories,
			&data.Serving_G,
			&data.Carbohydrate_G,
			&data.Protein_G,
			&data.Fat_G,
			&data.Is_Active,
		)

		if err != nil {
			log.Println("[atom][masterdata_food][resource_db][GetAllActiveFoodListDB] error while scanning : ", err.Error())
			return nil, false, err
		}

		datas = append(datas, data)
	}

	return datas, true, nil
}

func GetFoodListByNameDB(foodName string) ([]FoodDataResponse, bool, error) {

	query := `
	SELECT
	 "food_id",
		"food_description",
    "food_name",
    "food_type",
    "brand_name",
    "calories",
    "serving_g",
    "carbohydrate_g",
    "protein_g",
    "fat_g",
    "is_active"
  FROM "public"."tbl_master_food"
  WHERE "food_name" ILIKE '%' || $1 || '%'
	`

	rows, err := database.DB.Query(query, foodName)
	if err != nil {
		log.Println("[atom][masterdata_food][resource_db][getFoodByName] error while executing query: ", err.Error())
		return nil, false, err
	}
	defer rows.Close()

	var datas []FoodDataResponse

	for rows.Next() {
		var data FoodDataResponse

		err := rows.Scan(
			&data.Food_ID,
			&data.Food_Description,
			&data.Food_Name,
			&data.Food_Type,
			&data.Brand_Name,
			&data.Calories,
			&data.Serving_G,
			&data.Carbohydrate_G,
			&data.Protein_G,
			&data.Fat_G,
			&data.Is_Active,
		)

		if err != nil {
			log.Println("[atom][masterdata_food][resource_db][getFoodByName] error while scanning: ", err.Error())
			return nil, false, err
		}

		datas = append(datas, data)
	}

	return datas, true, nil
}

func PostCreateFoodDB(inputData PostCreateFoodRequest) (bool, error) {

	query := `
  INSERT INTO 
    "public"."tbl_master_food"(
      "food_description",
      "food_name", 
      "food_type", 
      "brand_name", 
      "calories", 
      "serving_g", 
      "carbohydrate_g", 
      "protein_g", 
      "fat_g"
    )
  VALUES 
    ($1, $2, $3, $4, $5, $6, $7, $8, $9)
  `

	_, err := database.DB.Exec(query,
		inputData.Food_Description,
		inputData.Food_Name,
		inputData.Food_Type,
		inputData.Brand_Name,
		inputData.Calories,
		inputData.Serving_G,
		inputData.Carbohydrate_G,
		inputData.Protein_G,
		inputData.Fat_G,
	)

	if err != nil {
		log.Println("[atom][masterdata_food][resource_db][PostCreateFoodDB] error while executing query: ", err.Error())
		return false, err
	}

	return true, nil
}

func PutUpdateFoodDB(inputData PutUpdateFoodRequest) (bool, error) {

	query := `
  UPDATE 
    "public"."tbl_master_food"
  SET 
		"food_description" = $1,
    "food_name" = $2,
    "food_type" = $3,
    "brand_name" = $4,
    "calories" = $5,
    "serving_g" = $6,
    "carbohydrate_g" = $7,
    "protein_g" = $8,
    "fat_g" = $9
  WHERE "food_id" = $10
  `

	_, err := database.DB.Exec(query,
		inputData.Food_Description,
		inputData.Food_Name,
		inputData.Food_Type,
		inputData.Brand_Name,
		inputData.Calories,
		inputData.Serving_G,
		inputData.Carbohydrate_G,
		inputData.Protein_G,
		inputData.Fat_G,
		inputData.Food_ID,
	)

	if err != nil {
		log.Println("[atom][masterdata_food][resource_db][PutUpdateFoodDB] error while updating query: ", err.Error())
		return false, err
	}

	return true, nil
}

func PutDeleteFoodDB(inputData PutDeleteFoodRequest) (bool, error) {

	query := `
  UPDATE "public"."tbl_master_food"
    SET "is_active" = 0
  WHERE "food_id" = $1
  `

	_, err := database.DB.Exec(query, inputData.Food_ID)

	if err != nil {
		log.Println("[atom][masterdata_food][resource_db][PutDeleteFoodDB] error while executing query: ", err.Error())
		return false, err
	}

	return true, nil
}

func GetDetailFoodDB(foodId GetDetailFoodRequest) (FoodDataResponse, bool, error) {

	query := `
  SELECT 
    "food_id",
    "food_description",
    "food_name",
    "food_type",
    "brand_name",
    "calories",
    "serving_g",
    "carbohydrate_g",
    "protein_g",
    "fat_g",
    "is_active"
  FROM "public"."tbl_master_food"
  WHERE "food_id" = $1
  `

	var data FoodDataResponse
	err := database.DB.QueryRow(query, foodId.Food_ID).Scan(
		&data.Food_ID,
		&data.Food_Description,
		&data.Food_Name,
		&data.Food_Type,
		&data.Brand_Name,
		&data.Calories,
		&data.Serving_G,
		&data.Carbohydrate_G,
		&data.Protein_G,
		&data.Fat_G,
		&data.Is_Active,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("[atom][masterdata_food][resource_db][GetDetailFoodDB] no rows found for food_id: ", foodId.Food_ID)
			return data, true, errors.New("no rows found")
		}

		log.Println("[atom][masterdata_food][resource_db][GetDetailFoodDB] error while executing query: ", err.Error())
		return data, false, err
	}

	return data, true, nil
}
