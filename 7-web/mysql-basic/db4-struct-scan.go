package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Bk struct {
	ID      int
	Name    string
	AddedOn string
}

var db *sql.DB
var err error

func init() {
	//tmpDB, err := sql.Open("postgres", "user=postgres password=postgres dbname=dev sslmode=disable")
	tmpDB, err := sql.Open("mysql", "root:@tcp(localhost:3307)/test")
	logOnErr(err)
	db = tmpDB
	//	defer db.Close()
}

func logOnErr(err error) {
	if err != nil {
		//log.Println(err.Error())
		panic(err)
	}
}

/*
-Struct (ID, Name, AddedOn)

#Mandatory
-Query (seelct id,name from books)
-Scan (&ID, &Name)

*/

func main() {

	//multiple row ///////////////////////////////////////////////////
	rows, err := db.Query("SELECT id, name FROM `books` ")
	logOnErr(err)

	for rows.Next() {
		var book Bk
		err = rows.Scan(&book.ID, &book.Name)
		logOnErr(err)

		fmt.Println(book)
	}

}

//https://go.dev/play/p/bqp22tY0N6F
