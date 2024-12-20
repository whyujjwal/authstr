package main

import (
	"auth/config"
	"auth/pkg/database"
	"auth/pkg/logger"
	"fmt"
	"net/http"
	"time"

	v1 "auth/routes/v1"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

func main() {
	logger.InitializeDistributedLoggers()

	// Initialize database
	db := database.Init()
	if db == nil {
		log.Fatal().Msg("Failed to initialize database")
		return
	}

	gormDB := db.GetDB()
	if gormDB == nil {
		log.Fatal().Msg("Failed to get database instance")
		return
	}

	// Setup router
	router := mux.NewRouter()
	v1.SetupRoutes(router, gormDB)

	// Start server
	port := fmt.Sprintf(":%d", config.DefaultServerConfig().Port)
	log.Info().Msgf("Server starting on port %s", port)

	server := &http.Server{
		Addr:         port,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal().Err(err).Msg("Server failed to start")
	}
}
