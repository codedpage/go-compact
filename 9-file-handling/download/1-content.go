package main

import (
	"bytes"
	"log"
	"net/http"
	"time"
)

func main() {
	log.SetFlags(log.Lshortfile)
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		modtime := time.Now()

		data := []byte("Hello World")
		content := bytes.NewReader(data)

		// ServeContent uses the name for mime detection
		const name = "random.txt"

		// tell the browser the returned content should be downloaded
		w.Header().Add("Content-Disposition", "Attachment")

		http.ServeContent(w, req, name, modtime, content)
	})

	log.Print("Serving %s on HTTP port :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
