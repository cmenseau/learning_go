package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type meaningEntry struct {
	Definitions []struct {
		Synonyms []string
	}
	Synonyms []string
}

type dictEntry struct {
	Meanings []meaningEntry
}

func extractSynonyms(jsonBody io.ReadCloser) ([]string, error) {
	var entries []dictEntry

	if err := json.NewDecoder(jsonBody).Decode(&entries); err != nil {
		return []string{}, err
	}

	var synonyms []string
	for _, entry := range entries {
		for _, meaning := range entry.Meanings {
			if len(meaning.Synonyms) != 0 {
				synonyms = append(synonyms, meaning.Synonyms...)
			}
		}
	}

	fmt.Printf("synonyms returned %s\n", synonyms)

	return synonyms, nil
}

// from debug pprof
// 38 @ 0x43e22e 0x407e6d 0x407ad7 0x6d505e 0x474e41
// #	0x6d505d	main.leakSomeGoroutines.func1+0x1d	/home/menseau/Documents/Go/learning_go/0_go_class_matt_holiday/36_profiler/my_server.go:48

func leakSomeGoroutines() {
	var waitFirst = make(chan bool)
	for range 20 {
		go func() {
			waitFirst <- true
		}()
	}
	<-waitFirst
}

func synonymHandler(w http.ResponseWriter, r *http.Request) {
	queries.Inc()
	word := r.PathValue("word")

	if word == "leak" {
		leakSomeGoroutines()
	}

	resp_api, err := http.Get("https://api.dictionaryapi.dev/api/v2/entries/en/" + word)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	body := resp_api.Body

	synonyms, err := extractSynonyms(body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, str := range synonyms {
		fmt.Fprintln(w, str)
	}
}

var queries = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "all_queries",
	Help: "How many queries we've received.",
})

func main() {
	prometheus.MustRegister(queries)
	fmt.Println("curl localhost:8080/synonym/leak to leak some goroutines")
	fmt.Println("hit localhost:8080/debug/pprof to see profiling data")
	fmt.Println("hit localhost:8080/metrics to see prometheus metrics, counter")

	http.HandleFunc("GET /synonym/{word}", synonymHandler)
	http.Handle("/metrics", promhttp.Handler())

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
