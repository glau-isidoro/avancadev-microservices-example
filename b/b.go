package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Result struct {
	Status string
}

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":9191", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	coupon := r.PostFormValue("coupon")
	ccNumber := r.PostFormValue("ccNumber")

	log.Println(coupon)
	log.Println(ccNumber)

	result := Result{Status: "declined"}

	jsonData, err := json.Marshal(result)
	if err != nil {
		log.Fatal("Error processing json")
	}

	fmt.Fprintf(w, string(jsonData))
}

// func makeHTTPCall(URLMicroservice string, coupon string, ccNumber string) Result {
// 	values := url.Values{}
// 	values.Add("coupon", coupon)
// 	values.Add("ccNumber", ccNumber)

// 	res, err := http.PostForm(URLMicroservice, values)
// 	if err != nil {
// 		return Result{Status: "Microservice B out"}
// 	}

// 	defer res.Body.Close()

// 	data, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		return Result{Status: "Error processing result"}
// 	}

// 	result := Result{}
// 	json.Unmarshal(data, &result)

// 	return result
// }
