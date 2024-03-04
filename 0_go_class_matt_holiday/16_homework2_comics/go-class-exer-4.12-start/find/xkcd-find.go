package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

// { "month":      "4",
//   "day":        "20"
//   "year":       "2009",
//   "num":        571,
//   . . .
//   "transcript": "[[Someone is in bed, . . . long int.",
//   "img":        "https://imgs.xkcd.com/comics/cant_sleep.png",
//   "title":      "Can't Sleep",
// }

type xkcd struct {
	Num        int    `json:"num"`
	Day        string `json:"day"`
	Month      string `json:"month"`
	Year       string `json:"year"`
	Title      string `json:"title"`
	Transcript string `json:"transcript"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "no file given")
		os.Exit(-1)
	}

	var (
		input     io.ReadCloser
		items     []xkcd
		terms     []string
		found_cnt int
		err       error
	)

	fn := os.Args[1]
	input, err = os.Open(fn)

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(-1)
	}

	defer input.Close()

	for _, t := range os.Args[2:] {
		terms = append(terms, strings.ToLower(t))
	}

	if len(terms) < 1 {
		fmt.Fprintln(os.Stderr, "no search terms")
		os.Exit(-1)
	}

	// some code here

	err = json.NewDecoder(input).Decode(&items)

	if err != nil {
		fmt.Fprintln(os.Stderr, "issue when unmarshalling", err)
		os.Exit(-1)
	}

	fmt.Println("read", len(items), "comics")

	// select comics with all terms
	for _, comic := range items {
		var all_terms bool = true
		for _, term := range terms {
			if !strings.Contains(strings.ToLower(comic.Title), strings.ToLower(term)) &&
				!strings.Contains(strings.ToLower(comic.Transcript), strings.ToLower(term)) {
				all_terms = false
			}
		}
		if all_terms {
			fmt.Printf("https://xkcd.com/%d/ %s/%s/%s \"%s\"\n", comic.Num, comic.Day, comic.Month, comic.Year, comic.Title)
			found_cnt++
		}
	}

	fmt.Fprintf(os.Stderr, "found %d comics\n", found_cnt)
}
