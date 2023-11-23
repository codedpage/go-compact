package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type BookOrder struct {
	A int
	B string
	D sql.NullString
	E int
	F sql.NullString
	G string
}

type Book struct {
	BookOrder
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

func main() {

	rows, err := db.Query("SELECT * FROM `books` ")
	logOnErr(err)

	for rows.Next() {
		var book Book
		err = rows.Scan(&book.A, &book.B, &book.D, &book.E, &book.F, &book.G)
		logOnErr(err)

		fmt.Print(book.A, "====> ")
		fmt.Println(book)

	}

}

//https://go.dev/play/p/7GOSarJlude
