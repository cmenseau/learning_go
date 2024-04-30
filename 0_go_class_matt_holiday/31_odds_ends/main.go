package main

import "fmt"

func iota_test() {

	type matchGranularity int

	const (
		// default-value of matchGranularity because iota is 0
		AllGranularity matchGranularity = iota //0
		_                                      // 1
		// ignore comment, line jump

		WordGranularity // 2
		LineGranularity // 3
	)

	var gran matchGranularity
	fmt.Println(gran) // 0
	fmt.Println(gran == AllGranularity)

	// iota can be used with int, float
	// can't be used with string
	//type chair string
	type chair float64

	const (
		_              = 42
		Recliner chair = iota
		Sunbed
		Stool
		Armchair chair = iota
	)

	var ch chair
	fmt.Println(ch) // 0
	ch = Sunbed
	fmt.Printf("%#v\n", ch)         // 2
	fmt.Println(Sunbed == Armchair) // false

	type century int

	const (
		XVth century = 15 + iota
		XVIth
		XVIIth
		XVIIIth
		XIXth
		XXth
		XXIst
	)

	var era century
	fmt.Println(era)
	era = XXIst
	fmt.Println(era)

	type floor int

	const (
		secondBasement floor = iota - 2 // -2
		basement                        // -1
		_
		firstFloor  // 1
		secondFloor // 2
	)

	type shape int

	const (
		triangle  shape = 3
		square    shape = 4
		octogon   shape = 8
		dodecagon shape = 12
	)

	var sh shape = octogon
	fmt.Println(sh)

	type ByteSize int64

	const (
		_            = iota
		KiB ByteSize = 1 << (10 * iota) // 2^10
		MiB                             // 2^20
		GiB                             // 2^30
		TiB                             // 2^40
	)
}

func takes_variadic_args(vals ...any) {
	fmt.Printf("%T : %+[1]v\n", vals)
	// cannot use vals (variable of type []any) as []int value in argument to printInts
	// printInts(vals...)
}

func printInts(ints ...int) {
	for idx := range ints {
		fmt.Println(ints[idx])
	}
}

func var_arg_type() {

	printInts()
	// printInts(nil) // cannot use nil as int value
	printInts(nil...)
	printInts([]int{}...)
	printInts([]int{1, 2, 3}...)
	printInts(4, 5)

	takes_variadic_args([]int{1, 2, 3})
	// []interface {} : [[1 2 3]]
	takes_variadic_args([]any{1, true, 6.0, make(chan bool), make(map[int]int)})
	// []interface {} : [[1 true 6 0xc000102060 map[]]]

}

func main() {
	//iota_test()

	var_arg_type()
}
