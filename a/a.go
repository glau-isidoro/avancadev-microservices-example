package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/process", process)
	http.ListenAndServe(":9090", nil)

}

func home(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/checkout.html"))
	t.Execute(w, "")
}

func process(w http.ResponseWriter, r *http.Request) {
	log.Println(r.FormValue("coupon"))
	log.Println(r.FormValue("cc-number"))

	t := template.Must(template.ParseFiles("templates/checkout.html"))
	t.Execute(w, "")
}
