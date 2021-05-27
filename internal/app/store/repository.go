package store

import "github.com/bulbetski/learn_http-rest-api/internal/app/model"

type UserRepository interface {
	Create(user *model.User) error
	Find(id int) (*model.User, error)
	FindByEmail(string) (*model.User, error)
}
