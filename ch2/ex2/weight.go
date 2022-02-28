package main

import "fmt"

type Kilogram float64
type Pound float64

func (k Kilogram) String() string { return fmt.Sprintf("%gkg", k) }
func (p Pound) String() string    { return fmt.Sprintf("%glbs", p) }

func KgToLb(k Kilogram) Pound { return Pound(k / 0.45359237) }
func LbToKg(p Pound) Kilogram { return Kilogram(p * 0.45359237) }
