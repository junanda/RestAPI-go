package main

import (
	"github.com/gin-gonic/gin"
	"github.com/junanda/golang-aa/config"
	"github.com/junanda/golang-aa/controllers"
	"github.com/junanda/golang-aa/database"
	"github.com/junanda/golang-aa/repository"
	"github.com/junanda/golang-aa/services"
)

func main() {
	r := gin.Default()

	config := config.Initialize()

	db := database.InitDB(config)

	userRepo := repository.InitUserRepository(db)

	userService := services.InitUserService(userRepo)

	authController := controllers.InitAuthController(userService)

	authController.Handler(r)

	r.Run(":8080")
}
