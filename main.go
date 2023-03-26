package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/exp/slices"
)

var (
	stuffs []string
	port   string = ":3000"
)

func main() {
	http.HandleFunc("/api", APIController)
	fmt.Printf("API initialized. listening on %v...", port)
	http.ListenAndServe(port, nil)
}

func Get(w http.ResponseWriter) {
	encode(stuffs, w)
}

func Post(w http.ResponseWriter, r *http.Request) {
	var things []string
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&things)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	StuffAdd(things)
}

func Put(w http.ResponseWriter, r *http.Request) {}

func APIController(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/api" {
		switch r.Method {
		case http.MethodGet:
			Get(w)
		case http.MethodPost:
			Post(w, r)
		case http.MethodPut:
			Put(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func StuffAdd(things []string) {
	for _, thing := range things {
		stuffs = append(stuffs, thing)
	}
}

func StuffStripDupes(things []string) (sanitizedThings []string) {
	for _, thing := range things {
		if slices.Contains(things, thing) {
			return
		} else {
			sanitizedThings = append(sanitizedThings, thing)
		}
	}
	return
}

func encode(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}

func decode(data interface{}, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	dec.Decode(&data)
}
