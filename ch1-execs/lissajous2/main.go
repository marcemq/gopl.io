// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 15.
//!+main

// Lissajous2 generates GIF animations of random Lissajous figures
// using RGB color's palette for each animation.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var myPalette = []color.Color{color.White, color.RGBA{232, 163, 178, 255},
	color.RGBA{132, 188, 213, 255}, color.RGBA{194, 203, 147, 255},
	color.RGBA{255, 205, 42, 255}, color.RGBA{173, 90, 255, 255},
	color.RGBA{23, 123, 159, 255}, color.RGBA{255, 0, 0, 255}}

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 20    // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, myPalette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			rand.Seed(time.Now().UnixNano())
			randIndex := uint8(rand.Intn(len(myPalette)))
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				randIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

//!-main
