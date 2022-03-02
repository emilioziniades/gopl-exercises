package main

import (
	"testing"
)

const format = "got %v, wanted %v"

func testOne(t *testing.T) {

	var x, y IntSet

	x.Add(1)
	x.Add(144)
	x.Add(9)

	if got, want := x.String(), "{1 9 144}"; got != want {
		t.Errorf(format, got, want)
	}

	y.Add(9)
	y.Add(42)

	if got, want := y.String(), "{9 42}"; got != want {
		t.Errorf(format, got, want)
	}

	x.UnionWith(&y)

	if got, want := x.String(), "{1 9 42 144}"; got != want {
		t.Errorf(format, got, want)
	}

	if got, want := x.Has(9), true; got != want {
		t.Errorf(format, got, want)
	}

	if got, want := x.Has(123), false; got != want {
		t.Errorf(format, got, want)
	}
}

func testFour(t *testing.T) {
	var x IntSet
	x.AddAll(1, 4, 16, 128)

	if got, want := x.String(), "{1 4 16 128}"; got != want {
		t.Errorf(format, got, want)
	}
}

func TestAll(t *testing.T) {
	testOne(t)
	testFour(t)
}
