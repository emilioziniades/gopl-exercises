package main

import (
	"reflect"
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

func testTwo(t *testing.T) {

	var x IntSet
	x.Add(1)
	x.Add(9)
	x.Add(144)
	if got, want := x.Len(), 3; got != want {
		t.Errorf(format, got, want)
	}
	x.Add(23)
	if got, want := x.Len(), 4; got != want {
		t.Errorf(format, got, want)
	}

	x.Remove(9)
	x.Remove(22)

	if got, want := x.Len(), 3; got != want {
		t.Errorf(format, got, want)
	}

	x.Clear()

	if got, want := x.Len(), 0; got != want {
		t.Errorf(format, got, want)
	}
}

func testThree(t *testing.T) {

	var x IntSet

	x.Add(20)
	x.Add(23)
	x.Add(500)
	y := x.Copy()
	if got, want := y.String(), x.String(); got != want {
		t.Errorf(format, got, want)
	}

	y.Remove(20)

	if got, want := y.String(), "{23 500}"; got != want {
		t.Errorf(format, got, want)
	}

	if got, want := x.String(), "{20 23 500}"; got != want {
		t.Errorf(format, got, want)
	}
}

func testFour(t *testing.T) {
	var x IntSet
	x.AddAll(1, 4, 16, 128)

	if got, want := x.String(), "{1 4 16 128}"; got != want {
		t.Errorf(format, got, want)
	}
	if got, want := x.Elems(), []int{1, 4, 16, 128}; !reflect.DeepEqual(got, want) {
		t.Errorf(format, got, want)
	}
}

func TestAll(t *testing.T) {
	testOne(t)
	testTwo(t)
	testThree(t)
	testFour(t)
}
