package main

import (
	"log"
	"main/controllers"
	"main/dynamo"
	"main/middleware"
	"main/models"
	"main/validator"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

func newEcho() *echo.Echo {
	e := echo.New()
	e.Pre(echomiddleware.RemoveTrailingSlash())
	e.Use(echomiddleware.Recover())
	e.Use(echomiddleware.Logger())
	e.Use(echomiddleware.CORS())
	e.Use(echomiddleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		if len(reqBody) != 0 {
			log.Printf("ReqBody: %s", string(reqBody))
		}
	}))

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		code := http.StatusInternalServerError
		msg := err.Error()

		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
			msg = he.Message.(string)
		}
		if !c.Response().Committed {
			c.JSON(code, map[string]string{"error": msg})
		}
	}
	e.Validator = validator.NewValidator()

	return e
}

func setupRoutes(e *echo.Echo, db dynamo.DynamoDB) {
	// model
	leagueModel := models.NewLeagueModel(db)
	gameModel := models.NewGameModel(db)
	// controller
	authController := controllers.NewAuthController(leagueModel)
	leagueController := controllers.NewLeagueController(leagueModel, gameModel)
	gameController := controllers.NewGameController(gameModel)

	e.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Jantools-api-v2")
	})

	api := e.Group("/api/v2")

	auth := api.Group("/auth")
	{
		auth.POST("/token", authController.CreateAuthToken)
	}

	league := api.Group("/league", middleware.CustomAuthMiddleware())
	{
		league.GET("", leagueController.Get)
		league.POST("", leagueController.Post)
		league.PUT("", leagueController.Put)
		league.DELETE("", leagueController.Delete)
	}

	game := api.Group("/game", middleware.CustomAuthMiddleware())
	{
		game.GET("/list", gameController.List)
		game.GET("", gameController.Get)
		game.POST("", gameController.Post)
		game.PUT("", gameController.Put)
		game.DELETE("", gameController.Delete)
	}
}

func main() {
	e := newEcho()
	db := dynamo.NewDynamoDB()
	setupRoutes(e, db)

	if os.Getenv("ENV") == "local" {
		e.Start(":8080")
		return
	}
	lambda.Start(echoadapter.New(e).ProxyWithContext)
}
