package v1

import (
	"auth/internal/handlers"
	"auth/internal/repositories"
	"auth/internal/services"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func SetupRoutes(router *mux.Router, db *gorm.DB) {
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	router.HandleFunc("/v1/users", userHandler.CreateUser).Methods("POST")
}
