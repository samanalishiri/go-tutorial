package user

import (
	"encoding/json"
	"errors"
	"github.com/asdine/storm/v3"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"net/http"
	"restfull-api/src/main/go/contract"
	"restfull-api/src/main/go/utils"
)

type HandlerImpl struct {
	repository Repository
}

func NewHandler() contract.Handler {
	return &HandlerImpl{
		repository: NewUserRepository(),
	}
}

func (h *HandlerImpl) GetBody(c contract.Context, u *User) error {
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

func (h *HandlerImpl) SaveOne(c contract.Context) {
	u := new(User)
	err := h.GetBody(c, u)
	if err != nil {
		utils.ThrowError(c.Writer, http.StatusBadRequest)
		return
	}

	u.ID = bson.NewObjectId()

	err = h.repository.Save(u)
	if err != nil {
		if err == ErrRecordInvalid {
			utils.ThrowError(c.Writer, http.StatusBadRequest)
		} else {
			utils.ThrowError(c.Writer, http.StatusInternalServerError)
		}
		return
	}

	c.Writer.Header().Set("Location", "/users/"+u.ID.Hex())
	c.Writer.WriteHeader(http.StatusCreated)
}

func (h *HandlerImpl) GetOne(c contract.Context) {
	u, err := h.repository.FindById(bson.ObjectIdHex(c.Params["id"]))
	if err != nil {
		if err == storm.ErrNotFound {
			utils.ThrowError(c.Writer, http.StatusNotFound)
			return
		}
		utils.ThrowError(c.Writer, http.StatusInternalServerError)
		return
	}
	utils.CreatHttpResponse(c, http.StatusOK, u)
}

func (h *HandlerImpl) UpdateOne(c contract.Context) {
	u := new(User)
	err := h.GetBody(c, u)
	if err != nil {
		utils.ThrowError(c.Writer, http.StatusBadRequest)
		return
	}
	u.ID = bson.ObjectIdHex(c.Params["id"])
	err = h.repository.Save(u)
	if err != nil {
		if err == ErrRecordInvalid {
			utils.ThrowError(c.Writer, http.StatusBadRequest)
		} else {
			utils.ThrowError(c.Writer, http.StatusInternalServerError)
		}
		return
	}
	c.Writer.WriteHeader(http.StatusNoContent)
}

func (h *HandlerImpl) DeleteOne(c contract.Context) {
	err := h.repository.DeleteById(bson.ObjectIdHex(c.Params["id"]))
	if err != nil {
		if err == storm.ErrNotFound {
			utils.ThrowError(c.Writer, http.StatusNotFound)
			return
		}
		utils.ThrowError(c.Writer, http.StatusInternalServerError)
		return
	}
	c.Writer.WriteHeader(http.StatusNoContent)
}

func (h *HandlerImpl) PatchOne(c contract.Context) {
	u, err := h.repository.FindById(bson.ObjectIdHex(c.Params["id"]))
	if err != nil {
		if err == storm.ErrNotFound {
			utils.ThrowError(c.Writer, http.StatusNotFound)
			return
		}
		utils.ThrowError(c.Writer, http.StatusInternalServerError)
		return
	}
	err = h.GetBody(c, u)
	if err != nil {
		utils.ThrowError(c.Writer, http.StatusBadRequest)
		return
	}
	u.ID = bson.ObjectIdHex(c.Params["id"])
	err = h.repository.Save(u)
	if err != nil {
		if err == ErrRecordInvalid {
			utils.ThrowError(c.Writer, http.StatusBadRequest)
		} else {
			utils.ThrowError(c.Writer, http.StatusInternalServerError)
		}
		return
	}
	c.Writer.WriteHeader(http.StatusNoContent)
}

func (h *HandlerImpl) GetAll(c contract.Context) {
	users, err := h.repository.FindAll()
	if err != nil {
		utils.ThrowError(c.Writer, http.StatusInternalServerError)
		return
	}
	utils.CreatHttpResponse(c, http.StatusOK, users)
}

func (h *HandlerImpl) GetBasicMethod(c contract.Context) {
	utils.CreateOptionsResponse(c, []string{http.MethodGet, http.MethodPost, http.MethodHead, http.MethodOptions}, nil)
}

func (h *HandlerImpl) GetAllMethod(c contract.Context) {
	utils.CreateOptionsResponse(c, []string{http.MethodGet, http.MethodPost, http.MethodHead, http.MethodOptions}, nil)
}
