package atom_user

import "time"

type UserResponseList struct {
	Id_User          int        `json:"id_user"`
	Is_Admin         int        `json:"is_admin"`
	Username         string     `json:"username"`
	Height_Cm        *float64   `json:"height_cm"`
	Weight_Kg        *float64   `json:"weight_kg"`
	Have_Filled_Form int        `json:"have_filled_form"`
	BMI              *float64   `json:"bmi"`
	Gender           *string    `json:"gender"`
	DOB              *string    `json:"dob"`
	Age              *int       `json:"age"`
	Created_Date     time.Time  `json:"created_date"`
	Created_By       string     `json:"created_by"`
	Updated_Date     *time.Time `json:"updated_date"`
	Updated_By       *string    `json:"updated_by"`
	Deleted_Date     *time.Time `json:"deleted_date"`
	Deleted_By       *string    `json:"deleted_by"`
	Is_Active        int        `json:"is_active"`
}

type UpdateUserRequestList struct {
	Id_User          int       `json:"id_user"`
	Height_Cm        float64   `json:"height_cm"`
	Weight_Kg        float64   `json:"weight_kg"`
	Have_Filled_Form int       `json:"have_filled_form"`
	BMI              float64   `json:"bmi"`
	Gender           string    `json:"gender"`
	DOB              string    `json:"dob"`
	Age              int       `json:"age"`
	Updated_Date     time.Time `json:"updated_date"`
	Updated_By       string    `json:"updated_by"`
}

type PostCreateUserRequestList struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type DeleteUserRequestList struct {
	Id_User int `json:"id_user"`
}

type GetDetailUserRequestList struct {
	Id_User int `json:"id_user"`
}
