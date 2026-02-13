package atom_food_log

type FoodLogResponseModel struct {
	Id_Food_Log        int     `json:"id_food_log"`
	Id_User            int     `json:"id_user"`
	Food_Name          string  `json:"food_name"`
	Protein_Grams      float64 `json:"protein_grams"`
	Carbohydrate_Grams float64 `json:"carbohydrate_grams"`
	Fat_Grams          float64 `json:"fat_grams"`
	Weight_Grams       float64 `json:"weight_grams"`
	Created_Date       string  `json:"created_date"`
	Is_Active          int     `json:"is_active"`
	Calories           int     `json:"calories"`
	Occasion           string  `json:"occasion"`
}

type CreateFoodLogRequestModel struct {
	Id_User            int     `json:"id_user"`
	Food_Name          string  `json:"food_name"`
	Protein_Grams      float64 `json:"protein_grams"`
	Carbohydrate_Grams float64 `json:"carbohydrate_grams"`
	Fat_Grams          float64 `json:"fat_grams"`
	Weight_Grams       float64 `json:"weight_grams"`
	Calories           int     `json:"calories"`
	Occasion           string  `json:"occasion"`
}

type UpdateFoodLogRequestModel struct {
	Food_Name          string  `json:"food_name"`
	Protein_Grams      float64 `json:"protein_grams"`
	Carbohydrate_Grams float64 `json:"carbohydrate_grams"`
	Fat_Grams          float64 `json:"fat_grams"`
	Weight_Grams       float64 `json:"weight_grams"`
	Calories           int     `json:"calories"`
	Occasion           string  `json:"occasion"`
	Id_Food_Log        int     `json:"id_food_log"`
}

type DeleteFoodLogRequestList struct {
	Id_Food_Log int `json:"id_food_log"`
}

type DetailFoodLogRequestList struct {
	Id_User      int    `json:"id_user"`
	Created_Date string `json:"created_date"`
	Timezone     string `json:"timezone"`
}
