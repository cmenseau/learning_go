package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	a := 10
	fmt.Println(a)
	a = -1
	defer printInt(a)
	a = 3
	fmt.Println(a)
}

func printInt(a int) {
	fmt.Println(a)
}

func print_param_update_table() {
	/// INT ///
	var myInt int = 1
	var myIntCopy int = myInt
	assignInt(myInt)
	var assignIntB bool = myInt != myIntCopy

	myInt = myIntCopy
	updateInt(myInt)
	var updateIntB bool = myInt != myIntCopy

	/// ARRRAY ///
	myArray := [2]int{100, 200}
	var myArrayCopy = myArray

	assignArray(myArray)
	var assignArrayB bool = myArray != myArrayCopy

	myArray = myArrayCopy
	updateArray(myArray)
	var updateArrayB bool = myArray != myArrayCopy

	/// STRING ///
	myString := "Hi"
	myStringCopy := myString
	assignString(myString)
	var assignStringB bool = myString != myStringCopy

	myString = myStringCopy
	makeString(myString)
	var makeStringB bool = myString != myStringCopy

	/// SLICE ///
	mySlice := []int{100, 200}
	var mySliceCopy = make([]int, len(mySlice))
	copy(mySliceCopy, mySlice) // !!!! beware of the copy !!! only len of mySliceCopy is available !!!!

	assignSlice(mySlice)
	var assignSliceB bool = !slices.Equal(mySlice, mySliceCopy)

	mySlice = make([]int, len(mySliceCopy))
	copy(mySlice, mySliceCopy)
	updateSlice(mySlice)
	var updateSliceB bool = !slices.Equal(mySlice, mySliceCopy)

	mySlice = make([]int, len(mySliceCopy))
	copy(mySlice, mySliceCopy)
	makeSlice(mySlice)
	var makeSliceB bool = !slices.Equal(mySlice, mySliceCopy)

	/// SLICE PTR ///
	var mySlicePtr *[]int = &[]int{100, 200}
	var mySlicePtrCopy *[]int = &[]int{100, 200}
	assignSlicePtr(mySlicePtr)
	var assignSlicePtrB bool = !slices.Equal(*mySlicePtr, *mySlicePtrCopy)

	*mySlicePtr = make([]int, len(*mySlicePtrCopy))
	copy(*mySlicePtr, *mySlicePtrCopy)
	updateSlicePtr(mySlicePtr)
	var updateSlicePtrB bool = !slices.Equal(*mySlicePtr, *mySlicePtrCopy)

	copy(*mySlicePtr, *mySlicePtrCopy)
	makeSlicePtr(mySlicePtr)
	var makeSlicePtrB bool = !slices.Equal(*mySlicePtr, *mySlicePtrCopy)

	/// MAP ///
	var myMap = map[int]int{1: 10, 2: 20}
	var myMapCopy = make(map[int]int, len(myMap))
	for k, v := range myMap {
		myMapCopy[k] = v
	}
	assignMap(myMap)
	var assignMapB bool = !maps.Equal(myMap, myMapCopy)

	myMap = make(map[int]int, len(myMapCopy))
	for k, v := range myMapCopy {
		myMap[k] = v
	}
	updateMap(myMap)
	var updateMapB bool = !maps.Equal(myMap, myMapCopy)

	myMap = make(map[int]int, len(myMapCopy))
	for k, v := range myMapCopy {
		myMap[k] = v
	}
	makeMap(myMap)
	var makeMapB bool = !maps.Equal(myMap, myMapCopy)

	/// RESULTS ///
	fmt.Println("true -> actual param val was updated by operation in func")
	fmt.Printf("|%7s|%11s|%13s|%14s|%13s|%14s|%11s|\n",
		"func op", "int updated", "array updated", "string updated", "slice updated", "*slice updated", "map updated")
	fmt.Printf("|%7s|%11v|%13v|%14v|%13v|%14v|%11v|\n",
		"++", updateIntB, updateArrayB, "-", updateSliceB, updateSlicePtrB, updateMapB)
	fmt.Printf("|%7s|%11v|%13v|%14v|%13v|%14v|%11v|\n",
		"=", assignIntB, assignArrayB, assignStringB, assignSliceB, assignSlicePtrB, assignMapB)
	fmt.Printf("|%7s|%11v|%13v|%14v|%13v|%14v|%11v|\n",
		"make", "-", "-", makeStringB, makeSliceB, makeSlicePtrB, makeMapB)
}

func assignInt(in int) {
	in = 12
}

func updateInt(in int) {
	in++
}

func assignArray(in [2]int) {
	in = [2]int{3}
}

func updateArray(in [2]int) {
	in[0]++
}

func assignString(in string) {
	// in has a different address than actual parameter

	// in[0] = "0" // not allowed : string is immutable

	// in += " there"
	// fmt.Printf("in %#v %#v\n", in, &in)

	in = "something else"
}

func makeString(in string) {
	// in has a different address than actual parameter

	// in[0] = "0" // not allowed : string is immutable

	// in += " there"
	// fmt.Printf("in %#v %#v\n", in, &in)

	in = string(make([]byte, 5))
}

func assignSlice(in []int) {

	// in[0] = 999
	// fmt.Printf("in %#v\n", in)

	// in[0]--
	// fmt.Printf("in %#v\n", in)

	// in = append(in, 300)
	// fmt.Printf("in %#v\n", in)

	in = []int{99, 88, 77}
}

func updateSlice(in []int) {
	in[0]--
}

func makeSlice(in []int) {
	in = make([]int, 5)
}

func assignSlicePtr(in *[]int) {

	// in[0] = 999
	// fmt.Printf("in %#v\n", in)

	// in[0]--
	// fmt.Printf("in %#v\n", in)

	// in = append(in, 300)
	// fmt.Printf("in %#v\n", in)

	*in = []int{99, 88, 77}
}

func updateSlicePtr(in *[]int) {
	(*in)[0]--
}

func makeSlicePtr(in *[]int) {
	*in = make([]int, 5)
}

func assignMap(in map[int]int) {
	in = map[int]int{0: 0, 1: 1}
}

func updateMap(in map[int]int) {
	in[0]--
}

func makeMap(in map[int]int) {
	in = make(map[int]int, 5)
}
