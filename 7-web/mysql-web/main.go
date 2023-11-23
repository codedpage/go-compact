package main

import (
	"log"
	. "mysql/book"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("view/assets"))))

	http.HandleFunc("/", HandleListBooks)
	http.HandleFunc("/book.html", HandleViewBook)
	http.HandleFunc("/save", HandleSaveBook)
	http.HandleFunc("/delete", HandleDeleteBook)

	log.Println("Listening and serving HTTP on :3000")
	http.ListenAndServe(":3000", nil)
}
