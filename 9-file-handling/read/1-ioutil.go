package main

import (
    "fmt"
    "io/ioutil"
    "log"
)

func main() {

    content, err := ioutil.ReadFile("1.txt")

     if err != nil {
          log.Fatal(err)
     }

    fmt.Println(string(content))
}