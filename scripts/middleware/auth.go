package middleware

import (
	"net/http"
	"os"

	jwt "github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func CustomAuthMiddleware() echo.MiddlewareFunc {
	jwtMiddleware := echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	})

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// 認証不要なルートを先に判定
			if c.Request().Method == http.MethodGet {
				return next(c)
			}
			if c.Request().Method == http.MethodPost && c.Path() == "/api/v2/league" {
				return next(c)
			}
			return jwtMiddleware(func(c echo.Context) error {
				token, ok := c.Get("user").(*jwt.Token)
				if !ok {
					return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
				}

				claims, ok := token.Claims.(jwt.MapClaims)
				if !ok {
					return echo.NewHTTPError(http.StatusUnauthorized, "Invalid claims")
				}

				id, ok := claims["sub"].(string)
				if !ok {
					return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
				}
				c.Set("leagueID", id)
				return next(c)
			})(c)
		}
	}
}
