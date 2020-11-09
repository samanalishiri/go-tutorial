package handler

import (
	"encoding/json"
	"errors"
	"github.com/asdine/storm/v3"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"net/http"
	"restfull-api/src/main/go/user"
)

func extractUser(request *http.Request, u *user.User) error {
	if request.Body == nil {
		return errors.New("request body is empty")
	}

	if u == nil {
		return errors.New("a user is required")
	}

	body, err := ioutil.ReadAll(request.Body)

	if err != nil {
		return err
	}

	return json.Unmarshal(body, u)
}

func UserGetAll(response http.ResponseWriter, request *http.Request) {
	users, err := user.All()
	if err != nil {
		postError(response, http.StatusInternalServerError)
		return
	}
	postBodyResponse(response, http.StatusOK, jsonResponse{"users": users})
}

func UserSave(response http.ResponseWriter, request *http.Request) {
	u := new(user.User)
	err := extractUser(request, u)
	if err != nil {
		postError(response, http.StatusBadRequest)
		return
	}

	u.ID = bson.NewObjectId()

	err = u.Save()
	if err != nil {
		if err == user.ErrRecordInvalid {
			postError(response, http.StatusBadRequest)
		} else {
			postError(response, http.StatusInternalServerError)
		}
		return
	}

	response.Header().Set("Location", "/users/"+u.ID.Hex())
	response.WriteHeader(http.StatusCreated)
}

func UserGetOne(response http.ResponseWriter, _ *http.Request, id bson.ObjectId) {
	u, err := user.One(id)
	if err != nil {
		if err == storm.ErrNotFound {
			postError(response, http.StatusNotFound)
			return
		}
		postError(response, http.StatusInternalServerError)
		return
	}
	postBodyResponse(response, http.StatusOK, jsonResponse{"user": u})
}

func usersPutOne(w http.ResponseWriter, r *http.Request, id bson.ObjectId) {
	u := new(user.User)
	err := extractUser(r, u)
	if err != nil {
		postError(w, http.StatusBadRequest)
		return
	}
	u.ID = id
	err = u.Save()
	if err != nil {
		if err == user.ErrRecordInvalid {
			postError(w, http.StatusBadRequest)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"user": u})
}

func usersPatchOne(w http.ResponseWriter, r *http.Request, id bson.ObjectId) {
	u, err := user.One(id)
	if err != nil {
		if err == storm.ErrNotFound {
			postError(w, http.StatusNotFound)
			return
		}
		postError(w, http.StatusInternalServerError)
		return
	}
	err = extractUser(r, u)
	if err != nil {
		postError(w, http.StatusBadRequest)
		return
	}
	u.ID = id
	err = u.Save()
	if err != nil {
		if err == user.ErrRecordInvalid {
			postError(w, http.StatusBadRequest)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"user": u})
}

func UserDelete(response http.ResponseWriter, _ *http.Request, id bson.ObjectId) {
	err := user.Delete(id)
	if err != nil {
		if err == storm.ErrNotFound {
			postError(response, http.StatusNotFound)
			return
		}
		postError(response, http.StatusInternalServerError)
		return
	}
	response.WriteHeader(http.StatusOK)
}
