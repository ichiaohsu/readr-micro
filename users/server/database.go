package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// var DB database = database{nil}

// type database struct {
// 	*sqlx.DB
// }

func Connect(dbURI string) (*sqlx.DB, error) {
	d, err := sqlx.Open("postgres", dbURI)
	if err != nil {
		log.Panic(err)
	}
	if err = d.Ping(); err != nil {
		log.Panic(err)
	}
	return d, err
}

// func CreateConnection() (*gorm.DB, error) {
// 	return nil, nil
// }
