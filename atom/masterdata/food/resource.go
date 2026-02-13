package atom_masterdata_food

import (
	"errors"
	"log"
)

func GetAllFoodListUseCase() ([]FoodDataResponse, bool, error) {
	data, status, err := GetAllFoodListDB()

	if !status {
		log.Println("[atom][masterdata_food][usecase][GetAllFoodListUseCase] status error - status =", errors.New("status error"))
		return nil, status, nil
	}

	if err != nil {
		log.Println("[atom][masterdata_food][usecase][GetAllFoodListUseCase] error =", err.Error())
		return nil, false, err
	}

	if len(data) == 0 {
		log.Println("[atom][masterdata_food][usecase][GetAllFoodListUseCase] data is empty!")
	}

	return data, true, nil
}

func GetAllActiveFoodListUseCase() ([]FoodDataResponse, bool, error) {
	data, status, err := GetAllActiveFoodListDB()

	if !status {
		log.Println("[atom][masterdata_food][usecase][GetAllActiveFoodListUseCase] status error - status =", errors.New("status error"))
		return nil, status, nil
	}

	if err != nil {
		log.Println("[atom][masterdata_food][usecase][GetAllActiveFoodListUseCase] error =", err.Error())
		return nil, false, err
	}

	if len(data) == 0 {
		log.Println("[atom][masterdata_food][usecase][GetAllActiveFoodListUseCase] data is empty!")
	}

	return data, true, nil
}

func GetFoodListByNameUseCase(foodName string) ([]FoodDataResponse, bool, error) {
	data, status, err := GetFoodListByNameDB(foodName)

	if !status {
		log.Println("[atom][masterdata_food][usecase][GetFoodByNameUseCase] status error - status =", errors.New("status error"))
		return nil, status, nil
	}

	if err != nil {
		log.Println("[atom][masterdata_food][usecase][GetFoodByNameUseCase] error =", err.Error())
		return nil, false, err
	}

	if len(data) == 0 {
		log.Println("[atom][masterdata_food][usecase][GetFoodByNameUseCase] no food found!")
		// Note: Returning true, nil (empty list) is often preferred over an error
		return nil, true, errors.New("no food found")
	}

	return data, true, nil
}

func PostCreateFoodUseCase(inputData PostCreateFoodRequest) (bool, error) {

	status, err := PostCreateFoodDB(inputData)

	if !status {
		log.Println("[atom][masterdata_food][usecase][PostCreateFoodUseCase] status error - status =", errors.New("status error"))
		return status, nil
	}

	if err != nil {
		log.Println("[atom][masterdata_food][usecase][PostCreateFoodUseCase] error =", err.Error())
		return false, err
	}

	return true, nil
}

func PutUpdateFoodUseCase(inputData PutUpdateFoodRequest) (bool, error) {

	status, err := PutUpdateFoodDB(inputData)

	if !status {
		log.Println("[atom][masterdata_food][usecase][PutUpdateFoodUseCase] status error - status =", errors.New("status error"))
		return status, nil
	}

	if err != nil {
		log.Println("[atom][masterdata_food][usecase][PutUpdateFoodUseCase] error =", err.Error())
		return false, err
	}

	return true, nil
}

func PutDeleteFoodUseCase(inputData PutDeleteFoodRequest) (bool, error) {
	status, err := PutDeleteFoodDB(inputData)

	if !status {
		log.Println("[atom][masterdata_food][usecase][PutDeleteFoodUseCase] status error - status =", errors.New("status error"))
		return status, nil
	}

	if err != nil {
		log.Println("[atom][masterdata_food][usecase][PutDeleteFoodUseCase] error =", err.Error())
		return false, err
	}

	return true, nil
}

func GetDetailFoodUseCase(foodId GetDetailFoodRequest) (FoodDataResponse, bool, error) {
	data, status, err := GetDetailFoodDB(foodId)

	if !status {
		log.Println("[atom][masterdata_food][usecase][GetDetailFoodUseCase] status error - status =", errors.New("status error"))
		return data, status, nil
	}

	if err != nil {
		log.Println("[atom][masterdata_food][usecase][GetDetailFoodUseCase] error =", err.Error())
		return data, status, err
	}

	return data, true, nil
}
