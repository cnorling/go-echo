package main

import (
	"fmt"
	"net/http"
)

func main() {
	initialize()
	http.ListenAndServe(":8080", nil)
}

type api struct {
	input string
	port  string
}

func initialize() {
	fmt.Println("initializing api")
}

func (api api) get() {}

func (api api) post() {}
