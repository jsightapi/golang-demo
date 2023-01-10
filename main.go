package main

import (
	"fmt"
	"io"
	"net/http"
)

var jSight JSight

func main() {
	jSight = NewJSight("./plugins/jsightplugin.so")
	http.HandleFunc("/", handle)
	fmt.Printf("\n# JSight plugin info\n\n%s\n", jSight.Stat())
	fmt.Println("Listening on 8000 port…")
	http.ListenAndServe(":8000", nil)
}

func handle(w http.ResponseWriter, req *http.Request) {

	jsightSpecPath := "./jsight/myapi.jst"
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
		w.Write([]byte(err.ToJSON()))
		return
	}

	responseStatusCode := 200
	responseBody := "\"Cat is created!\""
	responseHeaders := http.Header{}

	err = jSight.ValidateHTTPResponse(
		jsightSpecPath,
		req.Method,
		req.RequestURI,
		responseStatusCode,
		responseHeaders,
		[]byte(responseBody),
	)

	if err != nil {
		w.Write([]byte(err.ToJSON()))
		return
	}

	w.WriteHeader(responseStatusCode)
	w.Write([]byte(responseBody))
}
