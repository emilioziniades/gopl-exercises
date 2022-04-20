// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package main

import (
	"math/rand"
	"testing"
	"time"
	"unicode"
)

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", true},
		{"Evil I did dwell; lewd did I live.", true},
		{"Able was I ere I saw Elba", true},
		{"été", true},
		{"Et se resservir, ivresse reste.", true},
		{"palindrome", false}, // non-palindrome
		{"desserts", false},   // semi-palindrome
	}
	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v", test.input, got)
		}
	}
}

func TestRandomPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%v) = false", p)
		}
	}
}

func TestRandomPalindromesWithPunct(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomPalindromeWithPunct(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%v) = false", p)
		}
	}
}

func TestRandomNonPalindromes(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		np := randomNonPalindrome(rng)
		if IsPalindrome(np) {
			t.Errorf("IsPalindrome(%v) = true", np)
		}
	}
}

func TestRandomNonPalindromesWithPunct(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		np := randomNonPalindromeWithPunct(rng)
		if IsPalindrome(np) {
			t.Errorf("IsPalindrome(%v) = true", np)
		}
	}
}

// randomPalindrome returns a palindrome whose length and contents
// are derived from the pseudo-random number generator rng.
func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := getRandomLetter(rng, 0x1000, 0)
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

// randomPalindromeWithPunct returns a palindrome whose length and contents
// are derived from the pseudo-random number generator rng. It includes punctuation
func randomPalindromeWithPunct(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := getRandomLetterOrPunct(rng, 0x1000, 0)
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

// For NonPalindromes (with or without Punct)
// split the unicode range into two equal parts [0, 0x800) and [0x800, 0x1000)
// for the first half of the string, draw from the first interval, and for
// the second half, draw from the second interval

func randomNonPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) + 2 // random length [2,27)
	runes := make([]rune, n)
	for i := 0; i < n; i++ {
		var r rune
		if i < n/2 {
			r = getRandomLetter(rng, 0x800, 0)
		} else {
			r = getRandomLetter(rng, 0x800, 0x800)
		}
		runes[i] = r
	}

	return string(runes)
}
func randomNonPalindromeWithPunct(rng *rand.Rand) string {
	n := rng.Intn(25) + 5 // random length [5,30) (extend minimum length to avoid strings with a single letter)
	runes := make([]rune, n)
	for i := 0; i < n; i++ {
		var r rune
		if i < n/2 {
			r = getRandomLetterOrPunct(rng, 0x800, 0)
		} else {
			r = getRandomLetterOrPunct(rng, 0x800, 0x800)
		}
		runes[i] = r
	}

	return string(runes)
}

//getRandomLetter and getRandomLetterOrPunct used in generating random palindromes (with or without punctuation) above.
// We want to ensure that the runes we randomly get are valid unicode, otherwise the length of the
// valid unicode characters in our random string is nondeterministic
// you could get a string which is a single rune long, and that would be a palindrome,
// so this is necessary to guarantee that the string is NOT a palindrome for NonPalindrome tests

func getRandomLetter(rng *rand.Rand, limit, shift int) rune {
	return getRandom(rng, limit, shift, unicode.IsLetter)
}

func getRandomLetterOrPunct(rng *rand.Rand, limit, shift int) rune {
	return getRandom(rng, limit, shift, unicode.IsLetter, unicode.IsPunct)
}
func getRandom(rng *rand.Rand, limit, shift int, checkFuncs ...func(r rune) bool) rune {
	var r rune
	for {
		r = rune(rng.Intn(limit) + shift)
		for _, check := range checkFuncs {
			if check(r) {
				return r
			}
		}
	}
}
