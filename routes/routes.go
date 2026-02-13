package routes

import (
	atom_activity_log "optimach_service/atom/activity_log/controller"
	atom_auth "optimach_service/atom/auth/controller"
	atom_calorie_diary "optimach_service/atom/calorie_diary/controller"
	atom_fatSecret "optimach_service/atom/fatSecret/controller"
	atom_food_log "optimach_service/atom/food_log/controller"
	atom_handlers "optimach_service/atom/handlers"
	atom_masterdata_food "optimach_service/atom/masterdata/food/controller"
	atom_user "optimach_service/atom/user/controller"
	middlewares "optimach_service/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	route := gin.Default()

	route.Use(cors.New(cors.Config{

		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:3001", "https://optimach-fe.vercel.app"},
		AllowMethods:     []string{"POST", "PUT", "PATCH", "DELETE", "GET", "OPTIONS", "TRACE", "CONNECT"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With", "Access-Control-Allow-Credentials", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Content-Length"},
	}))

	route.GET("/kaithhealthcheck", atom_handlers.SimpleHealthCheck)
	route.GET("/kaithheathcheck", atom_handlers.SimpleHealthCheck)
	auth := route.Group("/auth")
	{
		auth.POST("login", atom_auth.PostLoginUser)
		auth.PUT("change-password", atom_auth.PutUpdateUserPassword)
		auth.POST("logout", atom_auth.PostLogoutUser)
		auth.POST("register", atom_auth.PostCreateUser)
		auth.POST("refresh", atom_auth.RefreshToken)
	}

	fatSecret := route.Group("/fatsecret", middlewares.JwtAuthenticateMiddleware())
	{
		fatSecret.GET("food-search", atom_fatSecret.GetFoodSearch)
		fatSecret.GET("food-by-id", atom_fatSecret.GetFoodById)

	}

	masterdata_food := route.Group("/masterdata/food", middlewares.JwtAuthenticateMiddleware())
	{
		masterdata_food.GET("list", atom_masterdata_food.GetAllFoodList)
		masterdata_food.GET("list-active", atom_masterdata_food.GetAllActiveFoodList)
		masterdata_food.GET("list-search", atom_masterdata_food.GetFoodListByName)
		masterdata_food.POST("create", atom_masterdata_food.PostCreateFood)
		masterdata_food.PUT("update", atom_masterdata_food.PutUpdateFood)
		masterdata_food.PUT("delete", atom_masterdata_food.PutDeleteFood)
		masterdata_food.POST("detail", atom_masterdata_food.GetDetailFood)
	}

	activitylog := route.Group("/activitylog", middlewares.JwtAuthenticateMiddleware())
	{
		activitylog.GET("list", atom_activity_log.GetAllActivityLogList)
		activitylog.GET("list-active", atom_activity_log.GetAllActiveActivityLogList)
		activitylog.POST("create", atom_activity_log.PostCreateActivityLog)
		activitylog.PUT("update", atom_activity_log.PutUpdateActivityLogList)
		activitylog.PUT("delete", atom_activity_log.PutDeleteActivityLog)
		activitylog.POST("detail", atom_activity_log.PostGetDetailActivityLogListByIdUser)
	}

	foodlog := route.Group("/foodlog", middlewares.JwtAuthenticateMiddleware())
	{
		foodlog.GET("list", atom_food_log.GetFoodList)
		foodlog.GET("list-active", atom_food_log.GetActiveFoodList)
		foodlog.POST("create", atom_food_log.PostCreateFoodLog)
		foodlog.PUT("update", atom_food_log.PutUpdateFoodLog)
		foodlog.DELETE("delete", atom_food_log.PutDeleteFoodLog)
		foodlog.POST("detail", atom_food_log.GetDEtailFoodLogListByIdUser)

	}

	diary := route.Group("/diary", middlewares.JwtAuthenticateMiddleware())
	{
		diary.GET("all-total-by-date", atom_calorie_diary.GetAllTotalCalorieDate)
		diary.POST("total-by-date", atom_calorie_diary.GetTotalCalorieDate)
	}

	user := route.Group("/user", middlewares.JwtAuthenticateMiddleware())
	{
		user.GET("list", atom_user.GetAllUserList)
		user.GET("list-active", atom_user.GetAllActiveUserList)
		user.PUT("update", atom_user.PutUpdateUser)
		user.PUT("delete", atom_user.PutDeleteUserList)
		user.POST("detail", atom_user.PostGetDetailUser)
	}

	return route
}
