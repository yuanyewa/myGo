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
	"sync"
)

var mu sync.Mutex
var c int
var palette = []color.Color{color.Black, color.RGBA{255, 0, 0, 0xff}, color.RGBA{0, 255, 0, 0xff}, color.RGBA{0, 0, 255, 0xff}}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count/", count)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	c++
	mu.Unlock()
	// fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	lissa(w)
}
func count(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	c++
	mu.Unlock()
	fmt.Fprintf(w, "Count %d\n", c)
}
func lissa(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(math.Mod(float64(i), 3)))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
