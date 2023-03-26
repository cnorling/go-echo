package api

import (
	"encoding/json"
	"io"
	"net/http"

	"golang.org/x/exp/slices"
)

var (
	stuffs []string
)

func get(w http.ResponseWriter) {
	encode(stuffs, w)
}

func post(w http.ResponseWriter, r *http.Request) {
	var things []string
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&things)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	stuffAdd(things)
}

func put(w http.ResponseWriter, r *http.Request) {}

func Controller(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/api" {
		switch r.Method {
		case http.MethodGet:
			get(w)
		case http.MethodPost:
			post(w, r)
		case http.MethodPut:
			put(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func stuffAdd(things []string) {
	for _, thing := range things {
		stuffs = append(stuffs, thing)
	}
}

func stuffStripDupes(things []string) (sanitizedThings []string) {
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
