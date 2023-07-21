package client

import (
	"reflect"
	"testing"
)

func TestNormalizeRequestBody(t *testing.T) {
	cases := []struct {
		input  string
		output interface{}
	}{
		{
			input:  `{"type":"test"}`,
			output: `{"type":"test"}`,
		},
		{
			input:  `payload=%7B%22type%22%3A%22test%22%7D`,
			output: `{"type":"test"}`,
		},
		{
			input:  `payload={"type":"test"}`,
			output: `{"type":"test"}`,
		},
		{
			input:  `type=test`,
			output: `{"type":"test"}`,
		},
	}

	for _, c := range cases {
		got, err := normalizeRequestBody(c.input)
		if err != nil {
			t.Errorf("[ERROR] normalizeRequestBody(%q): %s", c.input, err)
		}

		if !reflect.DeepEqual(string(*got), c.output) {
			t.Errorf("normalizeRequestBody(%q) == %q, want %q", c.input, *got, c.output)
		}
	}
}
