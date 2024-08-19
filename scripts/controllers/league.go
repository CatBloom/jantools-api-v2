package controllers

import (
	"log"
	"main/models"
	"main/types"
	"net/http"

	"github.com/labstack/echo/v4"
)

type LeagueController interface {
	Get(echo.Context) error
	Post(echo.Context) error
}

type leagueController struct {
	m models.LeagueModel
}

func NewLeagueController(m models.LeagueModel) LeagueController {
	return &leagueController{m}
}

func (lc *leagueController) Get(c echo.Context) error {
	id := c.Param("id")

	res, err := lc.m.GetLeagueByID(id)
	if err != nil {
		log.Printf("error:%s", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func (lc *leagueController) Post(c echo.Context) error {
	req := types.League{}

	if err := c.Bind(&req); err != nil {
		log.Printf("error:%s", err.Error())
		return err
	}

	if err := c.Validate(req); err != nil {
		log.Printf("error:%s", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	id, err := lc.m.CreateLeague(req)
	if err != nil {
		log.Printf("error:%s", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res := map[string]string{
		"id": id,
	}

	return c.JSON(http.StatusOK, res)
}
