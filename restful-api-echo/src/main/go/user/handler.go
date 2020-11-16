package user

import (
	"encoding/json"
	"errors"
	"github.com/asdine/storm/v3"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"net/http"
)

var repository = NewUserRepository()

func GetBody(r *http.Request, u *User) error {
	if r.Body == nil {
		return errors.New("r body is empty")
	}

	if u == nil {
		return errors.New("requestMap user is required")
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return err
	}

	return json.Unmarshal(body, u)
}

func Save(c echo.Context) error {
	u := new(User)
	err := GetBody(c.Request(), u)
	if err != nil {
		ThrowError(c, http.StatusBadRequest, "")
		return nil
	}

	u.ID = bson.NewObjectId()

	err = repository.SaveUser(u)
	if err != nil {
		if err == ErrRecordInvalid {
			ThrowError(c, http.StatusBadRequest, "")
		} else {
			ThrowError(c, http.StatusInternalServerError, "")
		}
		return nil
	}

	c.Response().Header().Set("Location", "/users/"+u.ID.Hex())
	c.Response().WriteHeader(http.StatusCreated)
	return nil
}

func GetOne(c echo.Context) error {
	u, err := repository.FindById(bson.ObjectIdHex(c.Param("id")))
	if err != nil {
		if err == storm.ErrNotFound {
			ThrowError(c, http.StatusNotFound, "")
			return nil
		}
		ThrowError(c, http.StatusInternalServerError, "")
		return nil
	}
	CreatHttpResponse(c, http.StatusOK, u)
	return nil
}

func UpdateOne(c echo.Context) error {
	u := new(User)
	err := GetBody(c.Request(), u)
	if err != nil {
		ThrowError(c, http.StatusBadRequest, "")
		return nil
	}
	u.ID = bson.ObjectIdHex(c.Param("id"))
	err = repository.SaveUser(u)
	if err != nil {
		if err == ErrRecordInvalid {
			ThrowError(c, http.StatusBadRequest, "")
		} else {
			ThrowError(c, http.StatusInternalServerError, "")
		}
		return nil
	}
	c.Response().WriteHeader(http.StatusNoContent)
	return nil
}

func DeleteOne(c echo.Context) error {
	err := repository.DeleteById(bson.ObjectIdHex(c.Param("id")))
	if err != nil {
		if err == storm.ErrNotFound {
			ThrowError(c, http.StatusNotFound, "")
			return nil
		}
		ThrowError(c, http.StatusInternalServerError, "")
		return nil
	}
	c.Response().WriteHeader(http.StatusNoContent)
	return nil
}

func PatchOne(c echo.Context) {
	u, err := repository.FindById(bson.ObjectIdHex(c.Param("id")))
	if err != nil {
		if err == storm.ErrNotFound {
			ThrowError(c, http.StatusNotFound, "")
			return
		}
		ThrowError(c, http.StatusInternalServerError, "")
		return
	}
	err = GetBody(c.Request(), u)
	if err != nil {
		ThrowError(c, http.StatusBadRequest, "")
		return
	}
	u.ID = bson.ObjectIdHex(c.Param("id"))
	err = repository.SaveUser(u)
	if err != nil {
		if err == ErrRecordInvalid {
			ThrowError(c, http.StatusBadRequest, "")
		} else {
			ThrowError(c, http.StatusInternalServerError, "")
		}
		return
	}
	c.Response().WriteHeader(http.StatusNoContent)
}

func GetAll(c echo.Context) {
	users, err := repository.FindAll()
	if err != nil {
		ThrowError(c, http.StatusInternalServerError, "")
		return
	}
	CreatHttpResponse(c, http.StatusOK, users)
}
