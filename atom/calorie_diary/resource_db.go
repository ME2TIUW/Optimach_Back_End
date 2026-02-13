package atom_calorie_diary

import (
	"database/sql"
	"errors"
	"log"
	"optimach_service/config/database"
)

func GetAllTotalCalorieDateDB() ([]GetTotalCalorieByDateResponseModel, bool, error) {

	query := `
	SELECT 
    date, 
    total_calories_in, 
    total_calories_out, 
    net_calories,
		total_protein,
		total_carbohydrate,
		total_fat
  FROM 
    view_daily_diary
  ORDER BY 
  date DESC;
	`

	rows, err := database.DB.Query(query)

	if err != nil {
		log.Println("[atom][calorie_diary][resource_db][GetAllTotalCalorieDateDB] error executing query ", err.Error())
		return nil, false, err
	}

	defer rows.Close()

	var datas []GetTotalCalorieByDateResponseModel

	for rows.Next() {
		var data GetTotalCalorieByDateResponseModel

		err := rows.Scan(
			&data.Date,
			&data.Total_Calorie_In,
			&data.Total_Calorie_Out,
			&data.Net_Calories,
			&data.Total_Protein,
			&data.Total_Carbohydrate,
			&data.Total_Fat,
		)

		if err != nil {
			log.Println("[atom][calorie_diary][resource_db][GetAllTotalCalorieDateDB] error scanning data: ", err.Error())
			return nil, false, err
		}

		datas = append(datas, data)
	}

	return datas, true, nil

}

func GetTotalCalorieByDateDB(idUser int, date string) (GetTotalCalorieByDateResponseModel, bool, error) {

	query := `
	SELECT 
  date, 
  id_user,
  total_calories_in, 
  total_calories_out, 
  net_calories,
  total_breakfast,
  total_lunch,
  total_dinner,
  total_snack,
  total_protein,
  total_carbohydrate,
  total_fat
  FROM 
    view_daily_diary
	WHERE id_user = $1 AND date = $2::DATE
  ORDER BY 
  date DESC;
	`

	dateOnly := date[:10]
	log.Println("dateOnly:", dateOnly)

	var data GetTotalCalorieByDateResponseModel
	err := database.DB.QueryRow(query, idUser, dateOnly).Scan(
		&data.Date,
		&data.Id_User,
		&data.Total_Calorie_In,
		&data.Total_Calorie_Out,
		&data.Net_Calories,
		&data.Total_Breakfast,
		&data.Total_Lunch,
		&data.Total_Dinner,
		&data.Total_Snack,
		&data.Total_Protein,
		&data.Total_Carbohydrate,
		&data.Total_Fat,
	)

	if err != nil {

		if err == sql.ErrNoRows {
			log.Println("[atom][calorie_diary][resource_db][GetTotalCalorieByDateDB] no user found", err.Error())
			return GetTotalCalorieByDateResponseModel{}, true, errors.New("no calorie data found")
		}

		log.Println("[atom][calorie_diary][resource_db][GetTotalCalorieByDateDB] error on queryRow", err.Error())
		return data, false, err
	}

	return data, true, nil

}
