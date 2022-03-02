// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 10

	//orange
	initial := color.RGBA{170, 90, 0, 255}
	//black
	target := color.RGBA{0, 0, 0, 255}

	nColors := float64(iterations / contrast)

	delta := struct {
		R, G, B, A float64
	}{
		float64(target.R-initial.R) / nColors,
		float64(target.G-initial.G) / nColors,
		float64(target.B-initial.B) / nColors,
		float64(target.A-initial.A) / nColors,
	}

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{
				uint8(initial.R + uint8(delta.R*float64(n))),
				uint8(initial.G + uint8(delta.G*float64(n))),
				uint8(initial.B + uint8(delta.B*float64(n))),
				uint8(initial.A + uint8(delta.A*float64(n))),
			}
		}
	}
	return target
}
