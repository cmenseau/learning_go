package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type euro float32

func (e euro) String() string {
	return fmt.Sprintf("%.2f€", e)
}

func (e *euro) StringPtrReceiver() string {
	*e++
	return fmt.Sprintf("%.2f€", *e)
}

// cannot define new methods on non-local type int
// func (i int) String() string {
// 	return fmt.Sprintf("", i)
// }

type IntSlice []int // named-type

func (int_slice IntSlice) String() string {
	var strs []string
	for _, v := range int_slice {
		strs = append(strs, strconv.Itoa(v))
	}
	return "[" + strings.Join(strs, ";") + "]"
}

func (int_slice *IntSlice) StringPtrReceiver() string {
	var strs []string
	for _, v := range *int_slice {
		strs = append(strs, strconv.Itoa(v))
	}
	return "[" + strings.Join(strs, ";") + "]"
}

func basic_methods() {
	var price euro = 34.99
	price.String()
	euro.String(price) // this works too!
	fmt.Println(price)

	price.StringPtrReceiver()

	// with litterals !

	euro(3.7).String()
	IntSlice{1, 2, 3}.String()
	// doesn't work
	// (&euro(3.7)).StringPtrReceiver()
	(&IntSlice{1, 2, 3}).StringPtrReceiver()

	var s fmt.Stringer
	// s is an interface variable
	// you can assign to it anything satisfying the interface
	s = price
	fmt.Printf("%T %[1]v\n", s)
	// main.IntSlice [1;2;3]

	// var s2 fmt.Stringer = fct
	// won't compile
}

type myFuncType func(int, int) int

func (f myFuncType) execWith2(i int) int {
	return f(2, i)
}

type database map[string]int

func (db *database) make_in_method_on_ptr_receiver() {
	new_map := make(map[string]int, 5)
	new_map["attic"] = 1
	*db = database(new_map)
}

func (db database) make_in_method_on_val_receiver() {
	new_map := make(map[string]int, 5)
	new_map["garage"] = 1
	db = database(new_map)
}

func database_example() {
	db := database{"bathroom": 1, "bedroom": 3, "kitchen": 1}

	db.make_in_method_on_ptr_receiver()
	fmt.Println(db) // db updated

	db.make_in_method_on_val_receiver()
	fmt.Println(db) // db not updated
}

type ByteCounter int

func (b *ByteCounter) Write(p []byte) (int, error) {
	l := len(p)
	*b += ByteCounter(l) // cast is required
	return len(p), nil
}

func byte_counter() {
	f1, _ := os.Open("a.txt")

	var c ByteCounter
	f2 := &c

	n, _ := io.Copy(f2, f1)
	fmt.Println("copied", n, "bytes")
}

func main() {
	basic_methods()

	var fct myFuncType = func(i1, i2 int) int { return i1 + i2 }
	fct.execWith2(3)

	database_example()
	byte_counter()

}
