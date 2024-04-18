package main

import "fmt"

func main() {
	fmt.Println(fibonacciRec(10))
	fmt.Println(fibonacciSeq(10))

	lst := mkList(10)
	fmt.Println(sumList(lst))

	sl := mkSlice(10)
	fmt.Println(sumSlice(sl))
}
