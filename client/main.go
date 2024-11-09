package main

import (
	"fmt"
	"io"
	"text/template"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
    sf "client/sendFile"
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
        return c.Render(200, "index", nil)
    })

    e.POST("/sendFile", func(c echo.Context) error {
        file, err := c.FormFile("file")
        if err != nil {
            fmt.Println(err)
            return err
        }

        sf.SendFileData(file)
        return c.Render(200, "index", nil)
    })

    e.Logger.Fatal(e.Start(":1234"))
}
