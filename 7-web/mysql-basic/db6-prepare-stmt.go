package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type BookSingle struct {
	ID              int
	Name            string
	Author          sql.NullString
	Pages           int
	PublicationDate sql.NullString
	AddedOn         string
}

type Book struct {
	BookSingle
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

//simple
func main1() {

	var n int
	fmt.Printf("Enter the book id: ")
	fmt.Scanf("%d", &n)

	str := "SELECT name  FROM `books`  where id = ? "
	var name string
	err := db.QueryRow(str, n).Scan(&name)
	logOnErr(err)

	fmt.Println(name)
	defer db.Close()
}

//prepare + single value
func main2() {

	var n int
	fmt.Printf("Enter the book id: ")
	fmt.Scanf("%d", &n)

	str := "SELECT name  FROM `books`  where id = ? "
	stmt, err := db.Prepare(str)
	logOnErr(err)

	var name string
	fmt.Println(name)
	err = stmt.QueryRow(n).Scan(&name)
	logOnErr(err)

	fmt.Println(name)

	defer stmt.Close()
	defer db.Close()

}

//prepare + struct
func main() {

	//multiple row ///////////////////////////////////////////////////
	var n int
	fmt.Printf("Enter the book id: ")
	fmt.Scanf("%d", &n)

	var book Book
	fmt.Println(book)

	str := "SELECT *  FROM `books`  where id = ? "
	stmt, err := db.Prepare(str)
	logOnErr(err)

	err = stmt.QueryRow(n).Scan(&book.ID, &book.Name, &book.Author, &book.Pages, &book.PublicationDate, &book.AddedOn)
	logOnErr(err)
	fmt.Println(book)

	defer stmt.Close()
	defer db.Close()

}

//https://go.dev/play/p/wINQrUKBSox
