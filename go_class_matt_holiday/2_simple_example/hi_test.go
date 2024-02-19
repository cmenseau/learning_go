package hi_test

import (
	"fmt"
	"hi"
	"testing"
)

func TestHi(test *testing.T) {
	var subtests = []struct {
		in       []string
		expected string
	}{
		{
			expected: "Hello World!",
		},
		{
			in:       []string{"Alice"},
			expected: "Hello Alice!",
		},
		{
			in:       []string{"Alice", "Bobby"},
			expected: "Hello Alice, Bobby!",
		},
	}

	for _, subtest := range subtests {
		fmt.Println(subtest)
		if hi.Say(subtest.in) != subtest.expected {
			test.Errorf("wanted %s (%v), got %s", subtest.expected, subtest.in, hi.Say(subtest.in))
		}
	}
}
