package main
import (
    "net/http"

)
func main() {
    http.HandleFunc("/markdown", GenerateMarkdown)
    http.Handle("/", http.FileServer(http.Dir("public")))
    http.ListenAndServe(":8080", nil)

}
func GenerateMarkdown(rw http.ResponseWriter, r *http.Request) {
}
