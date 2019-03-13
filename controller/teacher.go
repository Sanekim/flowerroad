package controller

import (
	"net/http"

	session "github.com/ipfans/echo-session"
	"github.com/labstack/echo"
)

// TeacherAuthAPI 관리자 로그인 여부 확인
func TeacherAuthAPI(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		session := session.Default(c)
		if session.Get("cnsanetID") == nil {
			return c.Redirect(http.StatusMovedPermanently, "/teacher/login")
		}
		return next(c)
	}
}
