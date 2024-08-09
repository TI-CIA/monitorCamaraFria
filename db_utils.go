package main

import (
    "database/sql"
//    "encoding/json"
//    "fmt"
    "log"
//    "net/http"
//    "strconv"


    _ "github.com/go-sql-driver/mysql"
//    "github.com/gorilla/mux"
)

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("mysql","cold_storage_adm:test_password@tcp(127.0.0.1:3306)/cold_storage")
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(10)
}