package models

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func HandleError(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}

}

func GetDBConnection() *sql.DB {

	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/digiease_final")

	HandleError(err)

	return db

}

func GetCurrentTimeStamp() string {

	currentTime := time.Now().Local()

	strFormattedTime := currentTime.Format("2006-01-02 15:04:05")

	return strFormattedTime

}
