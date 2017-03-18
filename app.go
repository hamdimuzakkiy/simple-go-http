package main

import (
	"net/http"

	"github.com/hamdimuzakkiy/simple-go-http/src/handler"
)

func main() {
	http.HandleFunc("/", handler.MainHandler)

	http.ListenAndServe(":8080", nil)
}
