// Sequence: 7
// Title: Forms
// Source: https://gowebexamples.com/forms/

package main

import (
    "html/template"
    "net/http"
)

type ContactDetails struct {	// we declare the form fields via struct
    Email   string
    Subject string
    Message string
}

func main() {
    tmpl := template.Must(template.ParseFiles("forms.html"))	// we are trying to capture the parses that was captured in the forms of forms.html

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
            tmpl.Execute(w, nil)
            return
        }

        details := ContactDetails{
            Email:   r.FormValue("email"),
            Subject: r.FormValue("subject"),
            Message: r.FormValue("message"),
        }

        // do something with details
        _ = details

        tmpl.Execute(w, struct{ Success bool }{true})
    })

    http.ListenAndServe(":80", nil)
}