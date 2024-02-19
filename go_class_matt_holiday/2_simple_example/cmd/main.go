package main

import (
	"fmt"
	"hi"
	"os"
)

func main() {
	fmt.Println(hi.Say(os.Args[1:]))
}
