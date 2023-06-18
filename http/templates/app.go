package main

import (
	"fmt"
	"html/template"
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
	data  PageData
	templ *template.Template
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func getHome(response http.ResponseWriter, request *http.Request) {
	templ.Execute(response, data)
}

func getView(fileName string) (string, error) {
	file, err := os.ReadFile(fileName)
	return string(file), err
}

func getTemplate(templName, html string) *template.Template {
	return template.Must(template.New(templName).Parse(html))
}

func main() {
	// Set page data
	person := Person{Name: "Pedro", Age: 26, Alive: true, Hobbies: []string{"Movies", "Restaraunts"}}
	data = PageData{Title: "HTML Templates", Data: person}

	// Get template file content
	html, err := getView("index.html")
	handleError(err)

	templ = getTemplate("home", html)

	// Write rendered template on console
	// templ.Execute(os.Stdout, data)

	fmt.Println("\n⚡ Server on http://localhost:8080")
	http.HandleFunc("/", getHome)
	http.ListenAndServe(":8080", nil)
}
