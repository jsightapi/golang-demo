package main

import (
	"fmt"
	"net/http"
)

func handle(w http.ResponseWriter, req *http.Request) {

	handleResponse(w, req)

}

func main() {
	loadJSightPlugin("./jsight/jsight_validator_plugin.so")
	http.HandleFunc("/", handle)
	fmt.Println("Listening on 8000 port")
	http.ListenAndServe(":8000", nil)
}
