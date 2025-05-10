package controllers

import (
	"log"
	"main/models"
	"main/types"
	"main/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type LeagueController interface {
	Get(echo.Context) error
	Post(echo.Context) error
	Put(c echo.Context) error
	Delete(c echo.Context) error
}

type leagueController struct {
	m models.LeagueModel
	g models.GameModel
}

func NewLeagueController(m models.LeagueModel, g models.GameModel) LeagueController {
	return &leagueController{m, g}
}

func (lc *leagueController) Get(c echo.Context) error {
	req := types.ReqGetLeague{}

	if err := c.Bind(&req); err != nil {
		log.Printf("error:%s", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(req); err != nil {
		log.Printf("error:%s", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := lc.m.GetLeague(req)
	if err != nil {
		log.Printf("error:%s", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// パスワードは返却しないため、削除
	res.Password = ""

	return c.JSON(http.StatusOK, res)
}

func (lc *leagueController) Post(c echo.Context) error {
	req := types.ReqPostLeague{}

	if err := c.Bind(&req); err != nil {
		log.Printf("error:%s", err.Error())
		return err
	}

	if err := c.Validate(req); err != nil {
		log.Printf("error:%s", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// passwordのハッシュ化
	hashPass, err := utils.GenerateHashPassword(req.Password)
	if err != nil {
		log.Printf("error:%s", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	req.Password = hashPass

	id, err := lc.m.CreateLeague(req)
	if err != nil {
		log.Printf("error:%s", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := lc.m.GetLeague(types.ReqGetLeague{ID: id})
	if err != nil {
		log.Printf("error:%s", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func (lc *leagueController) Put(c echo.Context) error {
	req := types.ReqPutLeague{}

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
		req.ID = leagueID
	} else {
		return c.JSON(http.StatusBadRequest, "Invalid token")
	}

	res, err := lc.m.UpdateLeague(req)
	if err != nil {
		log.Printf("error:%s", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func (lc *leagueController) Delete(c echo.Context) error {
	req := types.ReqDeleteLeague{}

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
		req.ID = leagueID
	} else {
		return c.JSON(http.StatusBadRequest, "Invalid token")
	}

	gameReq := types.ReqGetGameList{
		LeagueID: req.ID,
	}
	gameList, err := lc.g.GetGameList(gameReq)
	if err != nil {
		log.Printf("error:%s", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	for _, v := range gameList {
		deleteReq := types.ReqDeleteGame{
			ID:       v.ID,
			LeagueID: req.ID,
		}
		_, err := lc.g.DeleteGame(deleteReq)
		if err != nil {
			log.Printf("error:%s", err.Error())
			return c.JSON(http.StatusBadRequest, err.Error())
		}
	}

	id, err := lc.m.DeleteLeague(req)
	if err != nil {
		log.Printf("error:%s", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res := map[string]string{
		"id": id,
	}

	return c.JSON(http.StatusOK, res)
}
