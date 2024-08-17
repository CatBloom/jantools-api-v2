package controllers

import (
	"encoding/json"
	"errors"
	"main/types"
	"main/validator"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type mockUserModel struct{}

func (m *mockUserModel) GetUsers(types.ReqUser) ([]types.User, error) {
	users := []types.User{
		{
			Name:      "name1",
			Email:     "test@example.com",
			CreatedAt: time.Now(),
		},
		{
			Name:      "name2",
			Email:     "test2@example.com",
			CreatedAt: time.Now(),
		},
	}
	return users, nil
}

func (m *mockUserModel) GetUserByID(_ int) (types.User, error) {
	user := types.User{
		Name:      "name1",
		Email:     "test@example.com",
		CreatedAt: time.Now(),
	}
	return user, nil
}

func (m *mockUserModel) CreateUser(types.ReqCreateUser) (int, error) {
	return 1, nil
}

type mockErrorUserModel struct{}

func (m *mockErrorUserModel) GetUsers(types.ReqUser) ([]types.User, error) {
	return []types.User{}, errors.New("Failed to get users")
}

func (m *mockErrorUserModel) GetUserByID(_ int) (types.User, error) {
	return types.User{}, errors.New("Failed to get user")
}

func (m *mockErrorUserModel) CreateUser(types.ReqCreateUser) (int, error) {
	return 0, errors.New("Failed to create user")
}

func TestGetUsersController(t *testing.T) {
	t.Run("Test with success response", func(t *testing.T) {
		e := echo.New()
		e.Validator = validator.NewValidator()

		mockModel := &mockUserModel{}
		controller := NewUserController(mockModel)

		req := httptest.NewRequest(http.MethodGet, "/users", nil)
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)

		if assert.NoError(t, controller.List(c)) {
			assert.Equal(t, http.StatusOK, res.Code)
		}
	})

	t.Run("Test with error response", func(t *testing.T) {
		e := echo.New()
		e.Validator = validator.NewValidator()

		mockErrorModel := &mockErrorUserModel{}
		controller := NewUserController(mockErrorModel)

		req := httptest.NewRequest(http.MethodGet, "/users", nil)
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)

		controller.List(c)

		assert.Equal(t, http.StatusBadRequest, res.Code)
		expectedResponse := "\"Failed to get users\"\n"
		assert.Equal(t, expectedResponse, res.Body.String())
	})
}

func TestGetUserByIDController(t *testing.T) {
	t.Run("Test with success response", func(t *testing.T) {
		e := echo.New()
		mockModel := &mockUserModel{}
		controller := NewUserController(mockModel)

		req := httptest.NewRequest(http.MethodGet, "/user/1", nil)
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, controller.Get(c)) {
			assert.Equal(t, http.StatusOK, res.Code)
		}
	})

	t.Run("Test with error response", func(t *testing.T) {
		e := echo.New()
		e.Validator = validator.NewValidator()

		mockErrorModel := &mockErrorUserModel{}
		controller := NewUserController(mockErrorModel)

		req := httptest.NewRequest(http.MethodGet, "/user/1", nil)
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)
		c.SetParamNames("id")
		c.SetParamValues("1")

		controller.Get(c)

		assert.Equal(t, http.StatusBadRequest, res.Code)
		expectedResponse := "\"Failed to get user\"\n"
		assert.Equal(t, expectedResponse, res.Body.String())
	})
}

func TestCreateUserController(t *testing.T) {
	t.Run("Test with success response", func(t *testing.T) {
		e := echo.New()
		e.Validator = validator.NewValidator()

		q := map[string]interface{}{
			"name":  "test",
			"email": "test@example.com",
		}
		enc, _ := json.Marshal(q)
		mockModel := &mockUserModel{}
		controller := NewUserController(mockModel)

		req := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(string(enc)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)

		if assert.NoError(t, controller.Post(c)) {
			assert.Equal(t, http.StatusOK, res.Code)
		}
	})

	t.Run("Test with error response", func(t *testing.T) {
		e := echo.New()
		e.Validator = validator.NewValidator()

		q := map[string]interface{}{
			"name":  "test",
			"email": "test@example.com",
		}
		enc, _ := json.Marshal(q)
		mockErrorModel := &mockErrorUserModel{}
		controller := NewUserController(mockErrorModel)

		req := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(string(enc)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)

		controller.Post(c)

		assert.Equal(t, http.StatusBadRequest, res.Code)
		expectedResponse := "\"Failed to create user\"\n"
		assert.Equal(t, expectedResponse, res.Body.String())
	})
}
