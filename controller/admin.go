package controller

import (
	"net/http"

	"github.com/ipfans/echo-session"
	"github.com/labstack/echo"
)

// func Admin~~(API)

// AdminAuthAPI 관리자 로그인 여부 확인
func AdminAuthAPI(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		session := session.Default(c)
		if session.Get("ID") == nil {
			return c.Redirect(http.StatusMovedPermanently, "/admin/login")
		}
		return next(c)
	}
}

// AdminLogin 관리자 로그인 페이지
func AdminLogin(c echo.Context) error {
	// 미완
	return c.Render(http.StatusOK, "", nil)
}
