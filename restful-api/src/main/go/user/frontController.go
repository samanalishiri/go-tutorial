package user

import (
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strings"
)

type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
	Params  map[string]string
}

type EndpointMapper struct {
	method   string
	url      string
	function func(c Context)
}

type FrontController struct {
}

func NewFrontController() FrontController {
	return FrontController{}
}

var endpointMappers = make([]EndpointMapper, 10)

func (f *FrontController) Post(url string, function func(c Context)) {
	endpointMappers = append(endpointMappers, EndpointMapper{url: url, method: http.MethodPost, function: function})
}

func (f *FrontController) Put(url string, function func(c Context)) {
	endpointMappers = append(endpointMappers, EndpointMapper{url: url, method: http.MethodPut, function: function})
}

func (f *FrontController) Get(url string, function func(c Context)) {
	endpointMappers = append(endpointMappers, EndpointMapper{url: url, method: http.MethodGet, function: function})
}

func (f *FrontController) Delete(url string, function func(c Context)) {
	endpointMappers = append(endpointMappers, EndpointMapper{url: url, method: http.MethodDelete, function: function})
}

func (f *FrontController) Patch(url string, function func(c Context)) {
	endpointMappers = append(endpointMappers, EndpointMapper{url: url, method: http.MethodPatch, function: function})
}

func (f *FrontController) Options(url string, function func(c Context)) {
	endpointMappers = append(endpointMappers, EndpointMapper{url: url, method: http.MethodOptions, function: function})
}

func (f *FrontController) Head(url string, function func(c Context)) {
	endpointMappers = append(endpointMappers, EndpointMapper{url: url, method: http.MethodHead, function: function})
}

func (f *FrontController) Dispatcher(w http.ResponseWriter, r *http.Request) {

	var path string

	var id string

	if r.URL.Path == "/users" {
		path = "/users"
	} else {
		id = strings.TrimPrefix(r.URL.Path, "/users/")
		if !bson.IsObjectIdHex(id) {
			ThrowError(w, http.StatusNotFound)
			return
		}
		path = "/users/:id"
	}

	for k := 0; k < len(endpointMappers); k++ {
		if endpointMappers[k].url == path && endpointMappers[k].method == r.Method {
			endpointMappers[k].function(Context{Writer: w, Request: r, Params: map[string]string{"id": id}})
		}
	}
}
