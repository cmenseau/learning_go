package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	regexp_can_panic()
}

func regexp_can_panic() {
	// won't compile : error parsing regexp: invalid nested repetition operator: `++
	// re := regexp.MustCompile("(x+x++y")

	fmt.Println("Enter malformed regexp : (x+x++y")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	re := regexp.MustCompile(text)

	fmt.Println(re.FindAllString("xxxxxxxxxxy", -1))
	// panic: regexp: Compile("(x+x++y\n"): error parsing regexp: invalid nested repetition operator: `++`
}

func catastrophic_backtracking_ok() {
	re := regexp.MustCompile("(x+x+)+y")
	fmt.Println(re.FindAllString("xxxxxxxxxxy", -1))
}

func replacing_several_phone_numbers() {
	re := regexp.MustCompile(`\(([0-9]{3})\) ([0-9]{3})-([0-9]{4})`)
	var txts = []string{"call me at (123) 456-7890 today",
		"personal:(098) 765-4321, professional: (232) 232-2323",
		"not a phone number : (12) 345-67"}
	var repl_txts []string

	for _, txt := range txts {
		repl_txts = append(repl_txts, re.ReplaceAllString(txt, "+1 $1-$2-$3"))
	}

	fmt.Printf("%#v\n", repl_txts)
}

func regexp_methods() {
	str := "abc bc abbbc bbc abcdefg xyz ac ab"
	re := regexp.MustCompile("a?(b+)c")
	fmt.Println("FindString", re.FindString(str))
	fmt.Println("FindAllString", re.FindAllString(str, -1))
	fmt.Println("FindAllStringIndex", re.FindAllStringIndex(str, -1))
	fmt.Println("FindAllStringSubmatch", re.FindAllStringSubmatch(str, -1))
	fmt.Println("FindAllStringSubmatchIndex", re.FindAllStringSubmatchIndex(str, -1))
	fmt.Println("ReplaceAllStringFunc", re.ReplaceAllStringFunc(str, strings.ToUpper))
	// Inside repl, $ signs are interpreted as in [Regexp.Expand]
	// --> placeholder values
	fmt.Println("ReplaceAllString", re.ReplaceAllString(str, "<replaced>"))
	// useful when replacing with $ values (placeholder values) which SHOULDN'T be expanded
	fmt.Println("ReplaceAllLiteralString", re.ReplaceAllLiteralString(str, "<replaced>"))
	fmt.Println("MatchString", re.MatchString(str))
}
