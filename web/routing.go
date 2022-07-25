// Sequence: 3
// Title: Routing (using gorilla/mux)
// Source: https://gowebexamples.com/routes-using-gorilla-mux/

/* 
Prerequisites on a new machine
In the terminal, get the gorilla/mux package.

user$: go get -u github.com/gorilla/mux
user$: go mod init directory_name
user$: go get github.com/gorilla/mux

Then the build should compile successfully.
*/


package main

import (
    "fmt"						// this allows us to use fmt.Fprintf, fmt.Println, etc.
    "net/http"					// this allows us to use the http handlers like http.HandleFunc and http.Handle
    "github.com/gorilla/mux"	// this is an added library which contains the gorilla/mux function
)

/*
We're going to create a program that would only read the URL syntax:
/books/{title}/page/{page}
/books/go-programming-blueprint/page/10

Anything else would go to a 404 page.

For this, we'll be using the gorilla/mux router which extracts segments from the request URL into single parameters.
This is useful when managing GET/POST handlers and defining domain restrictions.
*/
func main() {
    r := mux.NewRouter()	// here is where we create the (r) router which will receive all HTTP connections and pass it on to the request handlers.

    r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {	// define the directory of the handler, the 2 objects will be defined below
        vars := mux.Vars(r)																		// we create a variable (vars) that takes the HTTP request
        title := vars["title"]																	// this is where we define the title variable and put it in the vars function, this one is for the title in the directory
        page := vars["page"]																	// and this one is for the page in the directory

        fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)				// let's say the entire URL is http://localhost/books/garlic-juice/page/5, it would print below.
    })																							// Print: You've requested the book: garlic-juice on page 5
																								// If the URL changes to http://localhost/books/garlic-juice/pages/5 where pages wasn't defined, then we get a 404 page.
    
	// Below, we will create our very own result from what we've learned.
    r.HandleFunc("/{crypto}/trade/{pair}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		crypto := vars["crypto"]
		pair := vars["pair"]
		
        fmt.Fprintf(w, "You're trying to trade %s with %s pair", crypto, pair)			
	})
	/*
	So now, this program accepts 2 sets of directory.
	
	First is domain.com/{string}/garlic-juice/pages/{int} which prints a sample if URL is http://localhost/books/garlico/page/12 below.
	Prints: You've requested the book: garlico on page 12
	
	Second is domain.com/{string}/trade/{string} which prints a sample if URL is http://localhost/chainlink/trade/link-usdt below.
	Prints: You're trying to trade chainlink with link-usdt pair
	
	For everything else, it prints a "404 page not found"
	
	See more in https://gowebexamples.com/routes-using-gorilla-mux/ for the rest of the gorilla/mux features.
	*/
	
	http.ListenAndServe(":80", r)	// this has 2 syntax, the first one is which port to use and the second one is the DefaultServeMux (which is Golang's default server)
}