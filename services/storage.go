package services

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"url-shortener/configs"
)

func CreateDatabaseIFNotExist(connectionString, dbName string) error {
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return err
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(db)

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	if err != nil {
		return err
	}
	return nil
}

func SaveLongShortToDB(config configs.GlobalConfig, longUrl, shortUrl, dbName string) error {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/", config.MySQL.User, config.MySQL.Pass, config.MySQL.Host, config.MySQL.Port)
	err := CreateDatabaseIFNotExist(connectionString, dbName)
	if err != nil {
		return err
	}

	connectionString = connectionString + dbName

	// Connect to DB
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(db)

	// Ping the database to check the connection
	err = db.Ping()
	if err != nil {
		return err
	}

	// Database connection is established
	log.Info("Connected to database: ", dbName)

	// Create table url_map if not exist
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS url_map (long_url VARCHAR(255) NOT NULL, short_url_hash VARCHAR(255) NOT NULL, PRIMARY KEY (short_url_hash))")
	if err != nil {
		return err
	}

	// Insert longUrl and shortUrl to table url_map
	_, err = db.Exec("INSERT INTO url_map (long_url, short_url_hash) VALUES (?, ?)", longUrl, shortUrl)
	if err != nil {
		return err
	}

	return nil
}
