package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"echo-demo/config"
	form "echo-demo/forms"
	method "echo-demo/methods"
	model "echo-demo/models"
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

func Register(c echo.Context) error {
	u := new(form.User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	_, err := model.GetUserByUsername(u.Username)
	if err == nil {
		return c.JSON(http.StatusConflict, map[string]string{
			"message": "Tên đăng nhập đã tồn tại!",
		})
	}

	hashPassword, err := method.HashPassword(u.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	newUser := form.User {
		Username: u.Username,
		Password: hashPassword,
	}
	result, err := model.CreateUser(newUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, result)
}

func Login(c echo.Context) error {
	u := new(form.User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	user, err := model.GetUserByUsername(u.Username)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "Tên đăng nhập không tồn tại!",
		})
	}

	if checkPassword := method.CheckPassword(user.Password, u.Password); checkPassword == false {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "Tên đăng nhập hoặc mật khẩu không chính xác!",
		})
	}

	// Create access token
	aT, err := method.CreateToken(user.ID, env.AuthAccessTokenSecret, env.AuthAccessTokenExp)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err)
	}

	// Refresh token
	createNewRefreshToken := false

	if user.RefreshToken == "" {
		createNewRefreshToken = true;
	} else {
		_, err := method.VerifyToken(user.RefreshToken, env.AuthRefreshTokenSecret)
		if err != nil { // Co err tuc la refresh token da het han hoac xay ra loi => tao refresh token moi
			createNewRefreshToken = true
		}
	}

	if createNewRefreshToken == true {
		var err error
		var rT string
		rT, err = method.CreateToken(user.ID, env.AuthRefreshTokenSecret, env.AuthRefreshTokenExp)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err)
		}

		user.RefreshToken = rT

		_, err = model.UpdateUser(user)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err)
		}
	}

	return c.JSON(http.StatusOK, map[string]string{
		"accessToken": aT,
		"refreshToken":user.RefreshToken,
	})
}