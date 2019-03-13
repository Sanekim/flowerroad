package controller

import (
	"net/http"

	session "github.com/ipfans/echo-session"
	"github.com/labstack/echo"
)

// TeacherAuthAPI 교사 로그인 여부 확인
func TeacherAuthAPI(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		session := session.Default(c)
		if session.Get("cnsanetID") == nil {
			return c.Redirect(http.StatusMovedPermanently, "/teacher/login")
		}
		return next(c)
	}
}

// TeacherLogin 교사 로그인 페이지
func TeacherLogin(c echo.Context) error {
	return c.Render(http.StatusOK, "teacherLogin", nil)
}
