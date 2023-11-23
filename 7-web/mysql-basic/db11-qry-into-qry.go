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

	//Category
	rows, err := db.Query("SELECT id, name FROM `language`")
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		logOnErr(err)
		fmt.Println("Language :: ", id, name)

		//Item
		var id1 int
		var name1 string
		rows1, err := db.Query("SELECT id, name FROM `books`  where language = ? ", id)
		for rows1.Next() {
			err = rows1.Scan(&id1, &name1)
			logOnErr(err)

			fmt.Println("Book : ", id1, name1)

		}
		fmt.Println("--------------------------")

	}

}

//https://go.dev/play/p/XLYtV13L_WS
