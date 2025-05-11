package controllers

import (
	"log"
	"main/models"
	"main/types"
	"main/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController interface {
	CreateAuthToken(c echo.Context) error
}

type authController struct {
	lm models.LeagueModel
}

func NewAuthController(lm models.LeagueModel) AuthController {
	return &authController{lm}
}

func (ac *authController) CreateAuthToken(c echo.Context) error {
	req := types.ReqAuthToken{}
	if err := c.Bind(&req); err != nil {
		log.Printf("error:%s", err.Error())
		return err
	}

	if err := c.Validate(req); err != nil {
		log.Printf("error:%s", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	league, err := ac.lm.GetLeague(types.ReqGetLeague{ID: req.ID})
	if err != nil {
		log.Printf("error:%s", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if !utils.VerifyPassword(req.Password, league.Password) {
		log.Printf("error:%s", "Invalid password")
		return c.JSON(http.StatusBadRequest, "Invalid password")
	}
	jwt, _ := utils.GenerateJWT(league.ID)

	res := map[string]string{
		"token": jwt,
	}

	return c.JSON(http.StatusOK, res)
}
