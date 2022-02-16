package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"ascii-art-web/Functions"
)

type LargeText struct {
	Ltext string
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./Templates/index.html"))

	switch r.Method {
	case "GET":
		data := LargeText{Ltext: "Enter some text",}
		tmpl.Execute(w, data)
	case "POST":
		result := ""
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
	
		text := r.FormValue("input")
		banner := r.FormValue("banner")
	
		path := "Banners/"+banner+".txt"
		chars := functions.FileInit(path)
	
		for i, j:= 0, 0; i < len(text); i++ {
			if text[i] == '\r'{
				result += functions.Transform(text[j:i], chars)
				j = i + 2
			} else if i == len(text)-1 {
				result += functions.Transform(text[j:], chars)
			}
		}
		
		data := LargeText{Ltext: result,}
		tmpl.Execute(w, data)
	}
}

func main() {
	http.HandleFunc("/", formHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}