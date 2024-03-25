package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func cancel_context_timeout() {
	type result struct {
		url     string
		err     error
		latency time.Duration
	}

	// restrict write only
	var get = func(ctx context.Context, url string, ch chan<- result) {
		start := time.Now()

		req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

		if resp, err := http.DefaultClient.Do(req); err != nil {
			ch <- result{url, err, 0}
		} else {
			t := time.Since(start).Round(time.Millisecond)
			ch <- result{url, nil, t}
			resp.Body.Close()
		}
	}

	results := make(chan result)
	list := []string{"https://amazon.com", "https://google.com", "http://localhost:8080/"}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	for _, url := range list {
		go get(ctx, url, results)
	}

	for range list {
		r := <-results
		if r.err != nil {
			log.Printf("%-20s %s\n", r.url, r.err)
		} else {
			log.Printf("%-20s %s\n", r.url, r.latency)
		}
	}
}

func first_service_wins() {
	type result struct {
		url     string
		err     error
		latency time.Duration
	}

	// restrict write only
	var get = func(ctx context.Context, url string, ch chan<- result) {
		start := time.Now()

		req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

		if resp, err := http.DefaultClient.Do(req); err != nil {
			ch <- result{url, err, 0}
		} else {
			t := time.Since(start).Round(time.Millisecond)
			ch <- result{url, nil, t}
			resp.Body.Close()
		}
	}

	var first = func(ctx context.Context, urls []string) (*result, error) {
		results := make(chan result)

		for _, url := range urls {
			go get(ctx, url, results)
		}

		select {
		case r := <-results:
			fmt.Println("result received")
			return &r, nil
		case <-ctx.Done():
			// it is first's responsibility to check if parent context is done
			// ctx is a descendant of parent context, but it could be cancelled from above
			fmt.Println("parent context done")
			return nil, ctx.Err()
		}
	}

	// localhost answers after 4 seconds
	// 2 cases :
	// 1) timout before localhost could answer : select in first can read from <-ctx.Done()
	// 2) localhost has enough time : some results are received in results channel

	list := []string{"http://localhost:8080/"}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := first(ctx, list)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Printf("%s responsed first in %s\n", res.url, res.latency)
	}

}

func main() {
	//cancel_context_timeout()
	first_service_wins()
}
