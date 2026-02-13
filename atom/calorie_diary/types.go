package atom_calorie_diary

type CalorieDiaryResponseList struct {
	Id_Diary           int    `json:"id_diary"`
	Id_User            int    `json:"id_user"`
	Total_Calories_In  int    `json:"total_calories_in"`
	Total_Calories_out int    `json:"total_calories_out"`
	Diary_Date         string `json:"diary_date"`
	Net_Calorie        int    `json:"net_calorie"`
	Is_Active          int    `json:"is_active"`
}

type GetTotalCalorieByDateResponseModel struct {
	Date               string  `json:"date"`
	Id_User            int     `json:"id_user"`
	Total_Calorie_In   int     `json:"total_calorie_in"`
	Total_Calorie_Out  int     `json:"total_calorie_out"`
	Net_Calories       int     `json:"net_calories"`
	Total_Breakfast    int     `json:"total_breakfast"`
	Total_Lunch        int     `json:"total_lunch"`
	Total_Dinner       int     `json:"total_dinner"`
	Total_Snack        int     `json:"total_snack"`
	Total_Protein      float32 `json:"total_protein"`
	Total_Carbohydrate float32 `json:"total_carbohydrate"`
	Total_Fat          float32 `json:"total_fat"`
}

type CreateCalorieDiaryRequestList struct {
	Id_Diary           int `json:"id_diary"`
	Id_User            int `json:"id_user"`
	Total_Calories_In  int `json:"total_calories_in"`
	Total_Calories_out int `json:"total_calories_out"`
	Net_Calorie        int `json:"net_calorie"`
}

type DetailCalorieDiaryRequestList struct {
	Id_Diary int `json:"id_diary"`
}

type GetTotalCalorieByDateRequestModel struct {
	Id_User int    `json:"id_user"`
	Date    string `json:"date"`
}

type DetailCalorieDiaryByIdUserRequestList struct {
	Id_User int `json:"id_user"`
}

type UpdateCalorieDiaryRequestList struct {
	Id_Diary           int `json:"id_diary"`
	Id_User            int `json:"id_user"`
	Total_Calories_In  int `json:"total_calories_in"`
	Total_Calories_out int `json:"total_calories_out"`
	Net_Calorie        int `json:"net_calorie"`
}

// "diary_date" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ,
// "is_active" INTEGER NOT NULL DEFAULT 1 ,
