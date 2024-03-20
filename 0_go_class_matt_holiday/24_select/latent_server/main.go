package main

import (
	"fmt"
	"net/http"
	"time"
)

func myHandler(w http.ResponseWriter, req *http.Request) {
	time.Sleep(time.Duration(4) * time.Second)
	fmt.Fprintln(w, "OK from latent_server")
}

func main() {
	http.HandleFunc("/", myHandler)
	http.ListenAndServe("localhost:8080", nil)
}
