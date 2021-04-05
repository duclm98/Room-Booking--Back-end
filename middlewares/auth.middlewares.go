package middlewares

import (
	"fmt"
	"net/http"
	
	"github.com/labstack/echo"

	config "echo-demo/config"
	method "echo-demo/methods"
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

func Verify(r *http.Request) bool {
	token := method.ExtractToken(r)
	if token == "" {
		return false
	}
	id, err := method.VerifyToken(token, env.AuthAccessTokenSecret)
	if err != nil || id < 1 {
		return false
	}
	return true
}

func IsAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		isAuth := Verify(c.Request())
		if isAuth == false {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "Chưa đăng nhập vào hệ thống!",
			})
		}
		return next(c)
	}
}