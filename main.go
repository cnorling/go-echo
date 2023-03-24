package main

import (
	"net/http"
	"regexp"
)

type Stuff struct {
	things []string
}

type Controller struct {
	path *regexp.Regexp
}

func main() {
	RegisterControllers()
	http.ListenAndServe(":8080", nil)
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

func (c Controller) Get(w http.ResponseWriter, r *http.Request)  {}
func (c Controller) Post(w http.ResponseWriter, r *http.Request) {}
func (c Controller) Put(w http.ResponseWriter, r *http.Request)  {}

func newController() *Controller {
	return &Controller{
		path: regexp.MustCompile(`^/api`),
	}
}

func RegisterControllers() {
	c := newController()
	http.Handle("/api", c)
}
