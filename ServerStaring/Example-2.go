package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World ::::::::::::::::::----------------------")
	})
	fmt.Println("Server running to port:", 8080)
	http.ListenAndServe(":8080", nil)

}

