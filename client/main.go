package main

import (
	"text/template"
    "io"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
    templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Template {
    return &Template{
        templates: template.Must(template.ParseGlob("views/*.html")),
    }
}

func main() {
    e := echo.New()
    e.Use(middleware.Logger())

    e.Renderer = newTemplate()
    
    e.GET("/", func(c echo.Context) error {
        return c.Render(200, "index", "World")
    })

    e.Logger.Fatal(e.Start(":1234"))
}
