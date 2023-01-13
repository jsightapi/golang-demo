package main

import (
	"net/http"
	"os"
	"strconv"
	"strings"
)

var mockDir = "/tmp/mock/"

func mockStatusCode() int {
	bb, _ := os.ReadFile(mockDir + "response_code")
	c, _ := strconv.Atoi(string(bb))
	return c
}

func mockResponseBody() []byte {
	bb, _ := os.ReadFile(mockDir + "response_body")
	return bb
}

func mockResponseHeaders() http.Header {
	h, _ := os.ReadFile(mockDir + "response_headers")
	return parseHttpHeaders(string(h))
}

func parseHttpHeaders(rawHeaders string) http.Header {
	headers := http.Header{}
	for _, s := range strings.Split(rawHeaders, "\n") {
		h := strings.SplitN(s, ":", 2)
		if len(h) > 1 {
			k := h[0]
			v := h[1]
			headers.Add(k, v)
		}
	}
	return headers
}

func addHeaders(w http.ResponseWriter, headers http.Header) {
	for k, ha := range headers {
		for _, h := range ha {
			w.Header().Add(k, h)
		}
	}
}
