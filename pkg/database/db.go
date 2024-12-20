package database

import (
	"auth/config"
	"auth/internal/models"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *Database

type Database struct {
	DB *gorm.DB
}

// Models returns all models that need to be migrated
func Models() []interface{} {
	return []interface{}{
		&models.User{},
		&models.OAuthAccount{},
		&models.RefreshToken{},
		&models.LoginAttempt{},
	}
}

func Init() *Database {
	if DB != nil {
		return DB
	}

	config := config.DefaultDatabaseConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=Asia/Kolkata",
		config.Host, config.User, config.Password, config.Name, config.Port, config.SslMode)

	// Add connection retry logic
	var db *gorm.DB
	var err error
	maxRetries := 5

	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Warn().Msgf("Failed to connect to database, attempt %d of %d", i+1, maxRetries)
		time.Sleep(time.Second * 5)
	}

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database after multiple attempts")
		return nil
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get database instance")
		return nil
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	DB = &Database{DB: db}

	if err := DB.AutoMigrate(); err != nil {
		log.Fatal().Err(err).Msg("Failed to run database migrations")
		return nil
	}

	log.Info().Msg("Database connected and migrations completed successfully")
	return DB
}

func (d *Database) GetDB() *gorm.DB {
	return d.DB
}

func (d *Database) AutoMigrate() error {
	return d.DB.AutoMigrate(Models()...)
}
