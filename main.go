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

	result := functions.Transform(text, chars)
	fmt.Fprintf(w, result)
}

func main() {
	fileServer := http.FileServer(http.Dir("./Static"))
	http.Handle("/", fileServer)
	http.Handle("/ascii-art", formHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}