package utils

import (
	"strings"
	"testing"
)

func TestStringReader(t *testing.T) {
	testTable := []struct {
		input  string
		target string
	}{
		{
			input:  "   qwert 12, 13  \n",
			target: "qwert 12, 13",
		},
		{
			input:  "  dada da.12, ew/c\n",
			target: "dada da.12, ew/c",
		},
		{
			input:  "\n",
			target: "",
		},
		{
			input:  "     \n",
			target: "",
		},
	}

	for _, testCase := range testTable {
		in := strings.NewReader(testCase.input)

		line, _ := StringReader(in)
		if line != testCase.target {
			t.Errorf("expected %q, got %q", testCase.target, line)
		}
	}
}
