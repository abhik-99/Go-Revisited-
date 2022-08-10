package config

import (
	"gorm.io/driver/sqlite" // Sqlite driver based on GGO
	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"gorm.io/gorm"
)

var db *gorm.DB

// github.com/mattn/go-sqlite3

func Connect() {
	d, err := gorm.Open(sqlite.Open("bookstore.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db = d
}

func GetDb() *gorm.DB {
	return db
}
