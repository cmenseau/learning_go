package main

import "fmt"

func main() {
	s := make([]func(), 4)

	for i := 0; i < 4; i++ {
		s[i] = func() {
			fmt.Printf("%d @ %p\n", i, &i)
		}
	}
	for j := 0; j < 4; j++ {
		s[j]()
	}
}

func two_closures_over_var() {
	var a = 1

	incr := func() {
		a++
	}

	square := func() {
		a *= a
	}

	incr()   // 2
	square() // 4
	incr()   // 5
	incr()   // 6
	fmt.Println(a)
}

func square() func() int {
	var x = 1

	return func() int {
		res := x * x
		x++
		return res
	}
}

func fib() func() int {
	a, b := 0, 1

	return func() int {
		a, b = b, a+b
		return b
	}
}

func simpleClosure() {
	var a = 10

	var inner func() = func() {
		fmt.Println(a)
	}
	a++

	inner()
	fmt.Println(a)
}
