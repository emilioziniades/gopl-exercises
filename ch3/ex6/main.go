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
		yUp := (float64(py)+0.5)/height*(ymax-ymin) + ymin
		yDown := (float64(py)-0.5)/height*(ymax-ymin) + ymin

		for px := 0; px < width; px++ {
			// Image point (px, py) represents complex value z.
			xLeft := (float64(px)-0.5)/width*(xmax-xmin) + xmin
			xRight := (float64(px)+0.5)/width*(xmax-xmin) + xmin
			NW := mandelbrot(complex(xLeft, yUp))
			SW := mandelbrot(complex(xLeft, yDown))
			NE := mandelbrot(complex(xRight, yUp))
			SE := mandelbrot(complex(xRight, yDown))
			colorAvg := averageColour(NW, SW, NE, SE)
			img.Set(px, py, colorAvg)
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.RGBA {
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

func averageColour(NW, SW, NE, SE color.RGBA) color.RGBA {
	avgR := uint8((int(NW.R) + int(SW.R) + int(NE.R) + int(SE.R)) / 4)
	avgG := uint8((int(NW.G) + int(SW.G) + int(NE.G) + int(SE.G)) / 4)
	avgB := uint8((int(NW.B) + int(SW.B) + int(NE.B) + int(SE.B)) / 4)
	avgA := uint8((int(NW.A) + int(SW.A) + int(NE.A) + int(SE.A)) / 4)
	return color.RGBA{avgR, avgG, avgB, avgA}

}
