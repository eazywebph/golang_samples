// Sequence: 8
// Title: Middleware (Basic)
// Source: https://gowebexamples.com/basic-middleware/

package main

import (
    "fmt"
    "log"
    "net/http"
)

const home = '
<h1>Hello, you've reached the foo page.</h1>
<h3>Click <a href="bar">here</a> to proceed.</h3>
'

func logging(f http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        log.Println(r.URL.Path)
        f(w, r)
    }
}

func foo(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "foo")
}

func bar(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "bar")
}

func main() {
    http.HandleFunc("/foo", logging(foo))
    http.HandleFunc("/bar", logging(bar))

    http.ListenAndServe(":80", nil)
}