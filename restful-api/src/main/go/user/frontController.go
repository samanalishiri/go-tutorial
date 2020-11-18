package user

import (
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"restfull-api/src/main/go/contract"
	"restfull-api/src/main/go/utils"
	"strings"
)

type FrontController struct {
	contract.FrontControllerImpl
}

func NewFrontController() contract.FrontController {
	return &FrontController{
		contract.FrontControllerImpl{
			EndpointMappers: make([]contract.EndpointMapper, 0, 10),
		},
	}
}

func (f *FrontController) Route(w http.ResponseWriter, r *http.Request) {

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

	for i := 0; i < len(f.EndpointMappers); i++ {
		if f.EndpointMappers[i].URL == path && f.EndpointMappers[i].Method == r.Method {
			f.EndpointMappers[i].Function(contract.Context{Writer: w, Request: r, Params: map[string]string{"id": id}})
		}
	}
}
