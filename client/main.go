package main

import (
	"fmt"
	"io"
    "os"
    "strings"
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
    FileType string
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
    e.Static("/static", "public")

    e.Renderer = newTemplate()
    
    e.GET("/manageFile", func(c echo.Context) error {
        return c.Render(200, "addFile", nil)
    })

    e.GET("/", func(c echo.Context) error {
        filesPath := "../server/storage/"

        fsFiles, err := os.ReadDir(filesPath)
        if err != nil {
            fmt.Printf("Error reading directory: %s\n", err)
            return err
        }

        files := Files{}
        
        for _, file := range fsFiles {
            fileType := strings.Split(file.Name(), ".")[1]
            if fileType == "jpg" || fileType == "jpeg" || fileType == "png"  || fileType == "webp" {
                fileType = "image"
            } else if fileType == "mp4" || fileType == "avi" || fileType == "mkv" {
                fileType = "video"
            } else if fileType == "mp3" || fileType == "wav" {
                fileType = "audio"
            }

            files.Files = append(files.Files, File{FilePath: "/storage/", FileName: file.Name(), FileType: fileType})
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
