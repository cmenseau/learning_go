package grep_line_prefix_control

import "testing"

func TestGetPrefix(test *testing.T) {
	subtests := []struct {
		filename string
		lpr      LinePrefixRequest
		exp_out  string
	}{
		{
			filename: "whatever.txt",
			lpr:      LinePrefixRequest{},
			exp_out:  "",
		},
		{
			filename: "whatever.txt",
			lpr:      LinePrefixRequest{WithFilename: true},
			exp_out:  "whatever.txt",
		},
		{
			filename: "whatever.txt",
			lpr:      LinePrefixRequest{WithFilename: false},
			exp_out:  "",
		},
	}

	for _, subtest := range subtests {
		var out = subtest.lpr.GetPrefix(subtest.filename)

		if out != subtest.exp_out {
			test.Errorf("wanted %s (filename=%s with request=%+v), got %s",
				subtest.exp_out,
				subtest.filename, subtest.lpr,
				out)
		}
	}
}
