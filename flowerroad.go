package main

import (
	"html/template"
	"io"

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
			//관리자용

			// 공용
			"view/login.html",
			// 교사용
			"view/teacher/login.html", "view/teacher/mainpage.html",
			// 학생용
			"view/view/login.html", "view/view/mainpage.html",
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

	// ================ 학생 페이지 ===================
	// ================ 학생 API ======================

	// ================ 교사 페이지 ===================

	// ================ 교사 API =====================

	// ================ 관리자 페이지 =================

	// ================ 관리자 API ===================

	e.Start(":80")
}
