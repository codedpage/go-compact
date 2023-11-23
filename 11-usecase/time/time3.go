package main

import (
	"fmt"
	"time"
)

func main() {
	var stime, etime string
	ist, _ := time.LoadLocation("Asia/Kolkata")

	today := time.Now().In(ist)
	today = Bod(today)
	//go back to DayOfMonth days
	stime = today.AddDate(0, 0, -1).Format("2006-01-02 15:04:05")
	etime = today.AddDate(0, 0, 1).Format("2006-01-02 15:04:05")

	fmt.Println(stime, etime)
}

func main0() {
	fmt.Println("Hello, playground")
	//lmt := new(int)
	var lmt *int
	*lmt = 5000
	var stime, etime string
	ist, err := time.LoadLocation("Asia/Kolkata")

	today := time.Now().In(ist)
	today = Bod(today)
	//go back to DayOfMonth days
	stime = today.AddDate(0, 0, -1).Format("2006-01-02 15:04:05")
	etime = today.AddDate(0, 0, 1).Format("2006-01-02 15:04:05")

	fmt.Println(stime, etime)
	limit := 5000 //limit is 5000 by default
	if *lmt > 0 {
		limit = *lmt
	}

	fmt.Println("stime = ", stime)
	fmt.Println("etime = ", etime)
	fmt.Println("limit = ", limit)
	fmt.Println(err)

}

func Bod(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

//https://go.dev/play/p/w1mTVIIIZaC