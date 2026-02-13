package atom_food_log

import (
	"errors"
	"log"
)

func GetFoodListUseCase() ([]FoodLogResponseModel, bool, error) {
	data, status, err := GetFoodLogListDB()

	if !status {
		log.Println("[atom][food_log][resource][GetActiveFoodLogListUseCase]  status error - status =", errors.New("status error"))

		return nil, status, nil
	}

	if err != nil {
		log.Println("[atom][food_log][resource][GetActiveFoodLogListUseCase]  error =", err.Error())
		return nil, false, err
	}

	if len(data) == 0 {
		log.Println("[atom][food_log][resource][GetActiveFoodLogListUseCase]  data is empty!")
		return data, true, errors.New("data is empty")
	}

	return data, true, nil
}

func GetActiveFoodListUseCase() ([]FoodLogResponseModel, bool, error) {
	data, status, err := GetActiveFoodLogListDB()

	if !status {
		log.Println("[atom][food_log][resource][GetActiveFoodLogListUseCase] status error - status =", errors.New("status error"))

		return nil, status, nil
	}

	if err != nil {
		log.Println("[atom][food_log][resource][GetActiveFoodLogListUseCase] error =", err.Error())
		return nil, false, err
	}

	if len(data) == 0 {
		log.Println("[atom][food_log][resource][GetActiveFoodLogListUseCase] data is empty!")
		return data, true, errors.New("data is empty")
	}

	return data, true, nil
}

func PostCreateFoodLogUseCase(inputData CreateFoodLogRequestModel) (bool, error) {

	status, err := PostCreateFoodLogDB(inputData)

	if !status {
		log.Println("[atom][food_log][resource][PostCreateFoodLogUseCase] status error - status =", errors.New("status error"))

		return status, nil
	}

	if err != nil {
		log.Println("[atom][food_log][resource][PostCreateFoodLogUseCase] error =", err.Error())
		return false, err
	}

	return true, nil
}

func PutUpdateFoodLogUseCase(inputData UpdateFoodLogRequestModel) (bool, error) {

	status, err := PutUpdateFoodLogListDB(inputData)

	if !status {
		log.Println("[atom][food_log][resource][PutUpdateFoodLogUseCase] status error - status =", errors.New("status error"))

		return status, nil
	}

	if err != nil {
		log.Println("[atom][food_log][resource][PutUpdateFoodLogUseCase] error =", err.Error())
		return false, err
	}

	return true, nil
}

func PutDeleteFoodLogUseCase(idFood int) (bool, error) {
	status, err := PutDeleteFoodLogListDB(idFood)

	if !status {
		log.Println("[atom][food_log][resource][PutDeleteFoodLogUseCase] status error - status =", errors.New("status error"))

		return status, nil
	}

	if err != nil {
		log.Println("[atom][food_log][resource][PutDeleteFoodLogUseCase] error =", err.Error())
		return false, err
	}
	return true, nil
}

func GetDetailFoodLogByIdUserUseCase(idUser int, createdDate string, timezone string) ([]FoodLogResponseModel, bool, error) {
	data, status, err := GetDetailFoodLogByIdUserDB(idUser, createdDate, timezone)

	if !status {
		log.Println("[atom][food_log][resource][GetDetailFoodLogByIdUserDBUseCase] status error - status =", errors.New("status error"))

		return nil, status, nil
	}

	if err != nil {
		log.Println("[atom][food_log][resource][GetDetailFoodLogByIdUserDBUseCase] error =", err.Error())
		return nil, false, err
	}

	if len(data) == 0 {
		log.Println("[atom][food_log][resource][GetDetailFoodLogByIdUserDBUseCase]  data is empty!")
		return data, true, errors.New("data is empty")
	}

	return data, true, nil
}
