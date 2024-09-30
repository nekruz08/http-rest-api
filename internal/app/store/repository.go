package store

import "github.com/nekruz08/http-rest-api/internal/app/model"

type UserRepository interface {
	Create(*model.User)error
	Find(int) (*model.User,error)
	FindByEmail(string)(*model.User,error)
}