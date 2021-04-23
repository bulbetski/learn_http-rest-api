package teststore_test

import (
	"github.com/bulbetski/learn_http-rest-api/internal/app/model"
	"github.com/bulbetski/learn_http-rest-api/internal/app/store"
	"github.com/bulbetski/learn_http-rest-api/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s := teststore.New()
	u := model.TestUser(t)
	err := s.User().Create(u)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s := teststore.New()
	email := "user@example.org"

	t.Run("email does bot exist", func(t *testing.T) {
		_, err := s.User().FindByEmail(email)
		assert.EqualError(t, err, store.ErrRecordNotFound.Error())
	})

	t.Run("email exists", func(t *testing.T) {
		u := model.TestUser(t)
		u.Email = email
		s.User().Create(u)

		u, err := s.User().FindByEmail(email)
		assert.NoError(t, err)
		assert.NotNil(t, u)
		assert.Equal(t, email, u.Email)
	})
}
