package main

import (
//    "database/sql"
//    "encoding/json"
//    "fmt"
    "log"
//    "net/http"
//    "strconv"


    _ "github.com/go-sql-driver/mysql"
//    "github.com/gorilla/mux"
)

// func getAllItems() []Item {
    
// 	rows, err := db.Query("SELECT * FROM golang_api")
//     if err != nil {
//         log.Fatal(err)
//         defer rows.Close()
//     }

//     var items []Item
//     for rows.Next() {
//         var item Item
//         if err := rows.Scan(&item.ID, &item.Nome, &item.Descricao, &item.Preco); err != nil {
//             log.Fatal(err)
//         }
//         items = append(items, item)
//     }

//     return items
// }

// func getItemById(id int) []Item {

//     rows, err := db.Query("SELECT * FROM golang_api WHERE ID=? LIMIT 1",id)
//     if err != nil {
//         log.Fatal(err)
//         defer rows.Close()
//     }

//     var items []Item
//     for rows.Next() {
//         var item Item
//         if err := rows.Scan(&item.ID, &item.Nome, &item.Descricao, &item.Preco); err != nil {
//             log.Fatal(err)
//         }
//         items = append(items, item)
//     }

//     return items
// }

func createStatusLog_db(status Status) bool {

    _, err := db.Exec("INSERT INTO status_log (cold_storage_id,temp0,temp1,temp2,temp3,temp4,temp5,main_board_failure) VALUES (?,?,?,?,?,?,?,?)", status.ColdStorageID, status.Temperature0, status.Temperature1, status.Temperature2, status.Temperature3, status.Temperature4, status.Temperature5)
    if err != nil {
        log.Fatal(err)
        return false
    }
    return true
}

// func deleteItem_db(id int) bool {

//     _, err := db.Exec("DELETE FROM golang_api WHERE id=?",id)
//     if err != nil {
//         log.Fatal(err)
//         return false
//     }
//     return true
// }

// func updateItem_db(item Item) bool {
      
//     _, err :=db.Exec("UPDATE golang_api SET nome=?,descricao=?,preco=? WHERE id=?",item.Nome,item.Descricao,item.Preco,item.ID)
//     if err != nil {
//         log.Fatal(err)
//         return false
//     }
//     return true
// }