/*

This is a simple golang code that starts with a static page and when directing to /login, the page would show up a login form (username and password).
The activity is being logged in the background through ParseForm() then captures the method and the username and password being used.

Source: https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/04.1.html

*/

package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
    r.ParseForm() //Parse url parameters passed, then parse the response packet for the POST body (request body)
    // attention: If you do not call ParseForm method, the following data can not be obtained form
    fmt.Println(r.Form) // print information on server side.
    fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
    fmt.Fprintf(w, "Hello astaxie!") // write data to response
    fmt.Fprintf(w, "<a href='login'>Login</a>") // write data to response
}

func login(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method) //get request method
    if r.Method == "GET" {
        t, _ := template.ParseFiles("login.gtpl")
        t.Execute(w, nil)
    } else {
        r.ParseForm()
        // logic part of log in
        fmt.Println("username:", r.Form["username"])
        fmt.Println("password:", r.Form["password"])
    }
}

func hero(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("login.php")
	t.Execute(w, nil)
    // fmt.Fprintf(w, "This is a restricted space.") // write data to response
}

func main() {
    http.HandleFunc("/", sayhelloName) // setting router rule
    http.HandleFunc("/login", login)
    http.HandleFunc("/hero", hero)
    err := http.ListenAndServe(":80", nil) // setting listening port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}