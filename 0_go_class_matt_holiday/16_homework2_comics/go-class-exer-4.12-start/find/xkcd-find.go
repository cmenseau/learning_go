package main

import (
	"bufio"
	"encoding/json"
	"fmt"
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
		//input io.ReadCloser
		items []xkcd
		terms []string
		cnt   int
	)

	fn := os.Args[1]
	input, err := os.Open(fn)

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(-1)
	}

	for _, t := range os.Args[2:] {
		terms = append(terms, strings.ToLower(t))
	}

	if len(terms) < 1 {
		fmt.Fprintln(os.Stderr, "no search terms")
		os.Exit(-1)
	}

	// some code here

	var found_comics []xkcd
	var file_content []byte

	// for n, _ := input.Read(file_bytes); n != 0; n, _ = input.Read(file_bytes) {
	// 	fmt.Println(string(file_bytes[:50]))
	// 	fmt.Println(string(file_bytes[9950:]))

	// 	file_content = append(file_content, file_bytes...)
	// 	file_bytes = make([]byte, 0, 10000)
	// }

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		text := scanner.Text()
		file_content = append(file_content, []byte(text)...)
	}

	fmt.Println(len(file_content))

	err = json.Unmarshal(file_content, &items)

	if err != nil {
		fmt.Fprintln(os.Stderr, "issue when unmarshalling", err)
		os.Exit(-1)
	}

	//fmt.Println(items)

	fmt.Println("read", len(items), "comics")

	// select comics with all terms
	for _, comic := range items {
		var add bool = true
		for _, term := range terms {
			if !strings.Contains(strings.ToLower(comic.Title), strings.ToLower(term)) &&
				!strings.Contains(strings.ToLower(comic.Transcript), strings.ToLower(term)) {
				add = false
			}
		}
		if add {
			found_comics = append(found_comics, comic)
		}
	}

	for _, comic := range found_comics {
		fmt.Printf("https://xkcd.com/%d/ %s/%s/%s \"%s\"\n", comic.Num, comic.Day, comic.Month, comic.Year, comic.Title)
	}

	cnt = len(found_comics)

	fmt.Fprintf(os.Stderr, "found %d comics\n", cnt)
}
