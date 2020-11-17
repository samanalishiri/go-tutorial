package user

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"restfull-api/src/main/go/user"
	"strings"
	"testing"
)

var (
	model = user.User{
		Name: "James",
		Role: "Developer",
	}

	dispatcher = user.NewDispatcher().Init()
)

func Test1_UserSave_GivenUser_PostRequest_ThenReturnLocation(t *testing.T) {

	marshal, err := json.Marshal(model)
	checkError(t, err, "the user could not marshal")

	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(marshal))
	checkError(t, err, "create http POST request was failed")
	req.Header.Add("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(dispatcher.Endpoint)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.NotNil(t, rr.Header().Get("Location"))

	id := strings.TrimPrefix(rr.Header().Get("Location"), "/users/")
	model.ID = bson.ObjectIdHex(id)
}

func Test2_UserGetOne_GivenIdentity_GetRequest_ThenReturnUser(t *testing.T) {
	req, err := http.NewRequest("GET", "/users/"+model.ID.Hex(), nil)
	checkError(t, err, "create http GET request was failed")
	req.Header.Add("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(dispatcher.Endpoint)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.NotNil(t, rr.Body)

	body, err := ioutil.ReadAll(rr.Body)
	checkError(t, err, "reading body was failed")

	var u user.User
	err = json.Unmarshal(body, &u)
	checkError(t, err, "unmarshal response body was failed")

	assert.Equal(t, "James", u.Name)
	assert.Equal(t, "Developer", u.Role)
}

func Test3_UserUpdate_GivenIdentityAndUser_PutRequest_ThenReturnUser(t *testing.T) {
	model.Role = "Team Lead"
	marshal, err := json.Marshal(model)
	checkError(t, err, "the user could not marshal")

	req, err := http.NewRequest("PUT", "/users/"+model.ID.Hex(), bytes.NewReader(marshal))
	checkError(t, err, "create http PUT request was failed")
	req.Header.Add("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(dispatcher.Endpoint)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)
	assert.Nil(t, rr.Body.Bytes())

	req2, err2 := http.NewRequest("GET", "/users/"+model.ID.Hex(), nil)
	checkError(t, err2, "create http GET request was failed")
	req2.Header.Add("Content-Type", "application/json")

	rr2 := httptest.NewRecorder()
	handler2 := http.HandlerFunc(dispatcher.Endpoint)
	handler2.ServeHTTP(rr2, req2)

	assert.Equal(t, http.StatusOK, rr2.Code)
	assert.NotNil(t, rr2.Body)

	body, err2 := ioutil.ReadAll(rr2.Body)
	checkError(t, err2, "unmarshal response body was failed")

	var u user.User
	err = json.Unmarshal(body, &u)
	checkError(t, err, "unmarshal response body was failed")

	assert.Equal(t, "James", u.Name)
	assert.Equal(t, "Team Lead", u.Role)
}
