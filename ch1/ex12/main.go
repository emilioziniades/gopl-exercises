package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w, r.URL.Query())
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	lissajous(os.Stdout, nil)
}

func lissajous(out io.Writer, query url.Values) {
	var params = map[string]float64{
		"cycles":  5,
		"res":     0.001,
		"size":    100,
		"nframes": 64,
		"delay":   8,
	}
	//overwrite defaults from url query map
	for k, v := range query {
		flt, err := strconv.ParseFloat(v[0], 64)
		if err == nil {
			params[k] = flt
		}
	}
	fmt.Println(params)
	var (
		cycles  = params["cycles"]  // number of complete x oscillator revolutions
		res     = params["res"]     // angular resolution
		size    = params["size"]    // image canvas covers [-size..+size]
		nframes = params["nframes"] // number of animation frames
		delay   = params["delay"]   // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: int(nframes)}
	phase := 0.0 // phase difference
	for i := 0; i < int(nframes); i++ {
		rect := image.Rect(0, 0, 2*int(size)+1, 2*int(size)+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(int(size)+int(x*size+0.5), int(size)+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, int(delay))
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
