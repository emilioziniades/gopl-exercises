package main

import "fmt"

func Example_three() {

	fmt.Println("EXAMPLE 3")
	var x IntSet
	x.Add(1)
	x.Add(9)
	x.Add(144)
	fmt.Println(x.Len())
	x.Add(23)
	fmt.Println(x.Len())
	fmt.Println(" ")

	fmt.Println(&x)
	x.Remove(9)
	fmt.Println(&x)
	x.Remove(22)
	fmt.Println(&x)
	x.Clear()
	fmt.Println(&x)

}

func Example_four() {

	fmt.Println("EXAMPLE 4")

	var x IntSet
	x.Add(20)
	x.Add(23)
	x.Add(500)
	fmt.Println(&x)
	y := x.Copy()
	fmt.Println(y)
	y.Remove(20)
	fmt.Println(&x)
	fmt.Println(y)
}

func Example_five() {
	fmt.Println("EXAMPLE 5")
	var x IntSet
	x.AddAll(1, 4, 16, 128)
	fmt.Println(&x)
	fmt.Println(x.Elems())
}

func main() {
	Example_three()
	Example_four()
	Example_five()
}
