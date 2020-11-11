package user

import (
	"github.com/asdine/storm/v3"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2/bson"
	"restfull-api/src/main/go/user"
	"testing"
)

var u = &user.User{
	Name: "James",
	Role: "Developer",
}

func Test1_Save_GivenNewUser_WhenSave_ReturnID(t *testing.T) {
	u.ID = bson.NewObjectId()
	err := user.SaveUser(u)
	checkError(t, err, "save user was failed")
	assert.NotNil(t, u.ID)
}

func Test2_One_GivenID_WhenReadOne_ThenReturnUser(t *testing.T) {
	u, err := user.FindById(u.ID)
	checkError(t, err, "read user was failed")
	assert.Equal(t, "James", u.Name)
	assert.Equal(t, "Developer", u.Role)
}

func Test3_Update_GivenNewChanges_WhenUpdate_ThenApplyTheChanges(t *testing.T) {
	u.Role = "Team Lead"
	err := user.SaveUser(u)
	checkError(t, err, "update user was failed")
	u2, err := user.FindById(u.ID)
	checkError(t, err, "read user was failed")
	assert.Equal(t, u2.Name, "James")
	assert.Equal(t, u2.Role, "Team Lead")
}

func Test4_Delete_GivenId_WhenDelete_ThenRemoveFromDatabase(t *testing.T) {
	err := user.DeleteById(u.ID)
	checkError(t, err, "delete user was failed")
	_, err = user.FindById(u.ID)
	assert.Equal(t, err, storm.ErrNotFound)
}
