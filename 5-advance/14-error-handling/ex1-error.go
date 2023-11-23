package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
)

// import "os"
func main1() {
	f, err := os.Open("/test.txt")
	if err != nil {
		fmt.Println(err)         //no such file or directory
		fmt.Println(err.Error()) //no such file or directory
		return
	}
	fmt.Println(f.Name(), "opened successfully")
}

// import "os"
func main2() {
	f, err := os.Open("/test.txt")
	if err, ok := err.(*os.PathError); ok {
		fmt.Println(err)         //no such file or directory
		fmt.Println(err.Error()) //no such file or directory

		fmt.Println(err.Err, "|", err.Op, "|", err.Path) //no such file or directory | open | /test.txt
		fmt.Println(err.Error(), "|", err.Timeout(), "|", err.Unwrap(), "|", err.Err.Error(), "|", err.Unwrap().Error())
		//open /test.txt: no such file or directory | false | no such file or directory | no such file or directory | no such file or directory
		return
	}
	fmt.Println(f.Name(), "opened successfully")
}

// import "net"
func main3() {
	addr, err := net.LookupHost("manojkumars.info")
	if err, ok := err.(*net.DNSError); ok {
		if err.Timeout() {
			fmt.Println("operation timed out")
		} else if err.Temporary() {
			fmt.Println("temporary error")
		} else {
			fmt.Println("generic error: ", err)
		}
		return
	}
	fmt.Println(addr)
}

// import "path/filepath"
func main4() {
	f, err := filepath.Glob("[")
	if err != nil {
		fmt.Println(err)         //syntax error in pattern
		fmt.Println(err.Error()) //syntax error in pattern
		return
	}
	fmt.Println("matched files", f)
}

// supress error
func main() {
	f, _ := filepath.Glob("[")
	fmt.Println("matched files", f)
}

//https://go.dev/play/p/fx-sBT8xXU4
