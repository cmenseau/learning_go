package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Post struct {
	UserId int
	Id     int
	Title  string
	Body   string
}

var form = `<h1>{{.Title}} (#{{.Id}})</h1>
<div>{{printf "User %d" .UserId}}</div>
<div>{{.Body}}</div>`

func main() {
	const url = "https://jsonplaceholder.typicode.com"

	resp, err := http.Get(url + "/posts/1")

	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode == http.StatusOK {
		var post Post
		err = json.NewDecoder(resp.Body).Decode(&post)

		if err != nil {
			log.Fatal(err)
		}

		var tmpl = template.New("myTemplate")
		tmpl.Parse(form)
		err = tmpl.Execute(os.Stdout, post)

		if err != nil {
			log.Fatal(err)
		}
	}

}
