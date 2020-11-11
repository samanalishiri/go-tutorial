package endpoint

import (
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"restfull-api/src/main/go/userService"
	"restfull-api/src/main/go/web"
	"strings"
)

func UserEndpoint(w http.ResponseWriter, r *http.Request) {
	log.Println("URL: " + r.URL.Path)

	switch r.Method {

	case http.MethodGet:
		userService.GetAll(w, r)
		return
	case http.MethodPost:
		userService.Save(w, r)
		return
	case http.MethodHead:
		userService.GetAll(w, r)
		return
	case http.MethodOptions:
		web.CreateOptionsResponse(w, []string{http.MethodGet, http.MethodPost, http.MethodHead, http.MethodOptions}, nil)
		return
	default:
		web.ThrowError(w, http.StatusMethodNotAllowed)
		return
	}
}

func SubUserEndpoint(w http.ResponseWriter, r *http.Request) {
	log.Println("URL: " + r.URL.Path)

	pathVariable := strings.TrimPrefix(r.URL.Path, "/users/")
	if !bson.IsObjectIdHex(pathVariable) {
		web.ThrowError(w, http.StatusNotFound)
		return
	}

	id := bson.ObjectIdHex(pathVariable)
	switch r.Method {
	case http.MethodGet:
		userService.GetOne(w, r, id)
		return
	case http.MethodPut:
		userService.UpdateOne(w, r, id)
		return
	case http.MethodPatch:
		userService.UsersPatchOne(w, r, id)
		return
	case http.MethodDelete:
		userService.DeleteOne(w, r, id)
		return
	case http.MethodHead:
		userService.GetOne(w, r, id)
		return
	case http.MethodOptions:
		web.CreateOptionsResponse(w, []string{http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodHead, http.MethodOptions}, nil)
		return
	default:
		web.ThrowError(w, http.StatusMethodNotAllowed)
		return
	}
}
