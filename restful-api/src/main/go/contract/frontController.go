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
	Route(w http.ResponseWriter, r *http.Request)
}

type FrontControllerImpl struct {
	EndpointMappers []EndpointMapper
}

func (f *FrontControllerImpl) Route(w http.ResponseWriter, r *http.Request) {

}

func (f *FrontControllerImpl) Post(url string, function func(c Context)) {
	f.EndpointMappers = append(f.EndpointMappers, EndpointMapper{URL: url, Method: http.MethodPost, Function: function})
}

func (f *FrontControllerImpl) Put(url string, function func(c Context)) {
	f.EndpointMappers = append(f.EndpointMappers, EndpointMapper{URL: url, Method: http.MethodPut, Function: function})
}

func (f *FrontControllerImpl) Get(url string, function func(c Context)) {
	f.EndpointMappers = append(f.EndpointMappers, EndpointMapper{URL: url, Method: http.MethodGet, Function: function})
}

func (f *FrontControllerImpl) Delete(url string, function func(c Context)) {
	f.EndpointMappers = append(f.EndpointMappers, EndpointMapper{URL: url, Method: http.MethodDelete, Function: function})
}

func (f *FrontControllerImpl) Patch(url string, function func(c Context)) {
	f.EndpointMappers = append(f.EndpointMappers, EndpointMapper{URL: url, Method: http.MethodPatch, Function: function})
}

func (f *FrontControllerImpl) Options(url string, function func(c Context)) {
	f.EndpointMappers = append(f.EndpointMappers, EndpointMapper{URL: url, Method: http.MethodOptions, Function: function})
}

func (f *FrontControllerImpl) Head(url string, function func(c Context)) {
	f.EndpointMappers = append(f.EndpointMappers, EndpointMapper{URL: url, Method: http.MethodHead, Function: function})
}
