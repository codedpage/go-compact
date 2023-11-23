package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

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

	var n int
	fmt.Printf("Enter the book id: ")
	fmt.Scanf("%d", &n)

	//sql := fmt.Sprintf("DELETE FROM books WHERE id  = '%d' ; DELETE FROM books WHERE id  = '15' ", n)
	sql := fmt.Sprintf("DELETE FROM books WHERE id  = '%d' ", n)
	res, err := db.Exec(sql)
	logOnErr(err)

	affectedRows, err := res.RowsAffected()
	logOnErr(err)

	fmt.Printf("The statement affected %d rows\n", affectedRows)
}

//https://go.dev/play/p/0W-Je0DNPxj
