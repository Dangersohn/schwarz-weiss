package main

import "regexp"

//Card hat einen Namen und  eine Farbe
type Card struct {
	Name  string
	Color string
}

func getcolor(s string) string {
	var re = regexp.MustCompile(`([a-z])`)
	treffer := re.FindAllString(s, -1)
	if treffer[0] == "b" {
		return "black"
	} else {
		return "white"
	}
}
