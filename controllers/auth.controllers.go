package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	dto "echo-demo/DTOs"
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
	u := new(dto.User)
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

	newUser := dto.User{
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
	u := new(dto.User)
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
		createNewRefreshToken = true
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

	authInfo := form.Auth{
		AccessToken:  aT,
		RefreshToken: user.RefreshToken,
		User:         user,
	}

	return c.JSON(http.StatusOK, authInfo)
}

func Refresh(c echo.Context) error {
	accessToken := method.ExtractToken(c.Request())
	if accessToken == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Không tìm thấy access token!",
		})
	}

	idAT, err := method.DecodeToken(accessToken, env.AuthAccessTokenSecret)
	if err != nil || idAT < 1 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Access token không hợp lệ!",
		})
	}

	user, err := model.GetUserById(idAT)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "User không tồn tại!",
		})
	}

	a := new(form.Auth)
	if err := c.Bind(a); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	idRT, err := method.VerifyToken(a.RefreshToken, env.AuthRefreshTokenSecret)
	if a.RefreshToken != user.RefreshToken || err != nil || idRT < 1 {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "Refresh token không hợp lệ hoặc đã hết hạn, vui lòng đăng nhập lại!",
		})
	}

	if idAT != idRT {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "Thông tin chứa trong ascess token và refresh token không khớp!",
		})
	}

	// Create access token
	aT, err := method.CreateToken(idAT, env.AuthAccessTokenSecret, env.AuthAccessTokenExp)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err)
	}

	authInfo := form.Auth{
		AccessToken:  aT,
		RefreshToken: user.RefreshToken,
		User:         user,
	}

	return c.JSON(http.StatusOK, authInfo)
}
