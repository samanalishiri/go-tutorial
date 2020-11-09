package handler

import (
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strings"
)

func UserRouter(response http.ResponseWriter, request *http.Request) {
	path := strings.TrimPrefix(request.URL.Path, "/")

	if path == "users" {
		switch request.Method {
		case http.MethodGet:
			UserGetAll(response, request)
			return
		case http.MethodPost:
			UserSave(response, request)
			return
		case http.MethodHead:
			UserGetAll(response, request)
			return
		case http.MethodOptions:
			postOptionsResponse(response, []string{http.MethodGet, http.MethodPost, http.MethodHead, http.MethodOptions}, nil)
			return
		default:
			postError(response, http.StatusMethodNotAllowed)
			return
		}
	}

	path = strings.TrimPrefix(request.URL.Path, "/users/")

	if !bson.IsObjectIdHex(path) {
		postError(response, http.StatusNotFound)
		return
	}

	id := bson.ObjectIdHex(path)

	switch request.Method {
	case http.MethodGet:
		UserGetOne(response, request, id)
		return
	case http.MethodPut:
		usersPutOne(response, request, id)
		return
	case http.MethodPatch:
		usersPatchOne(response, request, id)
		return
	case http.MethodDelete:
		UserDelete(response, request, id)
		return
	case http.MethodHead:
		UserGetOne(response, request, id)
		return
	case http.MethodOptions:
		postOptionsResponse(response, []string{http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodHead, http.MethodOptions}, nil)
		return
	default:
		postError(response, http.StatusMethodNotAllowed)
		return
	}
}
