// Respond URL path if was given and counting visited times (../count)
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
var count int
var palette = []color.Color{color.White, color.Black}

const whiteIndex = 0
const blackIndex = 1

func main() {
	/*
		Behind the scenes
		* the server runs the handler for each incoming request in a
		* separate goroutine so that it can server multiple req simultaneously

		`Lock` & `Unlock`
		* it was meant to avoid a bug which is called as "race condition"
		* we must ensure that at most one goroutine access the var at a time
		* #TODO might add more details when I fully grasp it
	*/
	http.HandleFunc("/", handler_)
	http.HandleFunc("/count", counter)

	// Optional
	http.HandleFunc("/pattern", func(w http.ResponseWriter, r *http.Request) {
		lissajous(w)
	})

	log.Fatal(http.ListenAndServe("localhost:7777", nil))
}

// Echo the Path component of the requested URL
func handler_(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	_, _ = fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// Echo the number of calls so far
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	_, _ = fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func lissajous(out io.Writer) {
	/*
		This function is optional, since we're just using it for demonstration
		that is the `lissajous` function can also be written to the HTTP client.
	*/
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
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(
				size+int(x*size+0.5),
				size+int(y*size+0.5),
				blackIndex,
			)
		}
		phase += 0.1

		anim.Image = append(anim.Image, img)
		anim.Delay = append(anim.Delay, delay)
	}

	_ = gif.EncodeAll(out, &anim)
}
