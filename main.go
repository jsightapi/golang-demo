package main

import (
	"fmt"
	"io"
	"net/http"
)

func handle(w http.ResponseWriter, req *http.Request) {

	jsightSpecPath := "./jsight/myapi.jst"
	reqBody, _ := io.ReadAll(req.Body)

	err := jSight.ValidateHTTPRequest(
		jsightSpecPath,
		req.Method,
		req.RequestURI,
		req.Header,
		reqBody,
	)

	if err != nil {
		fmt.Fprintf(w, err.ToJSON())
		return
	}

	handleResponse(w, req)

	/*err = jSight.ValidateHTTPResponse(
		jsightSpecPath,
		req.Method,
		req.RequestURI,
	)*/
}

var jSight JSight

func main() {
	jSight = NewJSight("./plugins/jsightplugin.so")
	http.HandleFunc("/", handle)
	fmt.Println("Listening on 8000 port")
	http.ListenAndServe(":8000", nil)
}
