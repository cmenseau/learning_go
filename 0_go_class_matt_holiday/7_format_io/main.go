package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	wc_like()
}

func wc_like() {
	var lt, wt, ct int
	for _, filename := range os.Args[1:] {
		file, err := os.Open(filename)

		var (
			wc, cc int
			lc     = -1
		)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		} else {

			scanner := bufio.NewScanner(file)

			for scanner.Scan() {
				line := scanner.Text()

				wc += len(strings.Split(line, " "))
				cc += len(line)
				lc += 1
			}
		}

		file.Close()
		fmt.Printf("%3d%4d%4d %s\n", lc, wc, cc, filename)
		wt += wc
		lt += lc
		ct += cc
	}

	if len(os.Args[1:]) > 1 {
		fmt.Printf("%3d%4d%4d total\n", lt, wt, ct)
	}
}

func cat_like() {
	for _, filename := range os.Args[1:] {

		file, err := os.Open(filename)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		} else {
			scanner := bufio.NewScanner(file)

			for scanner.Scan() {
				fmt.Println(scanner.Text())
			}
		}
		file.Close()
	}
}

func printf_formats() {
	var myMap = map[string]float32{"x": 3.498372, "y": 86.338497}
	var mySlice = []int8{34, -12, 67}
	// %v : value
	// %#v : value using Go-syntax (with type)
	fmt.Printf("%v\n%#[1]v\n%+[1]v\n", myMap)
	fmt.Printf("%v\n%#[1]v\n%+[1]v\n", mySlice)

	var myString = "Hello World!"
	fmt.Printf("%s %[1]q %[1]v %#[1]v %+[1]v\n", myString)

	fmt.Fprintln(os.Stderr, "error message")

	myFloat := 58.2799999
	myStringForFloat := fmt.Sprintf("%.2f", myFloat)
	fmt.Println(myStringForFloat)
}
