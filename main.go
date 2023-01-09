package main

import (
	"fmt"
	"io"
	"net/http"
)

func handle(w http.ResponseWriter, req *http.Request) {

	reqBody, _ := io.ReadAll(req.Body)

	err := jSight.ValidateHTTPRequest(
		"./jsight/myapi.jst",
		req.Method,
		req.RequestURI,
		req.Header,
		reqBody,
	)

	fmt.Printf("Reported by: %s, %d\n", err.ReportedBy(), err.Position().Index())

	handleResponse(w, req)
}

var jSight JSight

func main() {
	jSight = NewJSight("./plugins/jsightplugin.so")
	http.HandleFunc("/", handle)
	fmt.Println("Listening on 8000 port")
	http.ListenAndServe(":8000", nil)
}
