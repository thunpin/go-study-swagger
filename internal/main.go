package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello, %q", html.EscapeString(request.URL.Path))
	})

	log.Println("Listening on port :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
