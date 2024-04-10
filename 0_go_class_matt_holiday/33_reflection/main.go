package main

import (
	"fmt"
	"reflect"
	"time"
)

func any_interface() {
	var a interface{}
	if time.Now().Day() == 9 {
		a = "today"
	} else {
		a = []int{1, 2}
	}
	fmt.Println(a)

	// aStr := a.(string) // may panic !!!!

	// aSli := a.([]int) // may panic !!!!
	// fmt.Println(aSli)

	aSli, ok := a.([]int)
	fmt.Println(aSli)
	fmt.Println(ok)

	if aSli, ok := a.([]int); ok {
		fmt.Println(len(aSli))
	}

	if aStr, ok := a.(string); ok {
		fmt.Println(len(aStr))
	}

	var b any = []int{3, 4, 5}
	bStr, ok := b.(string)
	if !ok {
		fmt.Printf("b to string not ok, yet bStr=%#v\n", bStr)
	}
	// when downcast not working (when ok = false)
	// val.(T) returns the default value of T

	var c any
	cStr, ok := c.(string)
	fmt.Printf("cStr=%#v, ok=%t\n", cStr, ok)
}

func switch_type() {
	var a any
	//a = "a string"
	//a = 23
	//a = 44.0

	switch a_t := a.(type) {
	case string:
		fmt.Printf("string value=%s\n", a_t)
	case int:
		fmt.Printf("int value=%d\n", a_t)
	default:
		fmt.Printf("other type %T=%[1]v\n", a_t)
	}

	// a "a string" -> string value="a string"
	// a 23         -> int value=23
	// a 44.0       -> other type float64=44
	// a nil        -> other type <nil>=<nil>

}

func reflect_deepeq() {
	type struct1 struct {
		vars []int
	}
	type struct2 struct {
		varsss []int
	}
	myStruct1 := struct1{vars: []int{}}
	myStruct2 := struct2{varsss: nil}

	if reflect.DeepEqual(myStruct1, myStruct2) {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}

	myPtr1 := &myStruct1
	var myPtr2 *int = nil

	if reflect.DeepEqual(myPtr1, myPtr2) {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}
}

func main() {
	//any_interface()
	//switch_type()
	reflect_deepeq()
}
