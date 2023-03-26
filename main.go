package main

import (
	"fmt"
	"net/http"

	"api/api"
)

var (
	port string = ":3000"
)

func main() {
	http.HandleFunc("/api", api.Controller)
	fmt.Printf("API initialized. listening on %v...", port)
	http.ListenAndServe(port, nil)
}
