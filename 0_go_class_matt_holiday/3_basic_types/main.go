package main

import "fmt"

func main() {
	var a int = 45678
	fmt.Printf("%v %[1]T\n", a)
	b := 1
	fmt.Printf("%v %[1]T\n", b)
	var c int
	fmt.Printf("%v %[1]T\n", c)
	var d int8 = 127
	// var d int8 = 128     // NumericOverflow
	fmt.Printf("%v %[1]T\n", d)

	var e float32 = 4.78
	fmt.Printf("%v %[1]T\n", e)
	var f = 4.79
	fmt.Printf("%v %[1]T\n", f)

	// Conversion
	// var g = d/f          // MismatchedTypes
	var g = float64(d) / f
	fmt.Printf("%v %[1]T\n", g)
	// var h = d / int(f)   // MismatchedTypes
	var h = d / int8(f)
	fmt.Printf("%v %[1]T\n", h)

	var i bool
	fmt.Printf("%v %[1]T\n", i)
	// j = bool(1) // InvalidConversion
	// j = bool(h) // InvalidConversion
	j := (true && false) || i
	fmt.Printf("%v %[1]T\n", j)
	j = b != c
	fmt.Printf("%v %[1]T\n", j)

	var k error
	fmt.Printf("%v %[1]T\n", k)

	var l struct {
		a int8
		b int16
		c bool
		d complex128
		e float32
	}
	fmt.Printf("%v %[1]T\n", l)
	l.e = 234.56789
	fmt.Printf("%v %[1]T\n", l)

	const m bool = true
	// m = false           // UnassignableOperand
	fmt.Printf("%v %[1]T\n", m)

	// won't compile : cannot have const not init
	// const n bool
	// fmt.Println(n)
}
