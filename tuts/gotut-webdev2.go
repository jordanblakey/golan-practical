package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Using template literals
	fmt.Fprintf(w, `
<h1>Hey there</h1>
<p>Hey there</p>
<p>...and simple!</p>
<p>You %s even add %s</p>
  `, "can", "<strong>variables</strong>")

	fmt.Fprintln(w, "\n<h1>Hey there</h1>")
	fmt.Fprintln(w, "<p>Hey there</p>")
	fmt.Fprintln(w, "<p>...and simple!</p>")
	fmt.Fprintf(w, "<p>You %s even add %s</p>\n", "can", "<strong>variables</strong>")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8000", nil)
}
