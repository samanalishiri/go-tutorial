package user

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"restful-api-echo/src/main/go/user"
	"strings"
	"testing"
)

var model = user.User{
	Name: "James",
	Role: "Developer",
}

func Test1_UserSave_GivenUser_PostRequest_ThenReturnLocation(t *testing.T) {

	marshal, err := json.Marshal(model)
	checkError(t, err, "the user could not marshal")

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(marshal))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users")
	err = user.Save(c)

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.NotNil(t, rec.Header().Get("Location"))

		id := strings.TrimPrefix(rec.Header().Get("Location"), "/users/")
		model.ID = bson.ObjectIdHex(id)
	}
}

func Test2_UserGetOne_GivenIdentity_GetRequest_ThenReturnUser(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues(model.ID.Hex())
	err := user.GetOne(c)

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.NotNil(t, rec.Body)

		body, err := ioutil.ReadAll(rec.Body)
		checkError(t, err, "unmarshal response body was failed")

		var u user.User
		json.Unmarshal(body, &u)
		assert.Equal(t, "James", u.Name)
		assert.Equal(t, "Developer", u.Role)
	}
}
