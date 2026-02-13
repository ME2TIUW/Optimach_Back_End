package atom_user

import (
	"log"
)

func GetAllUserListUseCase() ([]UserResponseList, bool, error) {

	datas, status, err := GetAllUserListDB()

	if !status {
		log.Println("[atom][user][resource_db][GetAllUserListDBUseCase] error : ", err.Error())
		return nil, status, err
	}

	if err != nil {
		log.Println("[atom][user][resource_db][GetAllUserListDBUseCase] error : ", err)

		return nil, false, err
	}

	if len(datas) == 0 {
		log.Println("[atom][user][resource_db][GetAllUserListDBUseCase] no data found ")
		return nil, true, nil
	}

	return datas, status, nil
}

func GetAllActiveUserListUseCase() ([]UserResponseList, bool, error) {

	datas, status, err := GetAllActiveUserListDB()

	if !status {
		log.Println("[atom][user][resource_db][GetAllUserListDBUseCase] error : ", err.Error())
		return nil, status, err
	}

	if err != nil {
		log.Println("[atom][user][resource_db][GetAllUserListDBUseCase] error : ", err)

		return nil, false, err
	}

	if len(datas) == 0 {
		log.Println("[atom][user][resource_db][GetAllUserListDBUseCase] no data found ")
		return nil, true, nil
	}

	return datas, status, nil
}

func PutUpdateUserUseCase(inputData UpdateUserRequestList) (bool, error) {

	status, err := PutUpdateUserDB(inputData)

	if !status {
		log.Println("[atom][user][resource][PutUpdateUserUseCase] status error, status = ", status)
		return status, nil
	}

	if err != nil {
		log.Println("[atom][user][resource][PutUpdateUserUseCase] error = ", err.Error())
		return false, err
	}

	return true, nil
}

func PutDeleteUserUseCase(inputData DeleteUserRequestList) (bool, error) {

	status, err := PutDeleteUserDB(inputData)

	if !status {
		log.Println("[atom][user][resource_db][putDeleteUserUseCase] status error - status =  ", status)
		return status, nil
	}

	if err != nil {
		log.Println("[atom][user][resource_db][GetAllUserListDBUseCase] error : ", err)

		return false, err
	}

	return true, nil
}

func GetDetailUserUseCase(userId int) (UserResponseList, bool, error) {

	data, status, err := GetDetailUser(userId)

	if !status {
		log.Println("[atom][user][resource_db][getDetailUserUseCase] status error - status =  ", status)
		return data, status, nil
	}

	if err != nil {
		log.Println("[atom][user][resource_db][getDetailUserUseCase] error : ", err.Error())
		return data, false, err
	}

	return data, true, nil
}
