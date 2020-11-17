package root

import (
	"net/http"
	"restfull-api/src/main/go/contract"
)

type FrontController struct {
}

func NewFrontController() contract.FrontController {
	return &FrontController{}
}

var endpointMappers = make([]contract.EndpointMapper, 10)

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

	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("The URL is invalid"))
		return
	}

	var path = "/"

	for i := 0; i < len(endpointMappers); i++ {
		if endpointMappers[i].URL == path && endpointMappers[i].Method == r.Method {
			endpointMappers[i].Function(contract.Context{Writer: w, Request: r})
		}
	}
}
