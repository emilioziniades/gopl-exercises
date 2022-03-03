package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

var rootColour = map[complex128]color.RGBA{
	1:   {255, 0, 0, 255},   //red
	-1:  {255, 255, 0, 255}, //yellow
	1i:  {0, 0, 255, 255},   //blue
	-1i: {0, 255, 0, 255},   //green
}

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
			img.Set(px, py, newton(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//    = z - (z^4 - 1) / (4 * z^3)
//    = z - (z - 1/z^3) / 4

// f(x) = 0 when x = 1; x = -1; x = i; x = -i
func newton(z complex128) color.RGBA {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if r := cmplx.Abs(z*z*z*z - 1); r < 1e-6 {
			colour := rootColour[rootRounded(z)]
			return shade(colour, contrast*i)
		}
	}
	return color.RGBA{0, 0, 0, 255}
}

func rootRounded(z complex128) complex128 {
	a := int(math.Round(real(z)))
	b := int(math.Round(imag(z)))
	return complex(float64(a), float64(b))
}

func shade(col color.RGBA, contrast uint8) color.RGBA {
	newColor := color.RGBA{0, 0, 0, 255}
	if col.R == 255 {
		newColor.R = 255 - contrast
	}
	if col.G == 255 {
		newColor.G = 255 - contrast
	}
	if col.B == 255 {
		newColor.B = 255 - contrast
	}
	return newColor
}
