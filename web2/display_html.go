package main

import (
    "fmt"
    "net/http"
)

const AddForm = `
<form method="POST" action="/sending-url">
URL: <input type="text" name="url">
<input type="submit" value="Add">
</form>
`

func main() {
    http.HandleFunc("/", Add)
    http.ListenAndServe(":80", nil)
}

func Add(w http.ResponseWriter, r *http.Request) {
    url := r.FormValue("url")
	if url == "" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, AddForm)
		return
	}
    fmt.Fprintf(w, "http://localhost:80/%s", url)
}