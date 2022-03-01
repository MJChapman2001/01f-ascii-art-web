package main

import (
	functions "ascii-art-web/Functions"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type LargeText struct {
	Input  string
	Banner string
	Ltext  string
}

var Tmpl = template.Must(template.ParseGlob("Templates/*.html"))

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 Status not found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		var data LargeText
		data.Input = ""
		data.Banner = "standard"
		data.Ltext = ""
		if err := Tmpl.Execute(w, data); err != nil {
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return
		}
	case "POST":
		result := ""
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		text := r.FormValue("input")
		banner := r.FormValue("banners")

		if text == "" || banner == "" {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		path := "Banners/" + banner + ".txt"
		chars := functions.FileInit(path)

		for i, j := 0, 0; i < len(text); i++ {
			if text[i] == '\r' {
				result += functions.Transform(text[j:i], chars)
				j = i + 2
			} else if i == len(text)-1 {
				result += functions.Transform(text[j:], chars)
			}
		}

		var data LargeText
		data.Input = text
		data.Banner = banner
		data.Ltext = result
		if err := Tmpl.Execute(w, data); err != nil {
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
}

func main() {
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", formHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
