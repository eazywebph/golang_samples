// Sequence: 5
// Title: Templates
// Source: https://gowebexamples.com/templates/

package main

import (
    "html/template" // this allows us to use fmt.Fprintf, fmt.Println, etc.
    "net/http"		// this allows us to use the http handlers like http.HandleFunc and http.Handle
)

type Todo struct {	// a struct is a collection of fields or variables which we can use to display dynamic results later on.
    Title string
    Done  bool
}

type TodoPageData struct {
    PageTitle string
    Todos     []Todo
}

func main() {
    tmpl := template.Must(template.ParseFiles("layout.html"))			// Must is a helper that wraps a call to a function returning (*Template, error) and panics if the error is non-nil. It is intended for use in variable initializations such as this line of code. 

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {	// here is where we are using an http.HandleFunc since we're not exposing the entire directory, we just want the single page to interact with the below values.
	
        data := TodoPageData{
            PageTitle: "My TODO list",
            Todos: []Todo{
                {Title: "Python", Done: false},
                {Title: "PHP", Done: true},
                {Title: "Golang", Done: false},
            },
        }
        tmpl.Execute(w, data)
    })
    http.ListenAndServe(":81", nil)
}

/*
When interacting with a database, we can fetch the data from the database and define them through struct which we can then use as variables which then we can pass on to the HTML file.

Control Structures
The templating language contains a rich set of control structures to render your HTML. Here you will get an overview of the most commonly used ones. To get a detailed list of all possible structures visit: text/template

Control Structure	Definition
{{/* a comment *//*}}	Defines a comment
{{.}}	Renders the root element
{{.Title}}	Renders the “Title”-field in a nested element
{{if .Done}} {{else}} {{end}}	Defines an if-Statement
{{range .Todos}} {{.}} {{end}}	Loops over all “Todos” and renders each using {{.}}
{{block "content" .}} {{end}}	Defines a block with the name “content”
*/