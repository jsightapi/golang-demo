package main

import (
	"fmt"
	"net/http"
)

func handleResponse(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
	fmt.Println("request")
}
