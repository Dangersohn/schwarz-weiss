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

var a = []string{"01b.png", "02b.png", "03b.png", "04b.png", "05b.png", "06b.png", "07b.png", "08b.png", "09b.png", "10b.png", "11b.png", "12b.png", "13b.png", "14b.png", "15b.png", "16b.png", "01w.png", "02w.png", "03w.png", "04w.png", "05w.png", "06w.png", "07w.png", "08w.png", "09w.png", "10w.png", "11w.png", "12w.png", "13w.png", "14w.png", "15w.png", "16w.png"}

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
