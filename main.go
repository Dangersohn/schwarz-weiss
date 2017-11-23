package main

import (
	"html/template"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

// Template wird für das Rendern benötigt
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
	speicher := a[rnd]

	col := getcolor(speicher)

	karte := Card{
		Name:  speicher,
		Color: col,
	}

	a = remove(a, a[rnd])
	return c.Render(http.StatusOK, "index.html", karte)
}
