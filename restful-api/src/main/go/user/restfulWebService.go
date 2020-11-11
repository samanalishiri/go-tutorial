package user

import (
	"encoding/json"
	"errors"
	"github.com/asdine/storm/v3"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"net/http"
	"restfull-api/src/main/go/web"
)

func GetBody(r *http.Request, u *User) error {
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
	users, err := FindAll()
	if err != nil {
		web.ThrowError(w, http.StatusInternalServerError)
		return
	}
	web.CreatHttpResponse(w, http.StatusOK, web.JsonResponse{"users": users})
}

func Save(w http.ResponseWriter, r *http.Request) {
	u := new(User)
	err := GetBody(r, u)
	if err != nil {
		web.ThrowError(w, http.StatusBadRequest)
		return
	}

	u.ID = bson.NewObjectId()

	err = SaveUser(u)
	if err != nil {
		if err == ErrRecordInvalid {
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
	u, err := FindById(id)
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
	u := new(User)
	err := GetBody(r, u)
	if err != nil {
		web.ThrowError(w, http.StatusBadRequest)
		return
	}
	u.ID = id
	err = SaveUser(u)
	if err != nil {
		if err == ErrRecordInvalid {
			web.ThrowError(w, http.StatusBadRequest)
		} else {
			web.ThrowError(w, http.StatusInternalServerError)
		}
		return
	}
	web.CreatHttpResponse(w, http.StatusOK, web.JsonResponse{"user": u})
}

func PatchOne(w http.ResponseWriter, r *http.Request, id bson.ObjectId) {
	u, err := FindById(id)
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
	err = SaveUser(u)
	if err != nil {
		if err == ErrRecordInvalid {
			web.ThrowError(w, http.StatusBadRequest)
		} else {
			web.ThrowError(w, http.StatusInternalServerError)
		}
		return
	}
	web.CreatHttpResponse(w, http.StatusOK, web.JsonResponse{"user": u})
}

func DeleteOne(w http.ResponseWriter, _ *http.Request, id bson.ObjectId) {
	err := DeleteById(id)
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
