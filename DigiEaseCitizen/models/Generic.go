package models

import "database/sql"
import "fmt"
import _ "github.com/go-sql-driver/mysql"

func HandleError(err error) {

	if err != nil {
		fmt.Println(err.Error())
	}

}

func GetDBConnection() *sql.DB {

	db, err := sql.Open("mysql", "root:manthan@tcp(127.0.0.1:3306)/digiease")
	HandleError(err)

	return db

}
