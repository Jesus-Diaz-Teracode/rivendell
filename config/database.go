package config

import (
	"fmt"

	"github.com/jedi4z/rivendell/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pkg/errors"
)

// Connect creates a new connection with database
// and return a db instance
func CreateDatabase() (*gorm.DB, error) {
	dbHost := "localhost"
	dbPort := "26267"
	dbName := "example_gorm"
	dbSsl := "disable"
	connString := fmt.Sprintf(
		"postgresql://root@%s:%s/%s?sslmode=%s",
		dbHost,
		dbPort,
		dbName,
		dbSsl,
	)

	db, err := gorm.Open("postgres", connString)
	if err != nil {
		return nil, errors.Wrap(err, "Is not posible connect with database")
	}

	sql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName)

	err = db.Exec(sql).Error
	if err != nil {
		return nil, errors.Wrap(err, "Is not posible create database")
	}

	db.AutoMigrate(&models.User{})

	return db, nil
}
