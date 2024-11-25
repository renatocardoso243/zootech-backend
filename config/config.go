package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
	logger *Logger
)

func Init() error {
	var err error
	//Initialize Postgres DB
	db, err = InitalizePostgres()
	if err != nil {
		return fmt.Errorf("failed to initialize postgres: %v", err)
	}
	return nil
}

func GetPostgres() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}