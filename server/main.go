package main

import (
	"fmt"
	"server/Application"
	"server/Models"
	"server/Routes"
)

func main() {
	app := Application.NewApp()
	// Migrate Project Models
	err := app.DB.AutoMigrate(&Models.User{})
	if err != nil {
		fmt.Println("Error Message for initiate server:", err.Error())
	}
	// Close DB connection
	Application.CloseConnection(app)
	// Init Routes
	Routes.InitRoutes(app)

	// Run the project
	err = app.Gin.Run(":3000")
	if err != nil {
		fmt.Println("Error Message for initiate server:", err.Error())
	}
}
