// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"sync"
)

func main() {
	// Generate()
	GenerateParallell(false)
}
func Generate(save bool) {
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
	if save {
		png.Encode(os.Stdout, img) // NOTE: ignoring errors
	}
}

type point struct {
	px, py int
	c      color.Color
}

func GenerateParallell(save bool) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	points := make(chan point)
	var wg sync.WaitGroup

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			// Image point (px, py) represents complex value z.
			z := complex(x, y)
			wg.Add(1)
			go func(px, py int, z complex128) {
				defer wg.Done()
				points <- point{px, py, mandelbrot(z)}
			}(px, py, z)
		}
	}

	go func() {
		wg.Wait()
		close(points)
	}()

	for pt := range points {
		img.Set(pt.px, pt.py, pt.c)
	}
	if save {
		png.Encode(os.Stdout, img) // NOTE: ignoring errors
	}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
