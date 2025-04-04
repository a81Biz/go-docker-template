package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from __MODULE_NAME__ on port 5235!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Listening on port 5235...")
	http.ListenAndServe(":5235", nil)
}
