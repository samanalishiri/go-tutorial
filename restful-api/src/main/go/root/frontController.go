package root

import (
	"net/http"
	"restfull-api/src/main/go/contract"
)

type FrontController struct {
	contract.FrontControllerImpl
}

func NewFrontController() contract.FrontController {
	return &FrontController{
		contract.FrontControllerImpl{EndpointMappers: make([]contract.EndpointMapper, 0, 10)},
	}
}

func (f *FrontController) Route(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("The URL is invalid"))
		return
	}

	var path = "/"

	for i := 0; i < len(f.EndpointMappers); i++ {
		if f.EndpointMappers[i].URL == path && f.EndpointMappers[i].Method == r.Method {
			f.EndpointMappers[i].Function(contract.Context{Writer: w, Request: r})
		}
	}
}
