package main

func fibonacciRec(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	}

	return fibonacciRec(n-1) + fibonacciRec(n-2)
}

func fibonacciSeq(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	}

	prevprev := 0
	prev := 1
	for i := 2; i <= n; i++ {
		prevprev, prev = prev, prev+prevprev
	}
	return prev
}
