package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func deadlock_no_receiver() {
	var messages chan string = make(chan string)

	// fatal error: all goroutines are asleep - deadlock!
	// cannot send to messages because no one ready to receive

	// can't send on an unbuffered channel that does not have a ready receiver
	messages <- "hello"

	read := <-messages
	fmt.Println(read)
}

func deadlock_no_sender() {
	var messages chan string = make(chan string)

	// fatal error: all goroutines are asleep - deadlock!
	// can't read on an unbuffered channel that does not have a message yet

	read := <-messages
	messages <- "hello"
	fmt.Println(read)
}

func chan_basic_example() {

	// not ok : deadlock : can't read on unbuffered without any sender

	// var messages chan string = make(chan string)
	// <-messages
	// go func() { messages <- "hello" }()

	// ok
	var messages2 chan string = make(chan string)
	go func() { <-messages2 }()
	messages2 <- "hello"

	// ok
	var messages chan string = make(chan string)
	go func() { messages <- "hello" }()
	<-messages

	// not ok : deadlock : can't write on unbuffered without any receiver
	// var messages3 chan string = make(chan string)
	// messages3 <- "hello"
	// go func() { <-messages3 }()

	var messages3 chan string = make(chan string)
	go func() { messages3 <- "hello" }()
	go func() { <-messages3 }()

}

type result struct {
	url     string
	err     error
	latency time.Duration
}

func get(url string, c chan result) {
	start := time.Now()
	_, err := http.Get(url)

	time_get := time.Since(start)

	var res result

	res.url = url
	res.err = err
	res.latency = time_get

	c <- res
}

func concurrent_gets() {
	results := make(chan result)
	list := []string{"https://amazon.com", "https://google.com",
		"https://www.nytimes.com/international/",
		"https://www.lemonde.fr/", "https://www.health.com/"}

	for _, url := range list {
		go get(url, results)
	}

	for range list {
		res := <-results

		if res.err != nil {
			log.Default().Printf("%s %s\n", res.url, res.err)
		} else {
			log.Default().Printf("%s %s\n", res.url, res.latency)
		}
	}
}

func counter_request() {

	counter := make(chan int)

	var myCountHandler = func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "%d", <-counter)
	}

	go func() {
		for i := 0; ; i++ {
			counter <- i
		}
	}()

	http.HandleFunc("/", myCountHandler)
	http.ListenAndServe("localhost:8080", nil)
}

func channel_test() {
	myChannel := make(chan string)
	go func() {
		myChannel <- "start"
		time.Sleep(1 * time.Second)
		myChannel <- "one"
		time.Sleep(1 * time.Second)
		myChannel <- "two"
		time.Sleep(2 * time.Second)
		myChannel <- "four"
		close(myChannel)
	}()

	for {
		out, ok := <-myChannel
		if !ok {
			log.Println("Channel closed")
			break
		}
		log.Println(out)
	}
}

func main() {
	// deadlock_no_receiver()
	// deadlock_no_sender()
	//chan_basic_example()
	//concurrent_gets()
	//counter_request()
	// sieve(100)
	channel_test()
}
