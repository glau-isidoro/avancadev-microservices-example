package main

import (
	"html/template"
	"net/http"
)

func main() {

	http.HandleFunc("/", home)
	http.ListenAndServe(":9090", nil)

}

func home(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/checkout.html"))
	t.Execute(w, "")
}
