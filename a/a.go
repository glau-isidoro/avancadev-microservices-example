package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", home)
	http.ListenAndServe(":9090", nil)

}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello, World</h1>")
}
