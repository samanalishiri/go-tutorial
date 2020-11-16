package user

import (
	"encoding/json"
	"errors"
	"github.com/asdine/storm/v3"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"net/http"
)

type Handler interface {
	SaveOne(c Context)
	GetOne(c Context)
	UpdateOne(c Context)
	DeleteOne(c Context)
	PatchOne(c Context)
	GetAll(c Context)
}

type HandlerImpl struct {
	repository Repository
}

func NewHandler() Handler {
	return &HandlerImpl{
		repository: NewUserRepository(),
	}
}

func (h *HandlerImpl) GetBody(c Context, u *User) error {
	if c.Request.Body == nil {
		return errors.New("request body is empty")
	}

	if u == nil {
		return errors.New("endpointMappers user is required")
	}

	body, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		return err
	}

	return json.Unmarshal(body, u)
}

func (h *HandlerImpl) SaveOne(c Context) {
	u := new(User)
	err := h.GetBody(c, u)
	if err != nil {
		ThrowError(c.Writer, http.StatusBadRequest)
		return
	}

	u.ID = bson.NewObjectId()

	err = h.repository.Save(u)
	if err != nil {
		if err == ErrRecordInvalid {
			ThrowError(c.Writer, http.StatusBadRequest)
		} else {
			ThrowError(c.Writer, http.StatusInternalServerError)
		}
		return
	}

	c.Writer.Header().Set("Location", "/users/"+u.ID.Hex())
	c.Writer.WriteHeader(http.StatusCreated)
}

func (h *HandlerImpl) GetOne(c Context) {
	u, err := h.repository.FindById(bson.ObjectIdHex(c.Params["id"]))
	if err != nil {
		if err == storm.ErrNotFound {
			ThrowError(c.Writer, http.StatusNotFound)
			return
		}
		ThrowError(c.Writer, http.StatusInternalServerError)
		return
	}
	CreatHttpResponse(c, http.StatusOK, u)
}

func (h *HandlerImpl) UpdateOne(c Context) {
	u := new(User)
	err := h.GetBody(c, u)
	if err != nil {
		ThrowError(c.Writer, http.StatusBadRequest)
		return
	}
	u.ID = bson.ObjectIdHex(c.Params["id"])
	err = h.repository.Save(u)
	if err != nil {
		if err == ErrRecordInvalid {
			ThrowError(c.Writer, http.StatusBadRequest)
		} else {
			ThrowError(c.Writer, http.StatusInternalServerError)
		}
		return
	}
	c.Writer.WriteHeader(http.StatusNoContent)
}

func (h *HandlerImpl) DeleteOne(c Context) {
	err := h.repository.DeleteById(bson.ObjectIdHex(c.Params["id"]))
	if err != nil {
		if err == storm.ErrNotFound {
			ThrowError(c.Writer, http.StatusNotFound)
			return
		}
		ThrowError(c.Writer, http.StatusInternalServerError)
		return
	}
	c.Writer.WriteHeader(http.StatusNoContent)
}

func (h *HandlerImpl) PatchOne(c Context) {
	u, err := h.repository.FindById(bson.ObjectIdHex(c.Params["id"]))
	if err != nil {
		if err == storm.ErrNotFound {
			ThrowError(c.Writer, http.StatusNotFound)
			return
		}
		ThrowError(c.Writer, http.StatusInternalServerError)
		return
	}
	err = h.GetBody(c, u)
	if err != nil {
		ThrowError(c.Writer, http.StatusBadRequest)
		return
	}
	u.ID = bson.ObjectIdHex(c.Params["id"])
	err = h.repository.Save(u)
	if err != nil {
		if err == ErrRecordInvalid {
			ThrowError(c.Writer, http.StatusBadRequest)
		} else {
			ThrowError(c.Writer, http.StatusInternalServerError)
		}
		return
	}
	c.Writer.WriteHeader(http.StatusNoContent)
}

func (h *HandlerImpl) GetAll(c Context) {
	users, err := h.repository.FindAll()
	if err != nil {
		ThrowError(c.Writer, http.StatusInternalServerError)
		return
	}
	CreatHttpResponse(c, http.StatusOK, users)
}
