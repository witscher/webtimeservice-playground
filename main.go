package main

import (
	"fmt"
	"net/http"
	"time"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	fmt.Fprintf(w, t.String())

}

func main() {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8080", nil)
}
