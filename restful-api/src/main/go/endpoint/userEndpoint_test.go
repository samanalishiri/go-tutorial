package endpoint

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

type Response struct {
	User user.User `json:"user"`
}

var u = &user.User{
	Name: "James",
	Role: "Developer",
}

func Test1_UserGetOne_GivenIdentity_GetRequest_ThenReturnUser(t *testing.T) {

	marshal, err := json.Marshal(u)
	checkError(t, err, "the user could not marshal")

	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(marshal))
	checkError(t, err, "create http POST request was failed")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UserEndpoint)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.NotNil(t, rr.Header().Get("Location"))

	id := strings.TrimPrefix(rr.Header().Get("Location"), "/users/")
	u.ID = bson.ObjectIdHex(id)
	t.Log(u.ID)
}

func Test2_UserGetOne_GivenIdentity_GetRequest_ThenReturnUser(t *testing.T) {
	req, err := http.NewRequest("GET", "/users/"+u.ID.Hex(), nil)
	checkError(t, err, "create http GET request was failed")
	req.Header.Add("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(SubUserEndpoint)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.NotNil(t, rr.Body)

	body, err := ioutil.ReadAll(rr.Body)
	checkError(t, err, "unmarshal response body was failed")

	var res Response
	json.Unmarshal(body, &res)

	assert.Equal(t, "James", res.User.Name)
	assert.Equal(t, "Developer", res.User.Role)
}

func checkError(t *testing.T, err error, text string) {
	if err != nil {
		t.Fatalf("%s: %s", text, err)
	}
}
