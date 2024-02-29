package main

import "fmt"

func main() {
	reference_to_loop_var_iterator2_fixed_in_go1_22()
}

func reference_to_loop_var_iterator2_fixed_in_go1_22() {
	type myStruct struct {
		age  int
		name string
	}
	type myStructPtr struct {
		age  *int
		name *string
	}
	var params = []myStruct{{age: 25, name: "Alice"}, {age: 23, name: "Bob"}}
	var result []myStructPtr

	for _, change := range params {
		fmt.Printf("%p\n", &change)
		result = append(result, myStructPtr{&change.age, &change.name})
		// WRONG !
	}

	fmt.Println(result)
}

func class_ex_okay_on_go1_22() {
	items := [][2]byte{{1, 2}, {3, 4}, {5, 6}}
	a := [][]byte{}

	for _, item := range items {
		a = append(a, item[:])
	}

	fmt.Println(items) // [[1 2] [3 4] [5 6]]
	fmt.Println(a)
	// [[1 2] [3 4] [5 6]]  // Go 1.22
	// [[5 6] [5 6] [5 6]]  // old Go
}

func reference_to_loop_var_iterator_okay_on_go1_22() {

	// with array

	array := [][2]int{{0, 1}, {1, 3}, {3, 5}}
	var arrayBadCopy [][2]int
	var arrayCopy [][2]int
	var arrayCopy2 [][]int

	for _, v := range array {
		arrayBadCopy = append(arrayBadCopy, v)
	}

	for _, v := range array {
		v2 := v
		arrayCopy = append(arrayCopy, v2)
	}

	for _, v := range array {
		arrayCopy2 = append(arrayCopy2, v[:])
	}

	fmt.Println(arrayBadCopy)
	fmt.Println(arrayCopy)
	fmt.Println(arrayCopy2)

	// with slice

	slice := [][]int{{0, 1}, {1, 3}, {3, 5}}
	var sliceBadCopy [][]int
	var sliceCopy [][]int

	for _, v := range slice {
		sliceBadCopy = append(sliceBadCopy, v)
	}

	for _, v := range slice {
		v2 := v
		sliceCopy = append(sliceCopy, v2)
	}

	fmt.Println(sliceBadCopy)
	fmt.Println(sliceCopy)
}

func dangerous_ptr_on_slice_struct_elem() {
	type user struct {
		name  string
		count int
	}

	addTo := func(u *user) { u.count++ }

	users := []user{{"alice", 0}, {"bob", 0}}
	alice := &users[0] // RISKY !

	amy := user{"amy", 1}
	users = append(users, amy)

	addTo(alice)       // alice is likely a stale pointer
	fmt.Println(users) // so alice's count will be 0 (because alice was obsolete)
}

func dangerous_ptr_on_slice_elem() {
	//////// ok in this case - because no reallocation ////////

	sliceA := []int{0, 1, 2, 3, 4, 5}
	threeA := &sliceA[3]
	*threeA++
	fmt.Println(sliceA) // [0 1 2 4 4 5]

	//////// reallocation when adding elements : stale ptr ////////

	sliceB := []int{0, 1, 2, 3, 4, 5}
	threeB := &sliceB[3]

	sliceB = append(sliceB, 6, 7, 8)
	*threeB++ // doesn't work

	fmt.Println(sliceB) // [0 1 2 3 4 5 6 7]

	//////// reallocation when reordering elements : stale ptr ////////

	sliceC := []int{4, 5, 6, 0, 1, 2, 3}
	threeC := &sliceC[6]

	*threeC++ // works
	sliceC = append(sliceC[3:], sliceC[0:3]...)

	*threeC-- // doesn't work

	fmt.Println(sliceC) // [0 1 2 4 4 5 6]
}
