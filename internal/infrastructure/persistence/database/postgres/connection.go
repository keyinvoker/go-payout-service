package postgres

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/keyinvoker/go-payout-service/internal/config"
)

func NewPostgresConnection() (*gorm.DB, error) {
	dbUser := config.GetConfig().DBUser
	dbPassword := config.GetConfig().DBPassword
	dbHost := config.GetConfig().DBHost
	dbPort := config.GetConfig().DBPort
	dbName := config.GetConfig().DBName

	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable TimeZone=UTC",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("NewPostgresConnection :: Failed connecting: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("NewPostgresConnection :: Failed to get underlying sql.DB: %w", err)
	}

	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("NewPostgresConnection :: Failed to ping Postgres: %w", err)
	}

	log.Println("NewPostgresConnection :: Successfully connected to Postgres database")

	return db, nil
}
