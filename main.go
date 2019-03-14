package main

import (
	"fmt"
	"net/http"
	"log"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Medium!")
}

func main() {
	fmt.Print("Hello App listening on 3000\n")
	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
