// Sequence: 6
// Title: Assets and Files
// Source: https://gowebexamples.com/static-files/

package main

import "net/http"

func main() {
    fs := http.FileServer(http.Dir("assets/")) 					// this target directory is masked by a URL target in the Handle request below.
    http.Handle("/static/", http.StripPrefix("/static/", fs))	// here, we use the /static instead of /assets as we wanted to mask the name of the directory

    http.ListenAndServe(":8080", nil)
}

/* 
We created an index.html file in the /assets directory and we tried to use the pipeline {{.}} but since we didn't declared that page using template.Must like in templates.go, the appearance of the index.html appeared the way it is.
*/