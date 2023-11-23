package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	/*
	       dt := time.Now()

	   	layout := "02/01/2006 15:04:05"
	   	dt, _ := time.Parse(layout, "28/07/2020 19:54:05") //time
	   	fmt.Printf("s1 == %v | %T", dt, dt)                //s1 == 2020-05-31 23:59:59 +0530 IST | time.Time

	   	i, _ := strconv.ParseInt("1590949799", 10, 64)
	   	dt := time.Unix(i, 0)                              //time
	   	fmt.Printf("s1 == %v | %T", dt, dt)                //s1 == 2020-05-31 23:59:59 +0530 IST | time.Time
	   	............................................................
	   	fmt.Println()

	   	ts := dt.Unix()
	   	fmt.Println("timestamp(sec)", ts)                  //1590949799

	   	tss := dt.UnixNano() / 1000000
	   	fmt.Println("timestamp(millisec)", tss)           //1590949799000

	   	timeString := time.Unix(ts, 0)
	   	fmt.Println("time String:", timeString)            //2020-05-31 23:59:59 +0530 IST

	   	dateString := dt.Format("02012006")
	   	fmt.Println("date String: ", dateString)           //31052020


	*/
	//ctime
	ts1 := time.Now().UnixNano() / 1000000
	fmt.Printf("\ntimestamp (millisecond) = %d", ts1) //1600922817509

	//0
	tnow := time.Now() //2020-09-08 12:19:24.6150492 +0530 IST m=+0.006501901
	fmt.Println("\ntm3=", tnow)
	fmt.Printf("s3 == %v | %T", tnow, tnow)

	ts := tnow.UnixNano() / 1000000
	fmt.Printf("\ntimestamp (millisecond) = %d", ts) //1600922817509

	timeString := time.Unix(ts/1000, 0)
	fmt.Printf("\ntime String: %s\n", timeString) //2020-09-24 10:16:57 +0530 IST
	fmt.Println("----\n")

	//1
	var ts2 int64 = 1590949799000
	i1 := ts2 / 1000
	tm1 := time.Unix(i1, 0)
	fmt.Println("tm1=", tm1)
	fmt.Printf("s1 == %v | %T", tm1, tm1)
	a, b, c := tm1.Date()
	fmt.Println("\na, b, c = ", a, b, int(b), c)

	fmt.Println("\nBase Month:", tm1)
	after := tm1.AddDate(0, 1, 0)
	fmt.Println("Add 1 Month:", after)
	fmt.Println("----")

	//2
	var tString = "1590949799"
	i2, _ := strconv.ParseInt(tString, 10, 64)
	tm2 := time.Unix(i2, 0)
	fmt.Println("\ntm2=", tm2)
	fmt.Printf("s2 == %v | %T", tm2, tm2)
	fmt.Println("\n----")

	//3
	tnow1 := time.Now() //2020-09-08 12:19:24.6150492 +0530 IST m=+0.006501901
	fmt.Println("\ntm3=", tnow1)
	fmt.Printf("s3 == %v | %T", tnow1, tnow1)

	ts = tnow.Unix() //1599547764
	fmt.Printf("\ntimestamp = %d", ts)

	timeString1 := time.Unix(ts, 0)
	fmt.Printf("\ntime String: %s\n", timeString1)
	fmt.Println("----\n")

	//4-1
	layout := "02-01-2006 15:04:05"
	dt, _ := time.Parse(layout, "12-05-2020 16:23:36")
	formatedDate := dt.Format("02012006")
	fmt.Println("dt", dt, "|", formatedDate)
	fmt.Printf("s41 == %v | %T", dt, dt)

	//4-2 Date to timestamp for database
	layout = "01/02/2006 3:04:05 PM"
	dt, _ = time.Parse(layout, "07/28/2020 7:54:05 PM")
	formatedDate = dt.Format("02012006")
	fmt.Println("\n\ndt", dt, "|", formatedDate)
	fmt.Printf("s42 == %v | %T", dt, dt)

	//4-3 Date to timestamp for csv
	layout = "02/01/2006 15:04:05"
	dt, _ = time.Parse(layout, "28/07/2020 19:54:05")
	tss := dt.UnixNano() / 1000000
	formatedDate = dt.Format("02012006")
	fmt.Println("\n\ndt", dt, "|", formatedDate)
	fmt.Printf("s43 == %v | %T", dt, dt)
	fmt.Println("\ntimestamp(millisec)", tss)

}

//https://go.dev/play/p/q1PHLxPU4E6