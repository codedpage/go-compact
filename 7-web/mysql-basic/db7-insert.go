package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type BookAll struct {
	ID              int
	Name            string
	Language        int
	Author          string
	Pages           int
	PublicationDate string
	AddedOn         string
}

type Book struct {
	BookAll
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

	defer db.Close()

	sql := "INSERT INTO books(name, pages) VALUES ('C1', 150)"
	res, err := db.Exec(sql)
	logOnErr(err)

	lastId, err := res.LastInsertId()
	logOnErr(err)

	fmt.Printf("The last inserted row id: %d\n", lastId)
}

//https://go.dev/play/p/wGcgrtYwSC7
