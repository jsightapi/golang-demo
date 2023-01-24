package main

import (
	"fmt"
	"io"
	"net/http"
)

var jSight JSight

func main() {
	jSight = NewJSight("./jsightplugin.so")
	http.HandleFunc("/", handle)
	fmt.Println("Listening on 8000 port…")
	http.ListenAndServe(":8000", nil)
}

func handle(w http.ResponseWriter, req *http.Request) {

	jsightSpecPath := "./my-api-spec.jst"
	reqBody, _ := io.ReadAll(req.Body)

	jSight.ClearCache() // Comment this line in production.

	err := jSight.ValidateHTTPRequest(
		jsightSpecPath,
		req.Method,
		req.RequestURI,
		req.Header,
		reqBody,
	)

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.ToJSON()))
		return
	}

	responseStatusCode := 200
	responseBody := []byte("\"User created!\"\n")
	responseHeaders := http.Header{}

	err = jSight.ValidateHTTPResponse(
		jsightSpecPath,
		req.Method,
		req.RequestURI,
		responseStatusCode,
		responseHeaders,
		responseBody,
	)

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.ToJSON()))
		return
	}

	w.WriteHeader(responseStatusCode)
	w.Write([]byte(responseBody))
}
