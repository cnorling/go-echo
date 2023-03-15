package main

import (
	"fmt"
	"net/http"
)

func hello(writer http.ResponseWriter, requester *http.Request) {
	fmt.Fprintf(writer, "hello\n")
}

func headers(writer http.ResponseWriter, requester *http.Request) {
	for name, headers := range requester.Header {
		for _, h := range headers {
			fmt.Fprintf(writer, "%v: %v\n", name, h)
		}
	}
}

func post() {
	http.HandleFunc("/headers", headers)
	http.ListenAndServe(":8090", nil)
}
