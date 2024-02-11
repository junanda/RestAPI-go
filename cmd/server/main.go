package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	berandaService := services.Init()

	authController := controllers.InitAuthController(userService)
	berandaController := controllers.Init(berandaService)

	authController.Handler(r)
	berandaController.Handler(r)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server.....")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	dbc, _ := db.DB()
	//catching ctx.Done()
	select {
	case <-ctx.Done():
		dbc.Close()
		log.Println("timeout of 5 seconds")
	}
	log.Println("Server exiting")

	// r.Run(":8080")
}
