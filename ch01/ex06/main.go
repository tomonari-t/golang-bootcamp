package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	bgColorIndex   = 0
	lineColorIndex = 1
)

func lissajous(out io.Writer, colorPalette []color.Color) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, colorPalette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), lineColorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

func to16(numStr string) (int64, error) {
	return strconv.ParseInt(numStr, 16, 16)
}

func main() {
	r, _ := to16(os.Args[1])
	g, _ := to16(os.Args[2])
	b, _ := to16(os.Args[3])

	var lineColor = color.RGBA{uint8(r), uint8(g), uint8(b), 0xFF}
	var palette = []color.Color{color.Black, lineColor}
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout, palette)
}
