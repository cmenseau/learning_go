package main

import (
	"fmt"
	"log"
	"net/http"
)

func myHandler(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hey there!\n"))
	fmt.Fprintln(resp, "Another hi :)")
}

func main() {
	// curl http://localhost:8080/myEndpoint

	http.HandleFunc("/myEndpoint", myHandler)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
	// this line is required to keep the server up and running
}
