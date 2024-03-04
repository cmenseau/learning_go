package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	var resp, err = http.Get("http://localhost:8080/myEndpoint")

	if err != nil {
		fmt.Println("Error when getting content :", err)
	}

	fmt.Printf("%+v\n", resp)

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var content, err = io.ReadAll(resp.Body)

		if err != nil {
			fmt.Println("Error when reading body conent :", err)
		}

		fmt.Println(string(content))
	}
}
