// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

var minZ, maxZ float64

func main() {
	// surface(f1)
	// surface(f2)
	surface(f3)

}

func surface(f func(float64, float64) float64) {
	maxZ, minZ = zMaxMin(xyrange, f)
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7; margin: auto' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, color := corner(i+1, j, f)
			bx, by, _ := corner(i, j, f)
			cx, cy, _ := corner(i, j+1, f)
			dx, dy, _ := corner(i+1, j+1, f)
			if areAnyInf(ax, ay, bx, by, cx, cy, dx, dy) {
				continue
			}
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='#%s' stroke='#%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color, color)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int, f func(float64, float64) float64) (float64, float64, string) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Determine colour based on z
	zRange := maxZ - minZ
	redShare := (z - minZ) / zRange
	blueShare := 1 - redShare
	red := 255 * redShare
	blue := 255 * blueShare
	red = math.Min(red, 255.00)
	blue = math.Max(blue, 0.00)
	color := (int(red) << 16) + int(blue)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, fmt.Sprintf("%06X", color)
}

func f1(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

// egg box
func f2(x, y float64) float64 {
	a := 0.15
	b := 0.05
	return a * (math.Sin(x/b) + math.Sin(y/b))
}

// saddle
func f3(x, y float64) float64 {
	c := 16.0
	d := 10.0
	return (math.Pow(x, 2) / math.Pow(c, 2)) - (math.Pow(y, 2) / math.Pow(d, 2))
}

func areAnyInf(x ...float64) bool {
	for _, e := range x {
		if math.IsInf(e, 0) || math.IsNaN(e) {
			return true
		}
	}
	return false
}

func zMaxMin(xyrange float64, f func(float64, float64) float64) (float64, float64) {
	minZ := math.Inf(1)
	maxZ := math.Inf(-1)
	for x := -xyrange; x < xyrange; x++ {
		for y := -xyrange; y < xyrange; y++ {
			z := f(x, y)
			if z < minZ {
				minZ = z
			}
			if z > maxZ {
				maxZ = z
			}
		}
	}
	return maxZ, minZ
}

//!-
