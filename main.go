package main

import (
	"fmt"
	"log"
	"net/http"

	"ascii-art-web/Functions"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
        fmt.Fprintf(w, "ParseForm() err: %v", err)
        return
    }

	text := r.FormValue("input")
	banner := r.FormValue("banner")

	path := "Banners/"+banner+".txt"
	chars := functions.FileInit(path)
	result := ""

	for i, j:= 0, 0; i < len(text); i++ {
		if text[i] == '\r'{
			result += functions.Transform(text[j:i], chars)
			j = i + 2
		} else if i == len(text)-1 {
			result += functions.Transform(text[j:], chars)
		}
	}

	fmt.Fprintf(w, result)
}

func main() {
	fileServer := http.FileServer(http.Dir("./Static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/ascii-art", formHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}