package contract

import "net/http"

type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
	Params  map[string]string
}

type EndpointMapper struct {
	Method   string
	URL      string
	Function func(c Context)
}

type FrontController interface {
	Options(url string, function func(c Context))
	Patch(url string, function func(c Context))
	Delete(url string, function func(c Context))
	Get(url string, function func(c Context))
	Put(url string, function func(c Context))
	Post(url string, function func(c Context))
	Head(url string, function func(c Context))
	Endpoint(w http.ResponseWriter, r *http.Request)
}
