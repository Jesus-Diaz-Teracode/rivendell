package config

import (
	"fmt"
	"os"

	"github.com/jedi4z/rivendell/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pkg/errors"
)

// CreateDatabase creates a new connection with
// database and return a db instance
func CreateDatabase() (*gorm.DB, error) {
	connString := fmt.Sprintf(
		"postgresql://root@%s:%s/%s?sslmode=%s",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_NAME"),
		os.Getenv("DATABASE_SSL"),
	)

	db, err := gorm.Open("postgres", connString)
	if err != nil {
		return nil, errors.Wrap(err, "Is not posible connect with database")
	}

	sql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", os.Getenv("DATABASE_NAME"))

	err = db.Exec(sql).Error
	if err != nil {
		return nil, errors.Wrap(err, "Is not posible create database")
	}

	db.AutoMigrate(&models.User{})

	return db, nil
}
