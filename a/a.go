package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Result struct {
	Status string
}

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/process", process)
	http.ListenAndServe(":9090", nil)

}

func home(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/checkout.html"))
	t.Execute(w, Result{})
}

func process(w http.ResponseWriter, r *http.Request) {
	result := makeHTTPCall("http://localhost:9191", r.FormValue("coupon"), r.FormValue("cc-number"))

	t := template.Must(template.ParseFiles("templates/checkout.html"))
	t.Execute(w, result)
}

func makeHTTPCall(URLMicroservice string, coupon string, ccNumber string) Result {
	values := url.Values{}
	values.Add("coupon", coupon)
	values.Add("ccNumber", ccNumber)

	res, err := http.PostForm(URLMicroservice, values)
	if err != nil {
		log.Fatal("Microservice B out")
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error processing result")
	}

	result := Result{}
	json.Unmarshal(data, &result)

	return result
}
