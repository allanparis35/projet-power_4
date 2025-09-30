
package main

import (
	"net/http"
	"projet-power_4/printcircle"
	"github.com/russross/blackfriday"
)

func main() {
	go play()
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/markdown", GenerateMarkdown)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8080", nil)
}

func GenerateMarkdown(rw http.ResponseWriter, r *http.Request) {
	markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
	rw.Write(markdown)
	
}
func play() {
	printcircle.PrintCircle()
}
