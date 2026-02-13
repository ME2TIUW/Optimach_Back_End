package atom_calorie_diary

import (
	"errors"
	"log"
)

func GetAllTotalCalorieDateUseCase() ([]GetTotalCalorieByDateResponseModel, bool, error) {

	datas, status, err := GetAllTotalCalorieDateDB()
	if !status {
		log.Println("[atom][calorie_diary][resource][GetAllTotalCalorieDateUseCase] status error - status = ", status)
		return nil, status, nil
	}

	if err != nil {
		log.Println("[atom][calorie_diary][resource][GetAllTotalCalorieDateUseCase] error  = ", err.Error())
		return nil, false, err
	}

	if len(datas) == 0 {
		log.Println("[atom][calorie_diary][resource][GetAllTotalCalorieDateUseCase] empty calorie diary data")

		return datas, true, errors.New("no calorie diary data in database")
	}

	return datas, true, nil
}

func GetTotalCalorieByDateUseCase(id_user int, date string) (GetTotalCalorieByDateResponseModel, bool, error) {

	data, status, err := GetTotalCalorieByDateDB(id_user, date)
	if !status {
		log.Println("[atom][calorie_diary][resource][GetTotalCalorieByIdDateUseCase] status error - status = ", status)
		return GetTotalCalorieByDateResponseModel{}, status, nil
	}

	if err != nil {
		log.Println("[atom][calorie_diary][resource][GetTotalCalorieByIdDateUseCase] error  = ", err.Error())
		return GetTotalCalorieByDateResponseModel{}, status, err
	}

	return data, true, nil
}
