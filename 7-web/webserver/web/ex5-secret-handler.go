package main

import (
	"fmt"
	"net/http"
	"time"
)

// UptimeHandler writes the number of seconds since starting to the response.
type UptimeHandler struct {
	Started time.Time
}

func NewUptimeHandler() UptimeHandler {
	return UptimeHandler{Started: time.Now()}
}

// --------
type SecretTokenHandler struct {
	next   http.Handler
	secret string
}

func main() {
	http.Handle("/", SecretTokenHandler{next: NewUptimeHandler(), secret: "MySecret"})
	http.ListenAndServe(":3000", nil)
	//http://localhost:3000/?secret_token=MySecret
}

// /////
func (h UptimeHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(
		w,
		fmt.Sprintf("Current Uptime: %s", time.Since(h.Started)),
	)
}

// ServeHTTP makes SecretTokenHandler implement the http.Handler interface.
func (h SecretTokenHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	if req.URL.Query().Get("secret_token") == h.secret {

		h.next.ServeHTTP(w, req)
	} else {

		http.NotFound(w, req)
	}
}
