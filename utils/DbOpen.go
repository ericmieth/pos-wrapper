package utils

import (
	"database/sql"
	//"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func DbOpen(dbName string, configMap map[string]string) *sql.DB {

	dbUser := configMap["dbUser"]
	dbPass := configMap["dbPass"]
	dbHost := configMap["dbHost"]
	dbPort := configMap["dbPort"]

	dbSettings := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName

	db, err := sql.Open("mysql", dbSettings)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
