package user

import (
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"restfull-api/src/main/go/web"
	"strings"
)

func Endpoint(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		GetAll(w, r)
		return
	case http.MethodPost:
		Save(w, r)
		return
	case http.MethodHead:
		GetAll(w, r)
		return
	case http.MethodOptions:
		web.CreateOptionsResponse(w, []string{http.MethodGet, http.MethodPost, http.MethodHead, http.MethodOptions}, nil)
		return
	default:
		web.ThrowError(w, http.StatusMethodNotAllowed)
		return
	}
}

func FirstLevelEndpoint(w http.ResponseWriter, r *http.Request) {

	pathVariable := strings.TrimPrefix(r.URL.Path, "/users/")
	if !bson.IsObjectIdHex(pathVariable) {
		web.ThrowError(w, http.StatusNotFound)
		return
	}

	id := bson.ObjectIdHex(pathVariable)
	switch r.Method {
	case http.MethodGet:
		GetOne(w, r, id)
		return
	case http.MethodPut:
		UpdateOne(w, r, id)
		return
	case http.MethodPatch:
		PatchOne(w, r, id)
		return
	case http.MethodDelete:
		DeleteOne(w, r, id)
		return
	case http.MethodHead:
		GetOne(w, r, id)
		return
	case http.MethodOptions:
		web.CreateOptionsResponse(w, []string{http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodDelete,
			http.MethodHead, http.MethodOptions}, nil)
		return
	default:
		web.ThrowError(w, http.StatusMethodNotAllowed)
		return
	}
}
