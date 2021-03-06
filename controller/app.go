package controller

import (
	"net/http"

	session "github.com/ipfans/echo-session"
	"github.com/labstack/echo"
)

// AuthAPI 로그인 여부 확인
func AuthAPI(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		session := session.Default(c)

		// 로그인 되어 있지 않을 때
		if session.Get("studentNumber") == nil {
			return c.Redirect(http.StatusMovedPermanently, "/login")
		}
		return next(c)
	}
}
