package store_test

import (
	"github.com/bulbetski/learn_http-rest-api/internal/app/model"
	"github.com/bulbetski/learn_http-rest-api/internal/app/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseUrl)
	defer teardown("users")

	u, err := s.User().Create(&model.User{
		Email: "user@example.org",
	})
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseUrl)
	defer teardown("users")

	email := "user@example.org"

	t.Run("email does bot exist", func(t *testing.T) {
		_, err := s.User().FindByEmail(email)
		assert.Error(t, err)
	})

	t.Run("email exists", func(t *testing.T) {
		s.User().Create(&model.User{
			Email: email,
		})

		u, err := s.User().FindByEmail(email)
		assert.NoError(t, err)
		assert.NotNil(t, u)
		assert.Equal(t, email, u.Email)
	})
}
