package main

import (
	"reflect"
	"testing"
	"unicode/utf8"
)

type characterCount struct {
	counts  map[rune]int
	utflen  [utf8.UTFMax + 1]int
	invalid int
}

func TestCharcount(t *testing.T) {
	var tests = []struct {
		input string
		want  characterCount
	}{
		{
			"hello!",
			characterCount{
				map[rune]int{'h': 1, 'e': 1, 'l': 2, 'o': 1, '!': 1},
				[5]int{0, 6, 0, 0},
				0},
		},
		{
			"λιγο Ελλινικα",
			characterCount{
				map[rune]int{'λ': 3, 'ι': 3, 'γ': 1, 'Ε': 1, 'ν': 1, 'κ': 1, 'α': 1, 'ο': 1, ' ': 1},
				[5]int{0, 1, 12, 0},
				0},
		},
		{
			"\xc5\xca",
			characterCount{
				map[rune]int{},
				[5]int{0, 0, 0, 0},
				2},
		},
	}

	for _, test := range tests {
		counts, utflen, invalid := Charcount(test.input)

		if !reflect.DeepEqual(counts, test.want.counts) {
			t.Errorf("Charcount(%q) = %v, want %v", test.input, counts, test.want.counts)
		}
		if !reflect.DeepEqual(utflen, test.want.utflen) {
			t.Errorf("Charcount(%q) = %v, want %v", test.input, utflen, test.want.utflen)
		}
		if !reflect.DeepEqual(invalid, test.want.invalid) {
			t.Errorf("Charcount(%q) = %v, want %v", test.input, invalid, test.want.invalid)
		}
	}
}
