package main

import (
	"encoding/json"
	"net/http"
	"regexp"

	"golang.org/x/exp/slices"
)

type Controller struct {
	path *regexp.Regexp
}

var (
	stuffs []string
)

func main() {
	RegisterControllers()
	http.ListenAndServe(":3000", nil)
}

func Get(w http.ResponseWriter, r *http.Request) {}

func Post(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	var things []string
	err := dec.Decode(&things)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	StuffAdd(things)
}

func Put(w http.ResponseWriter, r *http.Request) {}

func newAPIController() *Controller {
	return &Controller{
		path: regexp.MustCompile(`^/api`),
	}
}

func RegisterControllers() {
	http.HandleFunc("/api", APIController)
}

func APIController(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/api" {
		switch r.Method {
		case http.MethodGet:
			Get(w, r)
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
