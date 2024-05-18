package models

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type DBConfig struct {
	Host, Port, User, Pass, Name string
	Retries                      uint
	RetryInterval                time.Duration
}

func InitDB(config *DBConfig) error {
	var err error

	if DB != nil {
		return fmt.Errorf("DB is already connected")
	}

	if config == nil {
		return fmt.Errorf("no DBConfig provided")
	}

	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		config.User, config.Pass, config.Host, config.Port, config.Name,
	)
	retries := config.Retries

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	for err != nil && retries > 0 {
		log.Printf("%v | retries: %v\n", err, retries)
		retries -= 1
		time.Sleep(config.RetryInterval)
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}
	if err != nil {
		return err
	}

	if err = DB.AutoMigrate(&Subscriber{}); err != nil {
		return err
	}

	return nil
}
