package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// NOTE: don't do this in real life
type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

// http.Error
// http.StatusBadRequest
// conv str to float : strconv.ParseFloat
func (db database) add(w http.ResponseWriter, req *http.Request) {
	u := req.URL.Query()

	if !u.Has("item") || !u.Has("price") {
		str := fmt.Sprintln(w, "Error : need item and price to create")
		http.Error(w, str, http.StatusBadRequest)
		return
	}

	float_price, err := strconv.ParseFloat(u.Get("price"), 32)

	if err != nil {
		str := fmt.Sprintln(w, "Error : price not ok", err)
		http.Error(w, str, http.StatusBadRequest)
		return
	}

	item_name := u.Get("item")

	db[item_name] = dollars(float_price)
	fmt.Fprintf(w, "new item %s for %s\n", item_name, db[item_name])
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	u := req.URL.Query()

	if !u.Has("item") || !u.Has("price") {
		fmt.Fprintln(w, "Error : need item and price to update")
		return
	}

	float_price, err := strconv.ParseFloat(u.Get("price"), 32)

	if err != nil {
		fmt.Fprintln(w, "Error : price not ok", err)
		return
	}

	item_price := dollars(float_price)
	item_name := u.Get("item")

	db[item_name] = item_price
	fmt.Fprintf(w, "new price %s for %s\n", item_price, item_name)
}

func (db database) fetch(w http.ResponseWriter, req *http.Request) {

	var queried = strings.Split(req.URL.RawQuery, "=")

	if len(queried) < 2 {
		fmt.Fprintln(w, "Error : need item name to read")
		return
	}

	item_name := queried[1]
	price, exi := db[item_name]

	if exi {
		fmt.Fprintf(w, "item %s has price %s\n", item_name, price)
	} else {
		str := fmt.Sprintln(w, "No such item")
		http.Error(w, str, http.StatusNotFound)
	}

}

func (db database) drop(w http.ResponseWriter, req *http.Request) {
	var queried = strings.Split(req.URL.RawQuery, "=")

	if len(queried) < 2 {
		fmt.Fprintln(w, "Error : need item name to delete")
		return
	}

	item_name := queried[1]

	delete(db, item_name)
	fmt.Fprintf(w, "item %s deleted\n", item_name)
}

func main() {
	db := database{
		"shoes": 50,
		"socks": 5,
	}

	// NOTE that these are all method values
	// (closing over the object "db")

	http.HandleFunc("/list", db.list)
	http.HandleFunc("/create", db.add)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.drop)
	http.HandleFunc("/read", db.fetch)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
