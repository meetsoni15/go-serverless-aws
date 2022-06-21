package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/meetsoni15/go-serverless-aws/model"
)

// InitDB initializes the database
func InitDB() *pg.DB {
	//connection string
	var optsConnect *pg.Options
	if optsConnect == nil {
		optsConnect = &pg.Options{
			Addr:     fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Database: os.Getenv("DB_NAME"),
		}
	}

	//connect to database
	var db = pg.Connect(optsConnect)
	//error handling
	if db == nil {
		log.Println("Connection to database failed")
		panic("Connection to database failed")
	} else {
		log.Println("Connection to database successful")
	}

	//Ping database and checked
	ctx := context.Background()
	if err := db.Ping(ctx); err != nil {
		panic(err)
	}
	//Create table using data models
	if err := CreateSchema(db); err != nil {
		log.Printf("Error while CreateSchema %v", err)
	}

	return db
}

//CreateSchema -> Create table
func CreateSchema(db *pg.DB) error {
	models := []interface{}{
		&model.UserRole{},
		&model.User{},
	}

	tx, _ := db.Begin()
	// Make sure to close transaction if something goes wrong.
	defer tx.Close()
	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists:   true,
			FKConstraints: true,
			Temp:          false,
		})

		if err != nil {
			log.Printf("SQL Error while creating tables %v", err)
			tx.Rollback()
			return err
		}
	}

	// Commit on success.
	if err := tx.Commit(); err != nil {
		log.Printf("SQL Error while commiting create schema transaction %v", err)
		return err
	}

	return nil
}
