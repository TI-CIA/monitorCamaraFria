package main

import (
//    "database/sql"
//    "encoding/json"
    "fmt"
    "log"
    "net/http"
//    "strconv"


    _ "github.com/go-sql-driver/mysql"
    "github.com/gorilla/mux"
)


type Status struct {
    ColdStorageID       string  `json:"id"`
    Temperature0        float64 `json:"temp0"`
    Temperature1        float64 `json:"temp1"`
    Temperature2        float64 `json:"temp2"`
    Temperature3        float64 `json:"temp3"`
    Temperature4        float64 `json:"temp4"`
    Temperature5        float64 `json:"temp5"`
    MainBoardFailure    bool `json:"fail"`
}

var privateKey = "asdfg"

func main() {
	initDB()
	defer db.Close()

	// initialize router
	router := mux.NewRouter()
    // router.HandleFunc("/status", getStatus).Methods("GET")
    router.HandleFunc("/status", createStatusLog).Methods("POST")
    // router.HandleFunc("/status", deleteStatus).Methods("DELETE")
    // router.HandleFunc("/status", updateStatus).Methods("UPDATE")
    
	// start http server
	fmt.Println("Server listening on port 8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}