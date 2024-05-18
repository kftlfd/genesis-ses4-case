package main

import "testing"

func TestFormatRate(t *testing.T) {
	tests := []struct {
		rate int
		out  string
	}{
		{0, "0.00"},
		{1, "0.01"},
		{99, "0.99"},
		{100, "1.00"},
		{101, "1.01"},
		{110, "1.10"},
		{199, "1.99"},
		{3900, "39.00"},
		{3950, "39.50"},
	}

	for _, testcase := range tests {
		if res := formatRate(testcase.rate); res != testcase.out {
			t.Errorf("rate: %v\texpected: %v\tgot: %v", testcase.rate, testcase.out, res)
		}
	}
}
