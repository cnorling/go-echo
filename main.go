package main

import (
	"encoding/json"
	"net/http"
	"regexp"
)

type Controller struct {
	path *regexp.Regexp
}

var (
	stuff []string
)

func main() {
	RegisterControllers()
	http.ListenAndServe(":3000", nil)
}

func (c Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/api" {
		switch r.Method {
		case http.MethodGet:
			c.Get(w, r)
		case http.MethodPost:
			c.Post(w, r)
		case http.MethodPut:
			c.Put(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (c Controller) Get(w http.ResponseWriter, r *http.Request) {}
func (c Controller) Post(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	var things []string
	err := dec.Decode(&things)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	AddStuff(things)
}
func (c Controller) Put(w http.ResponseWriter, r *http.Request) {}

func newAPIController() *Controller {
	return &Controller{
		path: regexp.MustCompile(`^/api`),
	}
}

func RegisterControllers() {
	c := newAPIController()
	http.Handle("/api", c)
}

func AddStuff(things []string) {
	for _, thing := range things {
		stuff = append(stuff, thing)
	}
}
