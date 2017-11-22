package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

var a = []string{"01b.jpg", "02b.jpg", "03b.jpg", "04b.jpg", "05b.jpg", "06b.jpg", "07b.jpg", "08b.jpg", "09b.jpg", "10b.jpg", "11b.jpg", "12b.jpg", "13b.jpg", "14b.jpg", "15b.jpg", "16b.jpg", "01w.jpg", "02w.jpg", "03w.jpg", "04w.jpg", "05w.jpg", "06w.jpg", "07w.jpg", "08w.jpg", "09w.jpg", "10w.jpg", "11w.jpg", "12w.jpg", "13w.jpg", "14w.jpg", "15w.jpg", "16w.jpg"}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {

	t := &Template{
		templates: template.Must(template.ParseGlob("template/*.html")),
	}
	e := echo.New()
	e.Renderer = t
	e.GET("/", showrandom)
	e.Static("/images/*", "images")
	log.Fatal(e.Start(":" + os.Getenv("PORT")))
}

func show(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}

func showrandom(c echo.Context) error {
	if len(a) == 0 {
		return c.Render(http.StatusOK, "end.html", nil)
	}
	rnd := rand.Intn(len(a))
	fmt.Println(rnd)
	speicher := a[rnd]
	a = remove(a, a[rnd])
	return c.Render(http.StatusOK, "index.html", speicher)
}

// removes one item from the slice
func remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}
