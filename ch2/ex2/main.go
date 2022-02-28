package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := Fahrenheit(t)
		c := Celsius(t)
		kg := Kilogram(t)
		lb := Pound(t)
		m := Meter(t)
		ft := Feet(t)

		fmt.Printf("%s = %s, %s = %s\n", f, FToC(f), c, CToF(c))
		fmt.Printf("%s = %s, %s = %s\n", kg, KgToLb(kg), lb, LbToKg(lb))
		fmt.Printf("%s = %s, %s = %s\n", m, MToF(m), ft, FToM(ft))

	}

}
