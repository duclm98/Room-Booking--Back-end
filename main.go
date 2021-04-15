package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/net/http2"

	config "echo-demo/config"
	method "echo-demo/methods"
	model "echo-demo/models"
	route "echo-demo/routes"
)

var env config.Config

func init() {
	var err error
	env, err = config.LoadConfig(".")
	if err != nil {
		fmt.Println("Cannot load auth config:", err)
		return
	}
}

func Authentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := method.ExtractToken(c.Request());
		if token == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "Không tìm thấy access token!",
			})
		}

		id, err := method.VerifyToken(token, env.AuthAccessTokenSecret)
		if err != nil || id < 1 {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "Access token không hợp lệ hoặc đã hết hạn!",
			})
		}

		user, err := model.GetUserById(id)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "User không tồn tại!",
			})
		}

		c.Set("user", user)

		return next(c)
	}
}

func main() {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, status=${status}, uri=${uri}\n",
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080"},
	}))

	api := e.Group("/api")

	route.AuthRoute(api)

	// api.Use(middleware.JWTWithConfig(middleware.JWTConfig{
	// 	SigningKey: []byte(env.AuthAccessTokenSecret),
	// }))
	api.Use(Authentication);

	route.BuildingsRoute(api)
	route.BookingRoute(api)
	
	s := &http2.Server{
		MaxConcurrentStreams: 250,
		MaxReadFrameSize: 1048576,
		IdleTimeout: 10 * time.Second,
	}

	e.Logger.Fatal(e.StartH2CServer(":3000", s))
}