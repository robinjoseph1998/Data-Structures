package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	port := ":5000"
	fmt.Println("Server Running On :", port)
	http.ListenAndServe(port, nil)

}
