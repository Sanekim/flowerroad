package controller

import (
	"FlowerRoad/models"
	"net/http"

	session "github.com/ipfans/echo-session"
	"github.com/labstack/echo"
)

// AuthAPI 로그인 여부 확인
func AuthAPI(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		session := session.Default(c)

		// 로그인 되어 있지 않을 때
		if session.Get("cnsanetID") != nil {
			return c.Redirect(http.StatusMovedPermanently, "/publicIndex")
		}
		return next(c)
	}
}

// Index 학생 메인 페이지
func Index(c echo.Context) error {
	return c.Render(http.StatusOK, "index", nil)
}

// Login 학생 로그인 페이지
func Login(c echo.Context) error {
	return c.Render(http.StatusOK, "login", nil)
}

// LoginPost 로그인
func LoginPost(c echo.Context) error {
	isSuccessed, nameOrErr := models.Login(c.FormValue("studentNumber"), c.FormValue("password"))

	if isSuccessed == false {
		c.Redirect(http.StatusMovedPermanently, "/login?status="+nameOrErr)
	}

	session := session.Default(c)
	session.Set("studentNumber", c.FormValue("studentNumber"))
	session.Set("name", nameOrErr)
	session.Save()

	return c.Redirect(http.StatusMovedPermanently, "/")
} 

// Logout 로그아웃
func Logout(c echo.Context) error {
	session := session.Default(c)
	session.Clear()
	session.Save()

	return c.Redirect(http.StatusMovedPermanently, "/publicIndex")
}
