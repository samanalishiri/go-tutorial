package user

import (
	"errors"
	"gopkg.in/mgo.v2/bson"
)

type (
	User struct {
		ID   bson.ObjectId `json:"id" storm:"id"`
		Name string        `json:"name"`
		Role string        `json:"role"`
	}
)

const (
	dbPath = "users.db"
)

var (
	ErrRecordInvalid = errors.New("record is invalid")
)

func (u *User) validate() error {
	if u.Name == "" {
		return ErrRecordInvalid
	}
	return nil
}
