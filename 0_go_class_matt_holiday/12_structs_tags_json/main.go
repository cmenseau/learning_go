package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	struct_copy()
}

func struct_copy() {
	type MyStruct struct {
		MyInt int
	}

	var a = MyStruct{MyInt: 7}

	var aCopy = a

	fmt.Printf("a=%+v\n", a)
	fmt.Printf("aCopy=%+v\n", aCopy)

	a.MyInt++

	fmt.Printf("a=%+v\n", a)
	fmt.Printf("aCopy=%+v\n", aCopy)
}

func struct_compatibilty() {
	type album1 struct {
		title string
	}
	type album2 struct {
		title string
	}

	// anonymous struct
	var a1 = album1{
		"The White Album",
	}
	var a2 = album2{
		"The Black Album",
	}

	fmt.Println(a1, a2)
	// {The White Album} {The Black Album}

	// a1 = a2
	// won't compile
}

func anonymous_struct_compatibilty() {
	// anonymous struct
	var album1 = struct {
		title string
	}{
		"The White Album",
	}
	var album2 = struct {
		myTitle string
	}{
		"The Black Album",
	}

	fmt.Println(album1, album2)
	// {The White Album} {The Black Album}

	// album1 = album2
	// won't compile : different field names
}

func struct_printf() {
	type Passport struct {
		country         string
		number          string
		validity_period int
	}

	var myPassport = Passport{
		country:         "New Zealand",
		number:          "12345ABC",
		validity_period: 10,
	}

	fmt.Printf("%%v : %v\n", myPassport)
	fmt.Printf("%%#v : %#v\n", myPassport)
	fmt.Printf("%%+v : %+v\n", myPassport)
}

func nil_struct_not_allowed() {
	type Message struct {
		m string
	}

	var m Message

	fmt.Println(m)

	// won't compile : cannot use nil as Message value in assignment
	// m = nil
}

func json_omitempty() {
	type Response struct {
		Page  int      `json:"page"`
		Words []string `json:"words,omitempty"`
	}

	r := Response{Page: 1, Words: []string{"up", "in", "out"}}
	j, _ := json.Marshal(r)

	fmt.Println(string(j)) // j []byte

	var r2 Response
	_ = json.Unmarshal(j, &r2)

	r3 := Response{Page: 1}
	// won't ouptut words in json because of omitempty
	j3, _ := json.Marshal(r3)
	fmt.Println(string(j3))
}

func json_encode_decode() {

	// things to know :
	// - struct fields should start with a capital letter so that
	//   json lib can find them !!
	// - assignment between json tags and struct field is not case sensitive!
	// - struct tag is useful for different prop name
	//      -> for ex : MaxSize int `json:max_size`

	json_raw_str := `{
		"Header": "SVG Viewer",
		"items": [
			{"id": "Open"},
			{"id": "OpenNew", "label": "Open New"},
			null,
			{"id": "OriginalView", "label": "Original View"},
			{"id": "Quality"},
			null
		],
		"size":5,
		"height":6
	}`
	type Item = struct {
		Id    string
		Label string
	}

	type Response = struct {
		HeaDeR string
		Items  []Item
		Height int `json:"-"` // will unmarshal as 0, will ignore field in marshal
		Weight int
	}

	var r Response

	err := json.Unmarshal([]byte(json_raw_str), &r)
	if err != nil {
		fmt.Errorf("Unable to unmarshal from JSON due to %s", err)
	}

	fmt.Printf("Unmarshal : %v\n", r)

	res, err := json.Marshal(r)
	if err != nil {
		fmt.Errorf("Unable to marshal to JSON due to %s", err)
	}

	fmt.Printf("Marshal : %s\n", res)

	var myResponse = Response{
		HeaDeR: "abc",
		//Items:  nil,
		//Height: ,
		//Weight: ,
	}

	res, err = json.Marshal(myResponse)
	if err != nil {
		fmt.Errorf("Unable to marshal to JSON due to %s", err)
	}

	fmt.Printf("Marshal : %s\n", res)
}
