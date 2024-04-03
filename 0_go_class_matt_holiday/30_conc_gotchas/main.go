package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

func mutex_deadlock() {
	var m sync.Mutex
	done := make(chan bool)

	fmt.Println("START")

	go func() {
		m.Lock()
		// no unlock !
	}()

	go func() {
		time.Sleep(time.Second)

		m.Lock()
		defer m.Unlock()

		fmt.Println("SIGNAL")
		done <- true
	}()

	<-done
	fmt.Println("DONE")
}

func mutexes_order() {
	m1, m2 := sync.Mutex{}, sync.Mutex{}
	done := make(chan bool)

	fmt.Println("START")

	go func() {
		m1.Lock()
		defer m1.Unlock()
		time.Sleep(time.Second)
		m2.Lock()
		defer m2.Unlock()
		fmt.Println("SIGNAL1")
		done <- true
	}()

	go func() {
		m2.Lock()
		defer m2.Unlock()
		time.Sleep(time.Second)
		m1.Lock()
		defer m1.Unlock()
		fmt.Println("SIGNAL2")
		done <- true
	}()

	<-done
	fmt.Println("DONE1")
	<-done
	fmt.Println("DONE2")
}

func select_lost_msg() {
	var i int = 0
	nb_gen := func() int {
		i++
		return i
	}

	output := make(chan int, 3)

	for {
		x := nb_gen()

		select {
		case output <- x: // writing x to output chan
			// takes time
			time.Sleep(time.Second)
			fmt.Println("processing", x)
		default:
			fmt.Println("finished")
			return
		}
	}
}

// 2024/04/03 17:23:06 FINISHED PROCESSING 0
// 2024/04/03 17:23:07 FINISHED PROCESSING 1
// 2024/04/03 17:23:07 FINISHED PROCESSING 2
// 2024/04/03 17:23:08 FINISHED PROCESSING 3
// 2024/04/03 17:23:08 FINISHED PROCESSING 4
// 2024/04/03 17:23:08 FINISHED PROCESSING 5
// 2024/04/03 17:23:09 FINISHED PROCESSING 6
// 2024/04/03 17:23:09 FINISHED PROCESSING 7
// 2024/04/03 17:23:09 WORK DONE
// cannot process last outputs in the chan because done signal was badly used
// it stopped the process before all values could be processed
func select_terminate_premature_done() {

	timeout := time.After(10 * time.Second)
	output := make(chan int, 3)
	done := make(chan bool)

	go func() {
		for i := range 10 {
			output <- i
		}
		done <- true
	}()

	for {
		select {
		case r := <-output:
			// sleep for 1 to 5 sec
			time.Sleep(time.Duration((rand.Intn(500))) * time.Millisecond)
			log.Println("FINISHED PROCESSING", r)
		case <-done:
			log.Println("WORK DONE")
			return
		case <-timeout:
			log.Println("TIMED OUT")
			return
		}
	}
}

func select_terminate_closing_ok() {

	timeout := time.After(10 * time.Second)
	output := make(chan int, 3)

	go func() {
		for i := range 10 {
			output <- i
		}
		close(output)
	}()

	for {
		select {
		case r, ok := <-output:
			if !ok {
				log.Println("WORK DONE")
				return
			}
			// sleep for 1 to 5 sec
			time.Sleep(time.Duration((rand.Intn(500))) * time.Millisecond)
			log.Println("FINISHED PROCESSING", r)
		case <-timeout:
			log.Println("TIMED OUT")
			return
		}
	}
}

func main() {
	// mutex_deadlock()
	// mutexes_order()
	//select_lost_msg()
	// select_terminate_premature_done()
	select_terminate_closing_ok()
}
