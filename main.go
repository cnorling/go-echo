package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://wttr.in")
	if err != nil {
		fmt.Printf("got an error, %v", err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.ContentLength)
}
