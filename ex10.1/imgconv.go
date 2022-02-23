package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
)

var format string

func main() {
	flag.StringVar(&format, "out", "", "image output filetype, includes jpeg, png, and gif")
	flag.Parse()

	formats := []string{"jpeg", "png", "gif"}
	if !contains(formats, format) {
		log.Fatalf("imgconv: unrecognized filetype: %q\n", format)
	}
	if err := convert(os.Stdin, os.Stdout); err != nil {
		log.Fatalf("imgconv: conversion error: %s\n", err)
	}
}

func convert(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	fmt.Fprintf(os.Stderr, "Input image filetype: %s\n", kind)
	if err != nil {
		return err
	}
	fmt.Fprintf(os.Stderr, "Output image filetype: %s\n", format)
	return encodeFuncs[format](out, img)

}
func contains(arr []string, s string) bool {
	for _, e := range arr {
		if e == s {
			return true
		}
	}
	return false
}

func jpegEncode(w io.Writer, i image.Image) error {
	return jpeg.Encode(w, i, nil)
}

func pngEncode(w io.Writer, i image.Image) error {
	return png.Encode(w, i)
}

func gifEncode(w io.Writer, i image.Image) error {
	return gif.Encode(w, i, nil)
}

type EncodeFunc func(io.Writer, image.Image) error

var encodeFuncs = map[string]EncodeFunc{
	"jpeg": jpegEncode,
	"png":  pngEncode,
	"gif":  gifEncode,
}
