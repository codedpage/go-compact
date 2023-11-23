package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	//1
	tma := time.Unix(1590949799000/1000, 0)
	fmt.Println("tma= ", tma)

	i, err := strconv.ParseInt("1585502600", 10, 64)
	if err != nil {
		panic(err)
	}
	tm := time.Unix(i, 0)
	fmt.Println("tm=", tm)

	fmt.Printf("x1 == %v %T", tma, tma)
	fmt.Println()

	//2
	t := time.Now()
	fmt.Printf("x2 == %v %T", t, t)
	fmt.Println()

	//3
	a, b, c := tm.Date()
	fmt.Println("\na,b,c = ", a, b, c, int(b))

	fmt.Println("\nBase Month:", tm)
	after := tm.AddDate(0, 1, 0)
	fmt.Println("Add 1 Month:", after)

	//4
	timestamp := "12-05-2020 16:23:36"
	fmt.Println("\n", timestamp)
	var layout = "02-01-2006 15:04:05"
	dt, _ := time.Parse(layout, timestamp)
	FileCreationDate := dt.Format("02012006")
	fmt.Println("dt", dt)
	fmt.Println("FileCreationDate", FileCreationDate)

	//5
	tNow := time.Now()
	fmt.Printf("\nx3 == %v %T", tNow, tNow)

	//time.Time to Unix Timestamp
	tUnix := tNow.Unix()
	fmt.Printf("\ntimeUnix %d\n", tUnix)

	//Unix Timestamp to time.Time
	timeT := time.Unix(tUnix, 0)
	fmt.Printf("time.Time: %s\n", timeT)

	//Date to timestamp for database
	layout = "01/02/2006 3:04:05 PM"
	t, err = time.Parse(layout, "07/28/2020 7:54:05 PM")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Date to timestamp(sec)", t.Unix())

	//Date to timestamp for csv
	layout = "02/01/2006 15:04:05"
	ctime, _ := time.Parse(layout, "28/07/2020 19:54:05")
	ts := ctime.UnixNano() / 1000000
	fmt.Println("Date to timestamp(millisec)", ts)

}
//https://go.dev/play/p/Ntn94MgghJ1
