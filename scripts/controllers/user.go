package controllers

import (
	"log"
	"main/models"
	"main/types"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController interface {
	List(echo.Context) error
	Get(echo.Context) error
	Post(echo.Context) error
}

type userController struct {
	m models.UserModel
}

func NewUserController(m models.UserModel) UserController {
	return &userController{m}
}

func (uc *userController) List(c echo.Context) error {
	req := types.ReqUser{}

	if err := c.Bind(&req); err != nil {
		log.Printf("error:%s", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(req); err != nil {
		log.Printf("error:%s", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	results, err := uc.m.GetUsers(req)
	if err != nil {
		log.Printf("error:%s", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res := types.Response{
		Count:   len(results),
		Results: results,
	}
	return c.JSON(http.StatusOK, res)
}

func (uc *userController) Get(c echo.Context) error {
	id := c.Param("id")

	intID, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("error:%s", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result, err := uc.m.GetUserByID(intID)
	if err != nil {
		log.Printf("error:%s", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	count := 0
	if result.ID > 0 {
		count = 1
	}

	res := types.Response{
		Count:   count,
		Results: result,
	}
	return c.JSON(http.StatusOK, res)
}

func (uc *userController) Post(c echo.Context) error {
	req := types.ReqCreateUser{}

	if err := c.Bind(&req); err != nil {
		log.Printf("error:%s", err.Error())
		return err
	}

	if err := c.Validate(req); err != nil {
		log.Printf("error:%s", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	id, err := uc.m.CreateUser(req)
	if err != nil {
		log.Printf("error:%s", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result, err := uc.m.GetUserByID(id)
	if err != nil {
		log.Printf("error:%s", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	count := 0
	if result.ID > 0 {
		count = 1
	}

	res := types.Response{
		Count:   count,
		Results: result,
	}
	return c.JSON(http.StatusOK, res)
}
