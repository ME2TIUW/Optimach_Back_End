package main

import (
	"optimach_service/config/database"
	"optimach_service/routes"
)

func main() {

	// if err := godotenv.Load(); err != nil {
	// 	log.Println("Info: No .env file found, relying on System Environment Variables.")
	// }

	database.Init()

	defer database.DB.Close()
	router := routes.SetupRoutes()
	router.Run(":8070")
}
