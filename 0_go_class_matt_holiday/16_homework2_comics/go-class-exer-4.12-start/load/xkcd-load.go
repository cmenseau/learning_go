package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

// we don't actually use the struct to unmarshall the JSON
// coming back from the server; we just dump it out as text

func getOne(i int) []byte {
	url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", i)
	resp, err := http.Get(url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "stopped reading: %s\n", err)
		os.Exit(-1)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// easter egg: #404 returns HTTP 404 - not found

		fmt.Fprintf(os.Stderr, "skipping %d: got %d\n", i, resp.StatusCode)
		return nil
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "bad body: %s\n", err)
		os.Exit(-1)
	}

	return body
}

func main() {
	var (
		output io.WriteCloser = os.Stdout
		err    error
		cnt    int
		fails  int
	)

	if len(os.Args) > 1 {
		output, err = os.Create(os.Args[1])

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}

		defer output.Close()
	} else {
		fmt.Fprintln(os.Stderr, "no file given")
		os.Exit(-1)
	}

	// the output will be in the form of a JSON array,
	// so add the brackets before and after

	fmt.Fprint(output, "[")
	defer fmt.Fprint(output, "]")

	// some code here

	// limit to 20 comics so that it doesn't take forever
	// start at 400 to work on num:404 : no comic case

	var start = 400
	for i := start; i-start <= 20; i++ {
		comic_data := getOne(i)

		if len(comic_data) != 0 {

			if i-start < 20 {
				comic_data = append(comic_data, byte(','))
			}
			io.Copy(output, bytes.NewBuffer(comic_data))
			cnt++
		} else {
			fails++
		}

	}

	fmt.Fprintf(os.Stderr, "read %d comics\n", cnt)
	fmt.Fprintf(os.Stderr, "read fail %d comics\n", fails)
}
