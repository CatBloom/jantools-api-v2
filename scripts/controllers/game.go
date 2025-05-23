package controllers

import (
	"log"
	"main/models"
	"main/types"
	"net/http"

	"github.com/labstack/echo/v4"
)

type GameController interface {
	List(c echo.Context) error
	Get(echo.Context) error
	Post(echo.Context) error
	Put(c echo.Context) error
	Delete(c echo.Context) error
}

type gameController struct {
	m models.GameModel
}

func NewGameController(m models.GameModel) GameController {
	return &gameController{m}
}

func (gc *gameController) List(c echo.Context) error {
	req := types.ReqGetGameList{}

	if err := c.Bind(&req); err != nil {
		log.Printf("error:%s", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(req); err != nil {
		log.Printf("error:%s", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := gc.m.GetGameList(req)
	if err != nil {
		log.Printf("error:%s", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func (gc *gameController) Get(c echo.Context) error {
	req := types.ReqGetGame{}

	if err := c.Bind(&req); err != nil {
		log.Printf("error:%s", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(req); err != nil {
		log.Printf("error:%s", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := gc.m.GetGame(req)
	if err != nil {
		log.Printf("error:%s", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func (gc *gameController) Post(c echo.Context) error {
	req := types.ReqPostGame{}

	if err := c.Bind(&req); err != nil {
		log.Printf("error:%s", err.Error())
		return err
	}

	if err := c.Validate(req); err != nil {
		log.Printf("error:%s", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// 認証情報を確認
	if leagueID, ok := c.Get("leagueID").(string); ok {
		req.LeagueID = leagueID
	} else {
		return c.JSON(http.StatusBadRequest, "Invalid token")
	}

	id, err := gc.m.CreateGame(req)
	if err != nil {
		log.Printf("error:%s", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	gameReq := types.ReqGetGame{
		ID:       id,
		LeagueID: req.LeagueID,
	}

	res, err := gc.m.GetGame(gameReq)
	if err != nil {
		log.Printf("error:%s", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func (gc *gameController) Put(c echo.Context) error {
	req := types.ReqPutGame{}

	if err := c.Bind(&req); err != nil {
		log.Printf("error:%s", err.Error())
		return err
	}

	if err := c.Validate(req); err != nil {
		log.Printf("error:%s", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// 認証情報を確認
	if leagueID, ok := c.Get("leagueID").(string); ok {
		req.LeagueID = leagueID
	} else {
		return c.JSON(http.StatusBadRequest, "Invalid token")
	}

	res, err := gc.m.UpdateGame(req)
	if err != nil {
		log.Printf("error:%s", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func (gc *gameController) Delete(c echo.Context) error {
	req := types.ReqDeleteGame{}

	if err := c.Bind(&req); err != nil {
		log.Printf("error:%s", err.Error())
		return err
	}

	if err := c.Validate(req); err != nil {
		log.Printf("error:%s", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// 認証情報を確認
	if leagueID, ok := c.Get("leagueID").(string); ok {
		req.LeagueID = leagueID
	} else {
		return c.JSON(http.StatusBadRequest, "Invalid token")
	}

	id, err := gc.m.DeleteGame(req)
	if err != nil {
		log.Printf("error:%s", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res := map[string]string{
		"id": id,
	}

	return c.JSON(http.StatusOK, res)
}
