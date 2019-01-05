package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "yo")
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Listenting on port 1234...")
	log.Fatal(http.ListenAndServe(":1234", nil))
}
