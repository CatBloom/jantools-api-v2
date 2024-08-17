package main

import (
	"context"
	"log"
	"main/controllers"
	"main/models"
	"main/validator"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	echoLambda *echoadapter.EchoLambda
	e          *echo.Echo
)

func init() {
	log.Println("init")

	// model
	userModel := models.NewUserModel()
	// controller
	userController := controllers.NewUserController(userModel)

	e = echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		log.Printf("ReqBody: %s", string(reqBody))
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

	e.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "GoRestTemplateAPI")
	})

	api := e.Group("/api")
	api.GET("/users", userController.List)
	api.GET("/user/:id", userController.Get)
	api.POST("/user", userController.Post)

	echoLambda = echoadapter.New(e)
}

func main() {
	if os.Getenv("ENV") == "local" {
		e.Start(":8080")
	} else {
		lambda.Start(handler)
	}
}

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return echoLambda.ProxyWithContext(ctx, req)
}
