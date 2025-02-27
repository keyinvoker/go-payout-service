package postgres

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/keyinvoker/go-payout-service/internal/domain/models"
)

func Connect() (*gorm.DB, error) {
	dbUser := os.Getenv("POSTGRES_DB_USER")
	dbPassword := os.Getenv("POSTGRES_DB_PASSWORD")
	dbHost := os.Getenv("POSTGRES_DB_HOST")
	dbPort := os.Getenv("POSTGRES_DB_PORT")
	dbName := os.Getenv("POSTGRES_DB_NAME")

	fmt.Println("POSTGRES_DB_USER:", dbUser)
	fmt.Println("POSTGRES_DB_PASSWORD:", dbPassword)
	fmt.Println("POSTGRES_DB_HOST:", dbHost)
	fmt.Println("POSTGRES_DB_PORT:", dbPort)
	fmt.Println("POSTGRES_DB_NAME:", dbName)

	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable TimeZone=UTC",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)

	fmt.Println("DSN:", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Postgres: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}

	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping Postgres: %w", err)
	}

	log.Println("Successfully connected to Postgres database")

	err = db.AutoMigrate(&models.Payout{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database migrated successfully!")

	return db, nil
}
