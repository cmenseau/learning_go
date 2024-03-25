package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func closed_channels() {
	// unbuffered chan
	myChan := make(chan string)

	go func() {
		myChan <- "hi"
		fmt.Println("sent hi")
		myChan <- "bob!"
		fmt.Println("sent bob!")
		close(myChan)
		fmt.Println("closed chan")
		//myChan <- "panic!!!!!"
	}()

for_loop:
	for {
		select {
		case str, ok := <-myChan:
			if ok {
				fmt.Println("Received :", str)
			} else {
				fmt.Println("Read closed from channel")
				break for_loop
			}
		}
	}

	fmt.Println("............")

	myChan2 := make(chan string)

	go func() {
		myChan2 <- "hey"
		fmt.Println("sent hey")
		myChan2 <- "alice!"
		fmt.Println("sent alice!")
		close(myChan2)
		fmt.Println("closed chan")
		//myChan2 <- "panic!!!!!"
	}()

	// do not use !!!!!!!!!!!
	// for range myChan2 {
	//     fmt.Println("Received :", <-myChan2)
	// }
	// it will read twice on every loop tour
	// 1) once in the for with range !!!!!!!!!!!!!!!
	//    equivalent to : for _ := range myChan2
	// 2) once in the <-myChan2

	for str := range myChan2 {
		fmt.Println("Received :", str)
	}

	fmt.Println("............")

	myBufferedChan := make(chan string, 3)

	go func() {
		myBufferedChan <- "hey"
		myBufferedChan <- "alice!"
		myBufferedChan <- "hi"
		myBufferedChan <- "bob"
		close(myBufferedChan)
		fmt.Println("Closed myBufferedChan")
		//myBufferedChan <- "panic!!!!!"
	}()

	for str := range myBufferedChan {
		fmt.Println("Received :", str)
	}
}

func nil_channels() {
	myChan := make(chan string)

	go func() {
		myChan <- "hey"
		fmt.Println("sent hey")
		myChan <- "alice!"
		fmt.Println("sent alice!")
		myChan = nil
		fmt.Println("niled chan")
	}()

	// // deadlock
	// for str := range myChan {
	// 	fmt.Println("Received :", str)
	// }

	// // deadlock
	// for_loop:
	// 	for {
	// 		select {
	// 		case str, ok := <-myChan:
	// 			if ok {
	// 				fmt.Println("Received :", str)
	// 			} else {
	// 				fmt.Println("Read closed from channel")
	// 				break for_loop
	// 			}
	// 		}
	// 	}

	// // deadlock
	// for str := range myChan {
	// 	fmt.Println(str)
	// }
}

func signature_vs_mailbox() {

	type T struct {
		i byte
		b bool
	}

	send_and_update := func(in int, channel chan *T) {
		t := &T{i: byte(in)}
		channel <- t
		t.b = true
	}

	send_5 := func(channel chan *T) {
		for i := range 5 {
			send_and_update(i, channel)
		}
		close(channel)
	}

	myChan := make(chan *T)
	vs := make([]T, 5)
	go send_5(myChan)

	for i := range vs {
		vs[i] = *<-myChan
	}

	// print later
	for _, v := range vs {
		fmt.Println(v)
	}

	myBuffChan := make(chan *T, 10)
	vsBuff := make([]T, 5)

	go send_5(myBuffChan)

	for i := range vsBuff {
		vsBuff[i] = *<-myBuffChan
	}

	// print later
	for _, v := range vsBuff {
		fmt.Println(v)
	}
}

func counting_semaphore() {
	wg := new(sync.WaitGroup)

	worker := func(i int, ch chan string) {
		// acquire semaphore
		ch <- "start"
		wg.Add(1)
		defer func() { <-ch; wg.Done() }() // release semaphore
		sleep_time := rand.Int63n(6)
		time.Sleep(time.Duration(sleep_time) * time.Second)
		fmt.Printf("Goroutine n°%d finished after for %ds\n", i, sleep_time)
	}

	tasks := make(chan string, 10)

	go func() {
		for i := range 40 {
			go worker(i+1, tasks)
		}
	}()

	time.Sleep(1 * time.Second)
	wg.Wait()

	fmt.Println("work done")
}

// will not work without a wait group because we don't know which
// goroutine will finish last
// none of the goroutine can send on done channel !!!
func counting_semaphore_v2() {
	done := make(chan int)

	worker := func(i int, ch chan string) {
		// acquire semaphore
		ch <- "start"
		defer func() { <-ch }() // release semaphore
		sleep_time := rand.Int63n(6)
		time.Sleep(time.Duration(sleep_time) * time.Second)
		fmt.Printf("Goroutine n°%d finished after for %ds\n", i, sleep_time)
	}

	tasks := make(chan string, 10)

	func() {
		for i := range 40 {
			go worker(i+1, tasks)
		}
	}()
	close(tasks)

	<-done
	fmt.Println("work done")
}

func main() {
	// closed_channels()
	// nil_channels()
	// signature_vs_mailbox()
	//counting_semaphore()
	counting_semaphore_v2()
}
