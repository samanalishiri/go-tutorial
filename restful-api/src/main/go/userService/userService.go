package userService

import (
	"encoding/json"
	"errors"
	"github.com/asdine/storm/v3"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"net/http"
	"restfull-api/src/main/go/user"
	"restfull-api/src/main/go/web"
)

func GetBody(r *http.Request, u *user.User) error {
	if r.Body == nil {
		return errors.New("r body is empty")
	}

	if u == nil {
		return errors.New("a user is required")
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return err
	}

	return json.Unmarshal(body, u)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := user.All()
	if err != nil {
		web.ThrowError(w, http.StatusInternalServerError)
		return
	}
	web.CreatHttpResponse(w, http.StatusOK, web.JsonResponse{"users": users})
}

func Save(w http.ResponseWriter, r *http.Request) {
	u := new(user.User)
	err := GetBody(r, u)
	if err != nil {
		web.ThrowError(w, http.StatusBadRequest)
		return
	}

	u.ID = bson.NewObjectId()

	err = u.Save()
	if err != nil {
		if err == user.ErrRecordInvalid {
			web.ThrowError(w, http.StatusBadRequest)
		} else {
			web.ThrowError(w, http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Location", "/users/"+u.ID.Hex())
	w.WriteHeader(http.StatusCreated)
}

func GetOne(w http.ResponseWriter, _ *http.Request, id bson.ObjectId) {
	u, err := user.One(id)
	if err != nil {
		if err == storm.ErrNotFound {
			web.ThrowError(w, http.StatusNotFound)
			return
		}
		web.ThrowError(w, http.StatusInternalServerError)
		return
	}
	web.CreatHttpResponse(w, http.StatusOK, web.JsonResponse{"user": u})
}

func UpdateOne(w http.ResponseWriter, r *http.Request, id bson.ObjectId) {
	u := new(user.User)
	err := GetBody(r, u)
	if err != nil {
		web.ThrowError(w, http.StatusBadRequest)
		return
	}
	u.ID = id
	err = u.Save()
	if err != nil {
		if err == user.ErrRecordInvalid {
			web.ThrowError(w, http.StatusBadRequest)
		} else {
			web.ThrowError(w, http.StatusInternalServerError)
		}
		return
	}
	web.CreatHttpResponse(w, http.StatusOK, web.JsonResponse{"user": u})
}

func UsersPatchOne(w http.ResponseWriter, r *http.Request, id bson.ObjectId) {
	u, err := user.One(id)
	if err != nil {
		if err == storm.ErrNotFound {
			web.ThrowError(w, http.StatusNotFound)
			return
		}
		web.ThrowError(w, http.StatusInternalServerError)
		return
	}
	err = GetBody(r, u)
	if err != nil {
		web.ThrowError(w, http.StatusBadRequest)
		return
	}
	u.ID = id
	err = u.Save()
	if err != nil {
		if err == user.ErrRecordInvalid {
			web.ThrowError(w, http.StatusBadRequest)
		} else {
			web.ThrowError(w, http.StatusInternalServerError)
		}
		return
	}
	web.CreatHttpResponse(w, http.StatusOK, web.JsonResponse{"user": u})
}

func DeleteOne(w http.ResponseWriter, _ *http.Request, id bson.ObjectId) {
	err := user.Delete(id)
	if err != nil {
		if err == storm.ErrNotFound {
			web.ThrowError(w, http.StatusNotFound)
			return
		}
		web.ThrowError(w, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
