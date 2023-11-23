package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

/*

CREATE TABLE `books` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL,
  `author` varchar(50) DEFAULT NULL,
  `pages` int DEFAULT NULL,
  `publication_date` date DEFAULT NULL,
  `addedOn` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=5 DEFAULT CHARSET=latin1

INSERT INTO `books` (`id`, `name`, `author`, `pages`, `publication_date`, `addedOn`) VALUES
(1, 'Go', 'A Singh', 400, '2020-12-02', '2021-04-03 04:10:18'),
(2, 'Java', 'B K SINGH', 230, '2021-03-24', '2021-04-03 04:10:18'),
(3, 'C', 'C K sing', 22, '2021-03-24', '2021-04-03 04:10:18'),
(4, 'C++', 'D k sing', 33, '2021-03-24', '2021-04-03 04:10:18');

*/

type BookAll struct {
	ID              int
	Name            string
	Author          sql.NullString
	Pages           int
	PublicationDate sql.NullString
	AddedOn         string
}

type Book struct {
	BookAll
}

var db *sql.DB

func init() {
	//tmpDB, err := sql.Open("postgres", "user=postgres password=postgres dbname=dev sslmode=disable")
	tmpDB, err := sql.Open("mysql", "root:@tcp(localhost:3307)/test")
	logOnErr(err)
	db = tmpDB
	//	defer db.Close()
}

func logOnErr(err error) {
	if err != nil {
		log.Println(err.Error())
		//panic(err)
	}
}

func main() {

	main1()
	main2()
	main3()
	main4()
	defer db.Close()
}

//By struct
func main1() {
	rows, err := db.Query("SELECT * FROM `books` ")
	logOnErr(err)

	for rows.Next() {
		var book Book
		err = rows.Scan(&book.ID, &book.Name, &book.Author, &book.Pages, &book.PublicationDate, &book.AddedOn)
		logOnErr(err)

		fmt.Print(book.ID, "====> ")
		fmt.Println(book)

	}

}

//by variables (without struct)
func main2() {
	rows, err := db.Query("SELECT id, name FROM `books` limit 100 ")
	logOnErr(err)

	for rows.Next() {
		//var book Book
		var id int
		var name string
		cl := []interface{}{&id, &name}
		err = rows.Scan(cl...)
		logOnErr(err)

		//fmt.Print(cl)
		fmt.Println(id, " - ", name)
	}
}

//Unknown field + Unknown datatype
func main3() {
	rows, err := db.Query("SELECT * FROM `books` limit 2 ")
	logOnErr(err)

	for rows.Next() {
		cols, err := rows.Columns()
		logOnErr(err)
		vals := make([]interface{}, len(cols))
		for i, _ := range cols {
			vals[i] = new(sql.RawBytes)
		}

		err = rows.Scan(vals...)
		logOnErr(err)

		//fmt.Println(vals)
		for k1, v1 := range vals {
			log.Printf("%d %s", k1, *v1.(*sql.RawBytes))
		}

	}

}

//Unknown field + Known datatype
func main4() {
	rows, err := db.Query("SELECT id, name, pages FROM `books` limit 2")
	logOnErr(err)

	for rows.Next() {
		val := []interface{}{
			new(uint64),
			new(string),
			new(uint64),
		}
		err = rows.Scan(val...)
		logOnErr(err)

		fmt.Printf("%d ", *val[0].(*uint64))
		fmt.Printf("%s ", *val[1].(*string))
		p := val[2]
		fmt.Printf("%d \n", *p.(*uint64))
	}
}

//https://go.dev/play/p/pzQpwn6t9Ol
