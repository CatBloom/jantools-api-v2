package models

import (
	"main/types"
)

type UserModel interface {
	GetUsers(types.ReqUser) ([]types.User, error)
	GetUserByID(int) (types.User, error)
	CreateUser(types.ReqCreateUser) (int, error)
}

type userModel struct {
}

func NewUserModel() UserModel {
	return &userModel{}
}

func (um *userModel) GetUsers(req types.ReqUser) ([]types.User, error) {
	results := []types.User{}
	return results, nil
}

func (um *userModel) GetUserByID(id int) (types.User, error) {
	result := types.User{}
	return result, nil
}

func (um *userModel) CreateUser(req types.ReqCreateUser) (int, error) {
	id := 0

	return int(id), nil
}
