package main

import (
	"fmt"
	"io"
    "os"
    "html/template"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
    sf "client/sendFile"
)

type Template struct {
    templates *template.Template
}

type Files struct {
    Files []File
}

type File struct {
    FilePath string
    FileName string
}


func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Template {
    return &Template{
        templates: template.Must(template.ParseGlob("views/*/*.html")),
    }
}

func main() {
    e := echo.New()
    e.Use(middleware.Logger())

    e.Static("/storage", "../server/storage")

    e.Renderer = newTemplate()
    
    e.GET("/addFile", func(c echo.Context) error {
        return c.Render(200, "addFile", nil)
    })

    e.GET("/viewFiles", func(c echo.Context) error {
        filesPath := "../server/storage/"

        fsFiles, err := os.ReadDir(filesPath)
        if err != nil {
            fmt.Printf("Error reading directory: %s\n", err)
            return err
        }

        files := Files{}
        
        for _, file := range fsFiles {
            files.Files = append(files.Files, File{FilePath: "/storage/", FileName: file.Name()})
        }

        return c.Render(200, "viewFiles", files)
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
