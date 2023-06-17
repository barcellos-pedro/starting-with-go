package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Person struct {
	Name    string
	Age     int
	Alive   bool
	Hobbies []string
}

type PageData struct {
	Title string
	Data  Person
}

var (
	PORT = 8080
	html string
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	if file, err := os.ReadFile("index.html"); err != nil {
		handleError(err)
	} else {
		html = string(file)
	}

	person := Person{Name: "Pedro", Age: 26, Alive: true, Hobbies: []string{"Movies", "Restaraunts"}}
	data := PageData{Title: "HTML Templates", Data: person}
	templ := template.Must(template.New("home").Parse(html))

	// Show template rendered result on console
	templ.Execute(os.Stdout, data)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templ.Execute(w, data)
	})

	fmt.Println("\n⚡ Server on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
