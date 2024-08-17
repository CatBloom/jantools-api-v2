package models

import (
	"main/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUsersModel(t *testing.T) {
	t.Run("Test with success response", func(t *testing.T) {
		req := types.ReqUser{
			Limit: 2,
			Order: "created_at",
		}

		model := NewUserModel()

		results, err := model.GetUsers(req)
		if assert.NoError(t, err) {
			for _, result := range results {
				assert.NotEmpty(t, result.ID)
				assert.NotEmpty(t, result.Name)
				assert.NotEmpty(t, result.Email)
				assert.NotEmpty(t, result.CreatedAt)
			}
		}
	})
}

func TestGetUserByIDModel(t *testing.T) {
	t.Run("Test with success response", func(t *testing.T) {
		model := NewUserModel()

		result, err := model.GetUserByID(1)
		if assert.NoError(t, err) {
			assert.NoError(t, err)
			assert.NotEmpty(t, result.ID)
			assert.NotEmpty(t, result.Name)
			assert.NotEmpty(t, result.Email)
			assert.NotEmpty(t, result.CreatedAt)
		}
	})
}

func TestCreateUserModel(t *testing.T) {
	t.Run("Test with success response", func(t *testing.T) {
		req := types.ReqCreateUser{
			Name:  "user",
			Email: "test@exsample.com",
		}

		model := NewUserModel()

		_, err := model.CreateUser(req)
		assert.NoError(t, err)
	})
}
