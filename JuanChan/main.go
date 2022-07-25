/*


 */

package main

import (
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

// Logging logs all requests with its path and the time it took to process
func Logging() Middleware {

	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {

			// Here, we're going to print some logs and save it as access.log and accessIP.log
			start := time.Now()

			addrs, err := net.InterfaceAddrs()
			if err != nil {
				os.Stderr.WriteString("Oops: " + err.Error() + "\n")
				os.Exit(1)
			}

			for _, a := range addrs {
				if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						saveIP, err := os.OpenFile("accessIP.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
						if err != nil {
							log.Println(err)
						}
						defer saveIP.Close()
						logIP := log.New(saveIP, "", log.LstdFlags)
						//logIP.Println(os.Stdout.WriteString(ipnet.IP.String() + "\n"))
						defer func() { logIP.Println(ipnet.IP.String()) }()
					}
				}
			}

			save, err := os.OpenFile("access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Println(err)
			}
			defer save.Close()
			logger := log.New(save, "", log.LstdFlags)
			logger.Println(r.URL.Path, time.Since(start))
			defer func() { log.Println(r.URL.Path, time.Since(start)) }()

			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}

// Method ensures that url can only be requested with a specific method, else returns a 400 Bad Request
func Method(m string) Middleware {

	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {

			// Do middleware things
			if r.Method != m {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}

// Chain applies middlewares to a http.HandlerFunc
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to JuanChan")
	fmt.Fprintln(w, "This page logs the IP address of the visitor.")
}

// We create the structs for the about.html page
type aboutHTML struct {
	PageTitle   string
	Author      string
	Description string
}

func main() {
	http.HandleFunc("/", Chain(Hello, Method("GET"), Logging()))

	//fs := http.FileServer(http.Dir("static/"))
	//http.Handle("/static/", http.StripPrefix("/static/", fs))

	// This is the URL handler for about
	tmplAbout := template.Must(template.ParseFiles("about.html"))
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		dataAbout := aboutHTML{
			PageTitle:   "About ImageBoard",
			Author:      "Miggy",
			Description: "This is a sample page written in Golang. The Title, Author and Description had variables in main.go and was passed on to about.html.",
		}
		tmplAbout.Execute(w, dataAbout)
	})

	/*
		The home page can log and save the access logs of all URLs
		But when it comes to handling /about, it's having trouble
	*/

	http.ListenAndServe(":80", nil)
}

/*

$ go run advanced-middleware.go
2017/02/11 00:34:53 / 0s

$ curl -s http://localhost:8080/
hello world

$ curl -s -XPOST http://localhost:8080/
Bad Request

*/
