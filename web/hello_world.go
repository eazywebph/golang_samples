// Sequence: 1
// Title: Hello World
// Source: https://gowebexamples.com/hello-world/

package main

import (
    "fmt" 		// this allows us to use fmt.Fprintf, fmt.Println, etc.
    "net/http"	// this allows us to use the http handlers like http.HandleFunc and http.Handle
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {		// use http.HandleFunc() if you want a function to respond to the HTTP request.
        fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)			// let's say this page would open localhost:81 and if you go to localhost:81/test, it would print "Hello, you've requested: /test"
    })

    http.ListenAndServe(":81", nil) // this has 2 syntax, the first one is which port to use and the second one is the DefaultServeMux (which is Golang's default server)
}