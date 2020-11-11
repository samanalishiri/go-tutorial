package user

import (
	"github.com/asdine/storm/v3"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

var u = &User{
	Name: "James",
	Role: "Developer",
}

func Test1_Save_GivenNewUser_WhenSave_ReturnID(t *testing.T) {
	u.ID = bson.NewObjectId()
	err := u.Save()
	checkError(t, err, "save user was failed")
	assert.NotNil(t, u.ID)
}

func Test2_One_GivenID_WhenReadOne_ThenReturnUser(t *testing.T) {
	u, err := One(u.ID)
	checkError(t, err, "read user was failed")
	assert.Equal(t, "James", u.Name)
	assert.Equal(t, "Developer", u.Role)
}

func Test3_Update_GivenNewChanges_WhenUpdate_ThenApplyTheChanges(t *testing.T) {
	u.Role = "Team Lead"
	err := u.Save()
	checkError(t, err, "update user was failed")
	u2, err := One(u.ID)
	checkError(t, err, "read user was failed")
	assert.Equal(t, u2.Name, "James")
	assert.Equal(t, u2.Role, "Team Lead")
}

func Test4_Delete_GivenId_WhenDelete_ThenRemoveFromDatabase(t *testing.T) {
	err := Delete(u.ID)
	checkError(t, err, "delete user was failed")
	_, err = One(u.ID)
	assert.Equal(t, err, storm.ErrNotFound)
}

func checkError(t *testing.T, err error, text string) {
	if err != nil {
		t.Fatalf("%s: %s", text, err)
	}
}
