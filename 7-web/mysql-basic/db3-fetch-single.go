package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type BookSingle struct {
	ID              int
	Name            string
	Author          sql.NullString //to avoid db NULL, OR => Set default value in field of table.
	Pages           int
	PublicationDate sql.NullString //sql.NullTime
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
func main() {

	defer db.Close()
	//multiple row ///////////////////////////////////////////////////
	var n int
	fmt.Printf("Enter the book id: ")
	fmt.Scanf("%d", &n)
	var book Book

	//case-1
	//err = db.QueryRow("SELECT * FROM `books` where id = ? ", n).Scan(&book.ID, &book.Name, &book.Author, &book.Pages, &book.PublicationDate, &book.AddedOn)

	//case-2 [ select * == Scan(...), here 6 = 6 ]
	row := db.QueryRow("SELECT * FROM `books` where id = ? ", n)                                          //6
	err = row.Scan(&book.ID, &book.Name, &book.Author, &book.Pages, &book.PublicationDate, &book.AddedOn) //6

	logOnErr(err)

	fmt.Print(book.ID, "====> ")
	fmt.Println(book)
	fmt.Println(book.Author.Value())
	fmt.Println(book.Author.String, book.Author.Valid)

	QueryStuff()
}

func QueryStuff() {

	defer db.Close()
	//multiple row ///////////////////////////////////////////////////
	var version string
	var ctime string
	err = db.QueryRow("SELECT VERSION(), now()").Scan(&version, &ctime)
	fmt.Println(version, ctime)

}

//https://go.dev/play/p/_f9hsD2BqN2
