package main

import (
	"fmt"
	"slices"
)

func main() {
	ijk_from_slice()
}

func ijk_from_slice() {
	a := []int{1, 2, 3}
	b := a[0:2:2]

	fmt.Printf("a[%p] = %[1]v\n", a)
	fmt.Printf("b[%p] = %[1]v\n", b)
	// a[ox12345] = [1 2 3]
	// b[ox12345] = [1 2]

	b[0] = 9

	fmt.Printf("a[%p] = %[1]v\n", a)
	fmt.Printf("b[%p] = %[1]v\n", b)
	// a[ox12345] = [9 2 3]
	// b[ox12345] = [9 2]

	b = append(b, 5)

	fmt.Printf("a[%p] = %[1]v\n", a)
	fmt.Printf("b[%p] = %[1]v\n", b)
	// a[ox12345] = [9 2 3]
	// b[ox12350] = [9 2 5]
}

func ij_ijk_equivalent() {
	tab := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	low := 4
	high := 7
	fmt.Println(tab[low:high:len(tab)], "cap=", cap(tab[low:high:len(tab)]))
	fmt.Println(tab[low:high], "cap=", cap(tab[low:high]))

	if slices.Equal(tab[low:high:len(tab)], tab[low:high]) {
		fmt.Println(true)
	}
}

func ijkop() {
	a := [4]int{1, 2, 3, 4}
	b1 := a[1:3:3]
	// won't compile
	// b2 := a[1:3:2] // invalid slice indices: 2 < 3compilerSwappedSliceIndices

	fmt.Println(b1, "len", len(b1), "cap", cap(b1))
	//fmt.Println(b2, "len", len(b2), "cap", cap(b2))

	c := b1[0:3]
	fmt.Println(c, "len", len(c), "cap", cap(c))

}

func ijop_unintuitive() {
	a := [4]int{1, 2, 3, 4}
	b := a[1:3]

	c := b[0:3] // WTF ? slice b has length of 2

	fmt.Println(b, len(b), cap(b))
	fmt.Println(c, len(c), cap(c))
	// [2 3]   2 3
	// [2 3 4] 3 3

}

func ij_slicing_compiler_checks() {

	a1 := [...]int{1}
	// won't compile : invalid argument: index 2 out of bounds [0:2]
	// b1 := a1[0:2]
	b1 := a1[0:1]

	fmt.Println("a1 =", a1) // a1 = [1]
	fmt.Println("b1 =", b1) // b1 = [1]

	// panic crash : slice bounds out of range [:2] with capacity 1
	// c1 := b1[0:2]
	// fmt.Println("c1 =", c1)

	// a3 := []int{1}
	// // panic crash : slice bounds out of range [:2] with capacity 1
	// b3 := a3[0:2]
	// fmt.Println("b3 =", b3)
}

func init_n_array_with_nplus_elems() {
	myArray := [4]int{5, 6, 7, 8}
	myLenOneArr := [1]int(myArray[:])

	fmt.Printf("myArray %v %[1]T cap=%d\n", myArray, cap(myArray))
	fmt.Printf("myLenOneArr %v %[1]T cap=%d\n", myLenOneArr, cap(myLenOneArr))

	// won't compile
	// myArr := [2]int{1, 2, 3} // index 2 is out of bounds (>= 2)
}

func ijop_type() {
	myArray := [4]int{5, 6, 7, 8}
	mySlicedArray := myArray[1:3] // [i:j] operator return value is by default a slice
	mySlicedArray2 := [2]int(myArray[1:3])

	fmt.Printf("myArray %v %[1]T cap=%d\n", myArray, cap(myArray))
	fmt.Printf("mySlicedArray %v %[1]T cap=%d\n", mySlicedArray, cap(mySlicedArray))
	fmt.Printf("mySlicedArray2 %v %[1]T cap=%d\n", mySlicedArray2, cap(mySlicedArray2))

	mySlicedArray = append(mySlicedArray, -1)

	fmt.Printf("myArray %v %[1]T cap=%d\n", myArray, cap(myArray))
	fmt.Printf("mySlicedArray %v %[1]T cap=%d\n", mySlicedArray, cap(mySlicedArray))
	fmt.Printf("mySlicedArray2 %v %[1]T cap=%d\n", mySlicedArray2, cap(mySlicedArray2))

	mySlicedArray = append(mySlicedArray, -2)

	fmt.Printf("myArray %v %[1]T cap=%d\n", myArray, cap(myArray))
	fmt.Printf("mySlicedArray %v %[1]T cap=%d\n", mySlicedArray, cap(mySlicedArray))
	fmt.Printf("mySlicedArray2 %v %[1]T cap=%d\n", mySlicedArray2, cap(mySlicedArray2))

}

// will cause panic
func insertnilmap() {
	var myMap map[string]int
	myMap["tomate"] = 5
}

// okay
func appendnilslice() {
	var s []int
	s = append(s, 5)
}
