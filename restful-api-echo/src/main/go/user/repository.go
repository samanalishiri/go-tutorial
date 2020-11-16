package user

import (
	"github.com/asdine/storm/v3"
	"gopkg.in/mgo.v2/bson"
)

type Repository interface {
	Save(u *User) error
	FindById(id bson.ObjectId) (*User, error)
	DeleteById(id bson.ObjectId) error
	FindAll() ([]User, error)
}

type RepositoryImpl struct {
}

func NewUserRepository() *RepositoryImpl {
	return &RepositoryImpl{}
}

func (impl *RepositoryImpl) SaveUser(u *User) error {

	if err := u.validate(); err != nil {
		return err
	}

	db, err := storm.Open(dbPath)
	if err != nil {
		return err
	}
	defer db.Close()

	return db.Save(u)
}

func (impl *RepositoryImpl) FindById(id bson.ObjectId) (*User, error) {
	db, err := storm.Open(dbPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	u := new(User)
	err = db.One("ID", id, u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (impl *RepositoryImpl) DeleteById(id bson.ObjectId) error {
	db, err := storm.Open(dbPath)
	if err != nil {
		return err
	}
	defer db.Close()

	u := new(User)
	err = db.One("ID", id, u)
	if err != nil {
		return err
	}

	return db.DeleteStruct(u)
}

func (impl *RepositoryImpl) FindAll() ([]User, error) {
	db, err := storm.Open(dbPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var users []User
	err = db.All(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}
