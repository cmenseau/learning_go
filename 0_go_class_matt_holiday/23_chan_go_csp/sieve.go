package main

import "fmt"

func number_generator(limit int, c chan<- int) {
	for i := 2; i < limit; i++ {
		c <- i // will wait until channel is writeable
	}
	close(c)
}

func filter(in <-chan int, out chan<- int, prime int) {
	// out : chan to send all numbers not divisible by prime

	for nb := range in {
		if nb%prime != 0 {
			out <- nb
		}
	}
	close(out)
}

func sieve(limit int) {

	numbers := make(chan int)
	go number_generator(limit, numbers)

	for {
		nb, ok := <-numbers
		// numbers is no longer the raw list of numbers from generator
		// it's the current last channel of the pyramid of channels we're building
		if !ok {
			break
		}
		fmt.Println(nb)
		// nb is prime because it was not filtered
		// nb couldn't be divided by any of the previous numbers

		not_divisible_by_nb := make(chan int)
		go filter(numbers, not_divisible_by_nb, nb)
		numbers = not_divisible_by_nb
	}
}
