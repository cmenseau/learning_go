package main

// go run cmd/main.go it that < input.txt
// go run cmd/main.go Léonie Élise < input.txt
// case unsensitive

import (
	"fmt"
	"search_replace"
	"strings"
)

func main() {
	fmt.Println(search_replace.Perform())
}

func string_properties() {
	a := "allo"
	b := "allô"
	fmt.Printf("%v %[1]T\n", a)
	fmt.Printf("%v %[1]T\n", b)

	fmt.Printf("%s len is %d\n", a, len(a))
	fmt.Printf("%s len is %d\n", b, len(b))

	fmt.Printf("%s byte len is %d\n", a, len([]byte(a)))
	fmt.Printf("%s byte len is %d\n", b, len([]byte(b)))

	fmt.Printf("%s rune len is %d\n", a, len([]rune(a)))
	fmt.Printf("%s rune len is %d\n", b, len([]rune(b)))

	c := 'A'
	fmt.Printf("%v %[1]T\n", c) // 65 int32

	var d rune = 'A'
	fmt.Printf("%v %[1]T\n", d) // 65 int32

	var e byte = 'A'
	fmt.Printf("%v %[1]T\n", e) // 65 uint8

	f := "A"
	fmt.Printf("%v %[1]T\n", f) // 1 string

	// imutability

	var g string = "Alice & Bob"
	h := g[:5] // Alice
	i := g
	g = "Alicia & Bobby"

	fmt.Printf("%v\n", g)
	fmt.Printf("%v\n", h)
	fmt.Printf("%v\n", i)

	// g[5] = 'a'         // UnassignableOperand
	g += " & Cecilia"
	fmt.Printf("%v\n", g)

	strings.ToLower(g)
	fmt.Printf("%v\n", g) // Alicia & Bobby & Cecilia, string is not modified cause it's immutable

	g = strings.ToLower(g)
	fmt.Printf("%v\n", g) // alicia & bobby & cecilia
}
