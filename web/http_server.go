// Sequence: 2
// Title: HTTP Server
// Source: https://gowebexamples.com/http-server/

package main

import (
    "fmt"		// this allows us to use fmt.Fprintf, fmt.Println, etc.
    "net/http"	// this allows us to use the http handlers like http.HandleFunc and http.Handle
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {		// use http.HandleFunc() if you want a function to respond to the HTTP request.
        fmt.Fprintf(w, "Welcome to my website!\n")							// Prints: Welcome to my website!
        fmt.Fprintf(w, "Directory: '%s", r.URL.Path)						// Prints: Directory: '/example' doesn't exist.
		fmt.Fprintf(w, "' doesn't exist.\n\n")								// w variable is the response writer that prints the front end and the r was used to get the URL path only
    })

    fs := http.FileServer(http.Dir("static/"))					// here we declare the catcher of the static directory and for everything else, it's handled by the http.HandleFunc()
    http.Handle("/static/", http.StripPrefix("/static/", fs))	// use http.Handle() to create an object type that responds to HTTP requests, it redirects.
	
	// Below, we will create our very own result from what we've learned.
	http.HandleFunc("/dynamic/", func(w http.ResponseWriter, r *http.Request) { 	// if a user visits localhost/dynamic, it will change the entire page to the one below. 
		fmt.Fprintf(w, "So, you're looking for dynamic huh.")						// Prints: "So, you're looking for dynamic huh."
	})
	/* 
	So now, this program does 3 things.
	
	First is if the directory URL isn't domain.com/static or domain.com/dynamic, it would print below for the samples: domain.com/strategic/sample and domain.com/mylife
	Prints: Welcome to my website!
			Directory: '/strategic/sample' doesn't exist.
				
			Welcome to my website!
			Directory: '/mylife' doesn't exist.
				
	Second is if the directory URL is domain.com/static, it would browse the static folder and show its content(s).

	Third is if the directory URL is domain.com/dynamic, it prints below.
	Prints:	So, you're looking for dynamic huh.
	*/

    http.ListenAndServe(":81", nil)		// this has 2 syntax, the first one is which port to use and the second one is the DefaultServeMux (which is Golang's default server)
}