package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func race_condition() {
	var i = 0
	incr := func() {
		i++
	}

	wg := sync.WaitGroup{}

	goroutine_nb, incr_nb := 100, 50

	for range goroutine_nb {
		wg.Add(1)
		go func() {
			for range incr_nb {
				incr()
			}
			wg.Done()
		}()
	}

	wg.Wait()
	if i != goroutine_nb*incr_nb {
		fmt.Printf("race condition : ")
	}
	fmt.Printf("i value=%d, expected %d\n", i, goroutine_nb*incr_nb)
}

func race_condition_mutex() {
	var i = 0

	mutex := sync.Mutex{}

	incr := func() {
		mutex.Lock()
		defer mutex.Unlock()
		i++
	}

	wg := sync.WaitGroup{}

	goroutine_nb, incr_nb := 100, 50

	for range goroutine_nb {
		wg.Add(1)
		go func() {
			for range incr_nb {
				incr()
			}
			wg.Done()
		}()
	}

	wg.Wait()
	if i != goroutine_nb*incr_nb {
		fmt.Printf("race condition : ")
	}
	fmt.Printf("i value=%d, expected %d\n", i, goroutine_nb*incr_nb)
}

func race_condition_semaphore() {
	ch := make(chan int, 1)
	var i = 0
	incr := func() {
		ch <- 1
		i++
		<-ch
	}

	wg := sync.WaitGroup{}

	goroutine_nb, incr_nb := 100, 50

	for range goroutine_nb {
		wg.Add(1)
		go func() {
			for range incr_nb {
				incr()
			}
			wg.Done()
		}()
	}

	wg.Wait()
	if i != goroutine_nb*incr_nb {
		fmt.Printf("race condition : ")
	}
	fmt.Printf("i value=%d, expected %d\n", i, goroutine_nb*incr_nb)
}

func map_concurrent() {
	// without the mutex : fatal error: concurrent map writes

	type database struct {
		db    map[string]int
		db_mx sync.Mutex
	}

	var myDb = database{
		db:    make(map[string]int),
		db_mx: sync.Mutex{},
	}
	incr := func() {
		myDb.db_mx.Lock()
		myDb.db["whatever"]++
		myDb.db_mx.Unlock()
	}

	wg := sync.WaitGroup{}

	goroutine_nb, incr_nb := 100, 50

	for range goroutine_nb {
		wg.Add(1)
		go func() {
			for range incr_nb {
				incr()
			}
			wg.Done()
		}()
	}

	wg.Wait()
	if myDb.db["whatever"] != goroutine_nb*incr_nb {
		fmt.Printf("race condition : ")
	}
	fmt.Printf("db[\"whatever\"] value=%d, expected %d\n", myDb.db["whatever"], goroutine_nb*incr_nb)
}

func race_condition_atomic() {
	var i int64 = 0
	incr := func() {
		atomic.AddInt64(&i, 1)
	}

	wg := sync.WaitGroup{}

	var goroutine_nb, incr_nb int64 = 100, 50

	for range goroutine_nb {
		wg.Add(1)
		go func() {
			for range incr_nb {
				incr()
			}
			wg.Done()
		}()
	}

	wg.Wait()
	if i != goroutine_nb*incr_nb {
		fmt.Printf("race condition : ")
	}
	fmt.Printf("i value=%d, expected %d\n", i, goroutine_nb*incr_nb)
}

func sync_once() {

	var i = 0
	incr := func() {
		i++
	}

	var once = sync.Once{}
	once.Do(incr)
	once.Do(incr)
	once.Do(incr)

	fmt.Println(i)
}

func main() {
	// go run -race main.go
	// race_condition()
	race_condition_mutex()
	race_condition_semaphore()
	map_concurrent()
	race_condition_atomic()
	sync_once()
}
