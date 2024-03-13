package main

import (
	"bytes"
	"fmt"
	"io"
)

func nil_val_interface_var() {
	var a io.Reader
	var b *bytes.Buffer

	fmt.Printf("a (val=%v, type=%[1]T)\n", a)
	fmt.Printf("b (val=%v, type=%[1]T)\n", b)
	// a (val=<nil>, type=<nil>)
	// b (val=<nil>, type=*bytes.Buffer)

	fmt.Println(a == nil)
	fmt.Println(b == nil)

	a = b

	fmt.Printf("a (val=%v, type=%[1]T)\n", a)
	fmt.Printf("b (val=%v, type=%[1]T)\n", b)
	// a (val=<nil>, type=*bytes.Buffer)
	// b (val=<nil>, type=*bytes.Buffer)

	fmt.Println(a == nil)
	fmt.Println(b == nil)

}

type anything interface {
	doSth()
	doSthPtr()
}
type my_struct struct{}

func (a my_struct) doSth() {
	fmt.Println("doSth")
}
func (a *my_struct) doSthPtr() {
	fmt.Println("doSthPtr")
}

func nil_val_interface_var_v2() {
	var a anything
	var b *my_struct

	fmt.Println(a == nil)
	fmt.Println(b == nil)
	// a.doSthPtr()
	// panic: runtime error: invalid memory address or nil pointer dereference

	a = b

	fmt.Println(a == nil)
	fmt.Println(b == nil)

	// a.doSth()
	// panic: value method main.my_struct.doSth called using nil *my_struct pointer
	// b.doSth()
	// panic: runtime error: invalid memory address or nil pointer dereference

	a.doSthPtr()
	b.doSthPtr()
}

type Point struct {
	x, y float32
}

func (p *Point) Add(x, y float32) {
	p.x, p.y = p.x+x, p.y+y
}
func (p Point) OffsetOf(p1 Point) (x, y float32) {
	x, y = p.x-p1.x, p.y-p1.y
	return
}

func pointer_value_receivers() {
	p1 := new(Point) // *Point (0,0)
	p2 := Point{1, 1}

	p1.OffsetOf(p2) // compiler : (*p1).OffsetOf(p2)  -> value receiver
	p2.Add(3, 4)    // compiler : (&p2).Add(3,4)      -> ptr receiver

	fmt.Println(p2)
}

func main() {
	//nil_val_interface_var()
	// nil_val_interface_var_v2()
	pointer_value_receivers()
}
