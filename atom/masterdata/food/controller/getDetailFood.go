package atom_masterdata_food

import (
	"fmt"
	"net/http"
	atom_masterdata_food "optimach_service/atom/masterdata/food"

	"github.com/gin-gonic/gin"
)

func getFloatVal(p *float64) float64 {
	if p == nil {
		return 0.0
	}
	return *p
}

func GetDetailFood(ctx *gin.Context) {
	var inputData atom_masterdata_food.GetDetailFoodRequest

	inputErr := ctx.ShouldBindJSON(&inputData)

	if inputErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"results": nil,
			"message": "Invalid request body: " + inputErr.Error(),
			"status":  400,
		})
		return
	}

	data, status, err := atom_masterdata_food.GetDetailFoodUseCase(inputData)

	if !status {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"results": nil,
			"status":  400,
			"message": err.Error(),
		})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"results": nil,
			"status":  400,
			"message": err.Error(),
		})
		return
	}

	caloriesVal := getFloatVal(data.Calories)
	proteinVal := getFloatVal(data.Protein_G)
	fatVal := getFloatVal(data.Fat_G)
	carbsVal := getFloatVal(data.Carbohydrate_G)
	servingGVal := getFloatVal(data.Serving_G)

	foodID := fmt.Sprintf("%v", data.Food_ID)
	foodName := fmt.Sprintf("%v", data.Food_Name)
	foodType := fmt.Sprintf("%v", data.Food_Type)

	caloriesStr := fmt.Sprintf("%.0f", caloriesVal)
	proteinStr := fmt.Sprintf("%.2f", proteinVal)
	fatStr := fmt.Sprintf("%.2f", fatVal)
	carbsStr := fmt.Sprintf("%.0f", carbsVal)
	servingGStr := fmt.Sprintf("%.3f", servingGVal)

	foodData := gin.H{
		"food_id":   foodID,
		"food_name": foodName,
		"food_type": foodType,
		"food_url":  fmt.Sprintf("https://yourdomain.com/food/%s", foodID), // Simulated URL
		"servings": gin.H{
			"serving": []gin.H{
				{
					// Map DB nutritional fields (converted to string)
					"calories":     caloriesStr,
					"carbohydrate": carbsStr,
					"fat":          fatStr,
					"protein":      proteinStr,

					// Map/Simulate serving size and descriptions
					"metric_serving_amount": servingGStr,
					"metric_serving_unit":   "g",
					// Use the extracted float value for serving description formatting
					"serving_description": fmt.Sprintf("%v g", servingGVal),

					// Simulated fields (based on the required structure but not in your DB)
					"measurement_description": "g",
					"serving_id":              foodID + "_master",
					"serving_url":             fmt.Sprintf("https://yourdomain.com/serving/%s_master", foodID),
				},
			},
		},
		"brand_name":       data.Brand_Name,
		"is_active":        data.Is_Active,
		"food_description": data.Food_Description,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Success get food detail data (FatSecret Format)",
		"results": gin.H{
			"food": foodData,
		},
	})
}
