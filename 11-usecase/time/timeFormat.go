package main

import (
	"fmt"
	"time"
)

var print = fmt.Println
var echo = fmt.Print
var printf = fmt.Printf
var br = fmt.Println

func main() {

	timestamp := "12-05-2020 16:23:36"
	//timestamp = "28-03-2016 11:50:50"

	fmt.Println(timestamp)

	const layout = "02-01-2006 15:04:05"

	dt, _ := time.Parse(layout, timestamp)

	FileCreationDate := dt.Format("02012006")

	fmt.Println(dt)

	fmt.Println(FileCreationDate)

}

/*
https://yourbasic.org/golang/format-parse-string-time-date-example/#standard-time-and-date-formats
Layout options
Type 	Options
Year 	06   2006
Month 	01   1   Jan   January
Day 	02   2   _2   (width two, right justified)
Weekday 	Mon   Monday
Hours 	03   3   15
Minutes 	04   4
Seconds 	05   5
ms μs ns 	.000   .000000   .000000000
ms μs ns 	.999   .999999   .999999999   (trailing zeros removed)
am/pm 	PM   pm
Timezone 	MST
Offset 	-0700   -07   -07:00   Z0700   Z07:00
*/

//https://go.dev/play/p/wLlHbW0qkG-