package main

// go run main.go < lorem_ipsum.txt

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	number_of_unique_words()
}

// printing 5 most common words of a text given in stdin
func number_of_unique_words() {
	var words map[string]int = make(map[string]int)

	scan := bufio.NewScanner(os.Stdin)
	scan.Split(bufio.ScanWords)

	for scan.Scan() {
		words[scan.Text()]++
	}

	// cannot sort map, so we need to sort another struct

	type kv struct {
		key   string
		value int
	}

	var sorted_words []kv

	for k, v := range words {
		sorted_words = append(sorted_words, kv{k, v})
	}

	sort.Slice(sorted_words, func(i, j int) bool {
		return sorted_words[i].value > sorted_words[j].value
	})

	for _, v := range sorted_words[:5] {
		fmt.Printf("%8v : %8d times\n", v.key, v.value)
	}
}

func array_slices_map_properties() {
	// Arrays
	a := [4]int{1, 2, 3, 4}
	fmt.Printf("%v %[1]T\n", a)
	b := [4]int{}
	fmt.Printf("%v %[1]T\n", b)
	var c = [4]int{1, 2, 3, 4}
	fmt.Printf("%v %[1]T\n", c)
	var d [4]int = [4]int{}
	fmt.Printf("%v %[1]T\n", d)
	var e [4]int
	fmt.Printf("%v %[1]T\n", e)
	// var f [4]int = {1,2,3,4} // syntax error
	var f [4]int = [4]int{1, 2, 3, 4}
	fmt.Printf("%v %[1]T\n", f)

	// use of ellipsis ...
	// var g [...]int = {2,3,4,5}
	var g = [...]int{2, 3, 4, 5}
	fmt.Printf("%v %[1]T\n", g)
	h := [...]int{1, 2, 3, 4}
	fmt.Printf("%v %[1]T\n", h)

	// Arrays
	var i []int
	fmt.Printf("%v %[1]T\n", i)
	var j = []string{"abc", "fer"}
	fmt.Printf("%v %[1]T\n", j)
	// not okay :
	// var k []int = {954, 653}

	// no i = {9,8,7,6,5,4,3} but :
	i = []int{9, 8, 7, 6, 5, 4, 3}

	// okay adding on Slice
	i = append(i, 2)
	fmt.Printf("%v %[1]T\n", i)
	// no adding on Array
	// h = append(h, 5)   // InvalidAppend

	// Assign Array to slice
	// i = h              // IncompatibleAssign
	i = h[:]
	fmt.Printf("%v %[1]T\n", i)

	// Assign slice to array
	// b = i              // IncompatibleAssign
	i = []int{9, 8, 7, 6, 5, 4, 3}
	b = [4]int(i) // only keep the first 4 int
	fmt.Printf("%v %[1]T\n", b)

	// len and cap on Arrays/Slices
	myArray := [...]int{1, 2, 3, 4}
	fmt.Printf("myArray %v, len=%d, cap=%d\n", myArray, len(myArray), cap(myArray))
	mySlice := []bool{true, false, true, true}
	fmt.Printf("myArray %v, len=%d, cap=%d\n", mySlice, len(mySlice), cap(mySlice))
	// make is only available on slices
	my2ndSlice := make([]bool, 3, 10)
	fmt.Printf("myArray %v, len=%d, cap=%d\n", my2ndSlice, len(my2ndSlice), cap(my2ndSlice))

	// pass by value, pass by reference

	// Maps

	var myMap = map[string]bool{"red": true, "blue": false, "green": true}
	fmt.Printf("%v %[1]T\n", myMap)

	myMap["red"] = false
	myMap["yellow"] = true
	fmt.Printf("%v %[1]T\n", myMap)

	var myIntMap = map[string]int{"dress": 3, "tshirt": 5, "socks": 6}
	fmt.Printf("%v %[1]T\n", myIntMap)

	myIntMap["bra"]++
	myIntMap["dress"]++
	myIntMap["socks"] = 3
	delete(myIntMap, "tshirt")
	fmt.Printf("%v %[1]T\n", myIntMap)

	var myNilMap map[string]bool
	fmt.Printf("%v %[1]T\n", myNilMap)
	// myNilMap["insert"] = false     // panic: assignment to entry in nil map

	myNilMap = make(map[string]bool)
	myNilMap["insert"] = true
	fmt.Printf("%v %[1]T\n", myNilMap)

	// Map lookup
	val, ok := myIntMap["bra"]
	fmt.Printf("myIntMap[\"bra\"] : val=%d, exists=%v\n", val, ok)
	val, ok = myIntMap["hat"]
	fmt.Printf("myIntMap[\"hat\"] : val=%d, exists=%v\n", val, ok)
}
