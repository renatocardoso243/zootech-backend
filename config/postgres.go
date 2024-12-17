package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/renatocardoso243/gopportunities.git/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializePostgres() (*gorm.DB, error) {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// String to connect PostgreSQL
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		"postgres", // Connect to default database
		os.Getenv("DB_PORT"),
	)

	// Initial connection to default postgres
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return nil, err
	}

	// Create database zootech if not exist
	if err := createDatabase(database, "zootech"); err != nil {
		log.Fatalf("Failed to create database 'zootech': %v", err)
		return nil, err
	}

	// Update dsn to connect to "zootech"
	dsn = fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		"zootech",
		os.Getenv("DB_PORT"),
	)

	// Conecta ao banco "zootech"
	database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database 'zootech': %v", err)
		return nil, err
	}

	DB = database
	fmt.Println("Database connected successfully on port: " + os.Getenv("DB_PORT"))

	// Migrate the schemas to created DB.
	if err := database.AutoMigrate(
		&schemas.Animal{},
		&schemas.User{},
		&schemas.Herd{},
		&schemas.Diet{},
		&schemas.Weight{},
	); err != nil {
		log.Fatalf("Database migration error: %v", err)
		return nil, err
	}

	fmt.Println("Database migration completed successfully.")
	return database, nil
}

// Function to create the DB if it doesn't exist
func createDatabase(db *gorm.DB, dbName string) error {
	// Check the existense of the database
	exists := false
	err := db.Raw("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = ?)", dbName).Scan(&exists).Error
	if err != nil {
		return fmt.Errorf("failed to check if database exists: %v", err)
	}

	// Create DB if it doesn't exist
	if !exists {
		if err := db.Exec("CREATE DATABASE " + dbName).Error; err != nil {
			return fmt.Errorf("failed to create database: %v", err)
		}
		fmt.Printf("Database '%s' created successfully.\n", dbName)
	} else {
		fmt.Printf("Database '%s' already exists.\n", dbName)
	}

	return nil
}
