package main

import (
	"log"
	"net/http"
)
  
func main() {
	fs := http.FileServer(http.Dir("./build"))
	http.Handle("/", fs)
  
	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		log.Fatal(err)
	}
}