package atom_activity_log

type ActivityLogListResponseModel struct {
	Id_Activities       int     `json:"id_activities"`
	Id_User             int 		`json:"id_user"`
	Id_Diary            int 		`json:"id_diary"`
	Length_Minutes      int 		`json:"length_minutes"`
	Calories_Burned     int 		`json:"calories_burned"`
	Logged_At           string 	`json:"logged_at"`
	Is_Active           int 		`json:"is_active"`
	Activities          string  `json:"activities"`
}



type CreateActivityLogListRequestModel struct {
	Id_User             int 		`json:"id_user"`
	Id_Diary            int 		`json:"id_diary"`
	Length_Minutes      int 		`json:"length_minutes"`
	Calories_Burned     int 		`json:"calories_burned"`
	Activities          string  `json:"activities"`
}



type UpdateActivityLogListRequestModel struct {
	Id_User             int 		`json:"id_user"`
	Id_Diary            int 		`json:"id_diary"`
	Length_Minutes      int 		`json:"length_minutes"`
	Calories_Burned     int 		`json:"calories_burned"`
	Id_Activities       int     `json:"id_activities"`
	Activities          string  `json:"activities"`
}



type DeleteActivityLogRequestModel struct {
	Id_Activities       int     `json:"id_activities"`
}

type DetailActivityLogResponseModel struct {
	Id_User             int 		`json:"id_user"`
}
