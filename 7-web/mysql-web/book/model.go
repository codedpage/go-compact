package book

import (
	"database/sql"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	//tmpDB, err := sql.Open("postgres", "user=postgres password=postgres dbname=dev sslmode=disable")
	tmpDB, err := sql.Open("mysql", "root:root@2021@tcp(localhost:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	db = tmpDB
}

func getBook(bookID int) (Book, error) {
	//Retrieve
	res := Book{}

	var id int
	var name string
	var author string
	var pages int
	var publicationDate mysql.NullTime

	err := db.QueryRow(`SELECT id, name, author, pages, publication_date FROM books where id = ?`, bookID).Scan(&id, &name, &author, &pages, &publicationDate)
	if err == nil {
		res = Book{ID: id, Name: name, Author: author, Pages: pages, PublicationDate: publicationDate.Time}
	}

	return res, err
}

func allBooks() ([]Book, error) {
	//Retrieve
	books := []Book{}

	rows, err := db.Query(`SELECT id, name, author, pages, publication_date FROM books order by id`)
	defer rows.Close()
	if err == nil {
		for rows.Next() {
			var id int
			var name string
			var author string
			var pages int
			var publicationDate mysql.NullTime

			err = rows.Scan(&id, &name, &author, &pages, &publicationDate)
			if err == nil {
				currentBook := Book{ID: id, Name: name, Author: author, Pages: pages}
				if publicationDate.Valid {
					currentBook.PublicationDate = publicationDate.Time
				}

				books = append(books, currentBook)
			} else {
				return books, err
			}
		}
	} else {
		return books, err
	}

	return books, err
}

func insertBook(name, author string, pages int, publicationDate time.Time) (int, error) {
	//Create
	var bookID int64
	res, err := db.Exec(`INSERT INTO books(name, author, pages, publication_date) VALUES(?, ?, ?, ?)`, name, author, pages, publicationDate) //.Scan(&bookID)

	bookID, err = res.LastInsertId()
	if err != nil {
		return 0, err
	}

	log.Printf("Last inserted ID: %v\n", int(bookID))
	return int(bookID), err
}

func updateBook(id int, name, author string, pages int, publicationDate time.Time) (int, error) {
	//Update
	_, err := db.Exec(`UPDATE books set name=?, author=?, pages=?, publication_date=? where id=? `, name, author, pages, publicationDate, id)
	if err != nil {
		return 0, err
	}

	//rowsUpdated, err := res.RowsAffected()
	rowsUpdated := id

	if err != nil {
		return 0, err
	}

	return int(rowsUpdated), err
}

func removeBook(bookID int) (int, error) {
	//Delete
	res, err := db.Exec(`delete from books where id = ?`, bookID)
	if err != nil {
		return 0, err
	}

	rowsDeleted, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowsDeleted), nil
}
