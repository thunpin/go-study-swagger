package main

import (
	"fmt"
	"html"
	"net/http"
)

func main() {
	fmt.Println("vim-go")
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello, %q", html.EscapeString(request.URL.Path))
	})
}
