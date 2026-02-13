package atom_activity_log

import (
	"errors"
	"log"
)


func GetAllActivityLogListUseCase()([]ActivityLogListResponseModel, bool, error){
	data, status, err := GetAllActivityLogListDB()

	if !status {
		log.Println("[atom][activity_log][resource][GetAllActivityLogListUseCase]  status error - status =", errors.New("status error"))

		return nil, status , nil
	 }

	 if err != nil {
		log.Println("[atom][activity_log][resource][GetAllActivityLogListUseCase]  error =", err.Error())
		return nil, false, err
	 }

	 if len(data) == 0 {
		log.Println("[atom][activity_log][resource][GetAllActivityLogListUseCase] data is empty!")
		return data, true, errors.New("data is empty")
	 }

	 return data, true, nil
}

func GetAllActiveActivityLogListUseCase()([]ActivityLogListResponseModel, bool, error){
	data, status, err := GetAllActiveActivityLogListDB()

	if !status {
		log.Println("[atom][activity_log][resource][GetAllActiveActivityLogListUseCase]  status error - status =", errors.New("status error"))

		return nil, status , nil
	 }

	 if err != nil {
		log.Println("[atom][activity_log][resource][GetAllActiveActivityLogListUseCase]  error =", err.Error())
		return nil, false, err
	 }

	 if len(data) == 0 {
		log.Println("[atom][activity_log][resource][GetAllActiveActivityLogListUseCase] data is empty!")
		return data, true, errors.New("data is empty")
	 }

	 return data, true, nil
}

func PostCreateActivityLogUseCase(inputData CreateActivityLogListRequestModel) (bool, error) {

	status, err := PostCreateActivityLogDB(inputData)

	if !status {
		log.Println("[atom][activity_log][resource][PostCreateActivityLogUseCase] status error - status =", errors.New("status error"))

		return status , nil
	 }

	 if err != nil {
		log.Println("[atom][activity_log][resource][PostCreateActivityLogUseCase] error =", err.Error())
		return false, err
	 }

	 return true, nil
}

func PutUpdateActivityLogUseCase(inputData UpdateActivityLogListRequestModel) (bool, error) {

	status, err := PutUpdateActivityLogDB(inputData)

	if !status {
		log.Println("[atom][activity_log][resource][PutUpdateActivityLogUseCase] status error - status =", errors.New("status error"))

		return status , nil
	 }

	 if err != nil {
		log.Println("[atom][activity_log][resource][PutUpdateActivityLogUseCase] error =", err.Error())
		return false, err
	 }

	 return true, nil
}

func PutDeleteActivityLogUseCase(idActivty int) (bool, error) {

	status, err := PutDeleteActivityLogListDB(idActivty)

	if !status {
		log.Println("[atom][activity_log][resource][PutDeleteActivityLogUseCase] status error - status =", errors.New("status error"))

		return status , nil
	 }

	 if err != nil {
		log.Println("[atom][activity_log][resource][PutDeleteActivityLogUseCase] error =", err.Error())
		return false, err
	 }

	 //TO DO
	 //add sql err no rows error from detail

	 return true, nil
}

func GetDetailActivityLogByIdUserUseCase (idUser int) ([]ActivityLogListResponseModel, bool, error){
	
	data, status, err := GetDetailActivityLogByIdUserDB(idUser)

	if !status {
		log.Println("[atom][activity_log][resource][GetDetailActivityLogByIdUserUseCase] status error - status =", errors.New("status error"))

		return nil, status , nil
	 }

	 if err != nil {
		log.Println("[atom][activity_log][resource][GetDetailActivityLogByIdUserUseCase] error =", err.Error())
		return nil, false, err
	 }

	 if len(data) == 0{
		log.Println("[atom][activity_log][resource][GetDetailActivityLogByIdUserUseCase]  data is empty!")
		return data, true, errors.New("data is empty")
	 }

	 return data, true, nil
}