package main

import (
      "fmt"
      "net/http"
)

func main() {
      //resp, err := http.Get("https://siongui.github.io/")
	  resp, err := http.Get("http://manojkumar.info/")
      if err != nil {
              panic(err)
      }
      defer resp.Body.Close()

      for k, v := range resp.Header {
              fmt.Print(k)
              fmt.Print(" : ")
              fmt.Println(v)
      }
}