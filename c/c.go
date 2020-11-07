package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Coupon struct {
	Code string
}

type Coupons struct {
	Coupon []Coupon
}

func (c Coupons) Check(code string) string {
	for _, item := range c.Coupon {
		if code == item.Code {
			return "valid"
		}
	}
	return "invalid"
}

type Result struct {
	Status string
}

var coupons Coupons

func main() {
	coupon1 := Coupon{Code: "WELLCOME"}
	coupon2 := Coupon{Code: "WEEKEND_DEAL"}
	coupon3 := Coupon{Code: "10OFF"}

	coupons.Coupon = append(coupons.Coupon, coupon1)
	coupons.Coupon = append(coupons.Coupon, coupon2)
	coupons.Coupon = append(coupons.Coupon, coupon3)

	http.HandleFunc("/", home)
	http.ListenAndServe(":9292", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	coupon := r.PostFormValue("coupon")
	valid := coupons.Check(coupon)

	result := Result{Status: valid}

	jsonResult, err := json.Marshal(result)
	if err != nil {
		log.Fatal("Error converting json")
	}

	fmt.Fprintf(w, string(jsonResult))
}
