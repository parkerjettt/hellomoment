package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", HelloServer)
	http.ListenAndServe(":8080", nil)

}

func HelloServer(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "hello world!")
	log.Println("hello world!")

}
