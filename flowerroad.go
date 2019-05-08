package main

import (
	"FlowerRoad/controller"
	"html/template"
	"io"
	"net/http"

	session "github.com/ipfans/echo-session"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Template 템플릿
type Template struct {
	templates *template.Template
}

// Render renders a template document
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	t := &Template{
		templates: template.Must(template.New("").Delims("[[", "]]").ParseFiles(
			// 공용
			"view/publicIndex.html",
			// 교사용
			"view/teacher/login.html",
			// 학생용
			"view/view/login.html", "view/view/index.html", "view/view/addRoadmap.html", "view/view/diploma.html", "view/view/passwordChange.html",
		)),
	}

	e := echo.New()

	// Set middlewares
	// Logger: loging all request and responses
	// Recover: Recover main thread if it fails
	e.Use(middleware.Logger(), middleware.Recover())

	// Session 설정
	store := session.NewCookieStore([]byte("secret"))
	e.Use(session.Sessions("CASESSION", store))

	// Set template renderer
	// We uses standard golang template
	e.Renderer = t

	// Set static serve files
	e.Static("/assets", "static")

	// 공용 인덱스
	e.GET("/publicIndex", func(c echo.Context) error {
		return c.Render(http.StatusOK, "publicIndex", nil)
	})
	// 로그아웃
	e.GET("/logout", controller.Logout)

	// ================ 학생 페이지 ===================
	// 로그인 페이지
	e.GET("/login", controller.Login)
	e.POST("/login", controller.LoginPost)

	// 메인 페이지
	e.GET("/", controller.Index)

	// 비밀번호 변경 페이지
	e.GET("/passwordChange", controller.PasswordChange)

	// 디플로마 소개 페이지
	// e.GET("/diploIntro")

	// 계획표 작성 페이지
	// e.GET("/addRoadmap")

	// 계획표 확인
	// e.GET("/checkRoadmap")

	// ================ 학생 API ======================

	// ================ 교사 페이지 ===================
	// 로그인 페이지
	e.GET("/teacher/login", controller.TeacherLogin)
	e.POST("/teacher/login", controller.TeacherLoginPost)

	tcr := e.Group("/teacher")
	tcr.Use(controller.TeacherAuthAPI)

	// 메인 페이지
	tcr.GET("/", controller.TeacherIndex)
	// ================ 교사 API =====================

	// ================ 관리자 페이지 =================

	// ================ 관리자 API ===================

	e.Start(":8888")
}
