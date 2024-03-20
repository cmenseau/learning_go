package main

import (
	"log"
	"net/http"
	"time"
)

func basic_multiplex_ex() {
	chans := []chan int{
		make(chan int),
		make(chan int),
	}

	send_every := func(interval int, out chan<- int) {
		for {
			time.Sleep(time.Duration(interval) * time.Second)
			out <- interval
		}
	}

	// can also use i := range chans
	for i := 1; i < 3; i++ {
		go send_every(i, chans[i-1])
	}

	for i := 0; i < 10; i++ {
		select {
		case in := <-chans[0]:
			log.Println(in)
		case in := <-chans[1]:
			log.Println(in)
		}
	}
}

func multiplex_ex_closing_channels() {
	chans := []chan int{
		make(chan int),
		make(chan int),
	}

	send_5 := func(interval int, out chan<- int) {
		for range 5 {
			time.Sleep(time.Duration(interval) * time.Second)
			out <- interval
		}
		close(out)
	}

	for i := range chans {
		go send_5(i+1, chans[i])
	}

	// var finished1, finished2 bool

	for {
		select {
		case in, ok := <-chans[0]:
			if !ok {
				chans[0] = nil
				// finished1 = true
			} else {
				log.Println(in)
			}
		case in, ok := <-chans[1]:
			if !ok {
				chans[1] = nil
				// finished2 = true
			} else {
				log.Println(in)
			}
		}
		if chans[0] == nil && chans[1] == nil {
			break
		}
		// if finished1 && finished2 {
		// 	break
		// }
	}
}

func get_channels_ex1() {
	type result struct {
		url     string
		err     error
		latency time.Duration
	}

	// restrict write only
	var get = func(url string, ch chan<- result) {
		start := time.Now()

		if resp, err := http.Get(url); err != nil {
			ch <- result{url, err, 0}
		} else {
			t := time.Since(start).Round(time.Millisecond)
			ch <- result{url, nil, t}
			resp.Body.Close()
		}
	}

	stopper := time.After(3 * time.Second) // returns a chan, sends a message when time's up
	results := make(chan result)
	list := []string{"https://amazon.com", "https://google.com", "http://localhost:8080/"}

	for _, url := range list {
		go get(url, results)
	}

	for range list {
		select {
		case r := <-results:
			if r.err != nil {
				log.Printf("%-20s %s\n", r.url, r.err)
			} else {
				log.Printf("%-20s %s\n", r.url, r.latency)
			}
		case <-stopper:
			log.Fatal("timeout reached")
		}
	}
}

func ticker() {
	// Print every 2 seconds, stop after 5 prints

	const rate = 2 * time.Second
	ticker := time.NewTicker(rate)
	stopper := time.After(5 * rate)

outer:
	for {
		select {
		case <-ticker.C:
			log.Println("tick")
		case <-stopper:
			log.Println("STOP !")
			break outer
		}
	}
}

func main() {
	// basic_multiplex_ex()
	// multiplex_ex_closing_channels()
	// get_channels_ex1()
	ticker()
}
