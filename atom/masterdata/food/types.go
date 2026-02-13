package atom_masterdata_food

type FoodDataResponse struct {
	Food_ID          string   `json:"food_id"`
	Food_Description *string  `json:"food_description"`
	Food_Name        string   `json:"food_name"`
	Food_Type        *string  `json:"food_type"`
	Brand_Name       *string  `json:"brand_name"`
	Calories         *float64 `json:"calories"`
	Serving_G        *float64 `json:"serving_g"`
	Carbohydrate_G   *float64 `json:"carbohydrate_g"`
	Protein_G        *float64 `json:"protein_g"`
	Fat_G            *float64 `json:"fat_g"`
	Is_Active        *int16   `json:"is_active"`
}

type PostCreateFoodRequest struct {
	Food_Description *string  `json:"food_description"`
	Food_Name        string   `json:"food_name"`
	Food_Type        *string  `json:"food_type"`
	Brand_Name       *string  `json:"brand_name"`
	Calories         *float64 `json:"calories"`
	Serving_G        *float64 `json:"serving_g"`
	Carbohydrate_G   *float64 `json:"carbohydrate_g"`
	Protein_G        *float64 `json:"protein_g"`
	Fat_G            *float64 `json:"fat_g"`
}

type PutUpdateFoodRequest struct {
	Food_ID          string   `json:"food_id"`
	Food_Description *string  `json:"food_description"`
	Food_Name        string   `json:"food_name"`
	Food_Type        *string  `json:"food_type"`
	Brand_Name       *string  `json:"brand_name"`
	Calories         *float64 `json:"calories"`
	Serving_G        *float64 `json:"serving_g"`
	Carbohydrate_G   *float64 `json:"carbohydrate_g"`
	Protein_G        *float64 `json:"protein_g"`
	Fat_G            *float64 `json:"fat_g"`
}

type PutDeleteFoodRequest struct {
	Food_ID string `json:"food_id"`
}

type GetDetailFoodRequest struct {
	Food_ID string `json:"food_id"`
}

type GetFoodByNameRequest struct {
	Query string `json:"query"`
}
