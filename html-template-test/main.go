package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.go.tmpl")),
	}
	e.Renderer = t
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "home.go.tmpl", struct {
			Name string
		}{
			Name: "World",
		})
	})
	e.Logger.Fatal(e.Start("localhost:8080"))
}
