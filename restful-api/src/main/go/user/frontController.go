package user

import (
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"restfull-api/src/main/go/contract"
	"restfull-api/src/main/go/utils"
	"strings"
)

type FrontController struct {
}

func NewFrontController() contract.FrontController {
	return &FrontController{}
}

var endpointMappers = make([]contract.EndpointMapper, 0, 10)

func (f *FrontController) Post(url string, function func(c contract.Context)) {
	endpointMappers = append(endpointMappers, contract.EndpointMapper{URL: url, Method: http.MethodPost, Function: function})
}

func (f *FrontController) Put(url string, function func(c contract.Context)) {
	endpointMappers = append(endpointMappers, contract.EndpointMapper{URL: url, Method: http.MethodPut, Function: function})
}

func (f *FrontController) Get(url string, function func(c contract.Context)) {
	endpointMappers = append(endpointMappers, contract.EndpointMapper{URL: url, Method: http.MethodGet, Function: function})
}

func (f *FrontController) Delete(url string, function func(c contract.Context)) {
	endpointMappers = append(endpointMappers, contract.EndpointMapper{URL: url, Method: http.MethodDelete, Function: function})
}

func (f *FrontController) Patch(url string, function func(c contract.Context)) {
	endpointMappers = append(endpointMappers, contract.EndpointMapper{URL: url, Method: http.MethodPatch, Function: function})
}

func (f *FrontController) Options(url string, function func(c contract.Context)) {
	endpointMappers = append(endpointMappers, contract.EndpointMapper{URL: url, Method: http.MethodOptions, Function: function})
}

func (f *FrontController) Head(url string, function func(c contract.Context)) {
	endpointMappers = append(endpointMappers, contract.EndpointMapper{URL: url, Method: http.MethodHead, Function: function})
}

func (f *FrontController) Endpoint(w http.ResponseWriter, r *http.Request) {

	var path string

	var id string

	if r.URL.Path == "/users" {
		path = "/users"
	} else {
		id = strings.TrimPrefix(r.URL.Path, "/users/")
		if !bson.IsObjectIdHex(id) {
			utils.ThrowError(w, http.StatusNotFound)
			return
		}
		path = "/users/:id"
	}

	for i := 0; i < len(endpointMappers); i++ {
		if endpointMappers[i].URL == path && endpointMappers[i].Method == r.Method {
			endpointMappers[i].Function(contract.Context{Writer: w, Request: r, Params: map[string]string{"id": id}})
		}
	}
}
