package models

import "fmt"
import "database/sql"
import "time"

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

func GetTimeStampString(strTimeStamp string) string {

	tmeTimeStamp, err := time.Parse("2006-01-02 15:04:05", strTimeStamp)
	HandleError(err)

	strFormattedTime := tmeTimeStamp.Format("20060102150405")

	return strFormattedTime

}
