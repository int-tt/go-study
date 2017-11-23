package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var palette = []color.Color{
	color.Black,
	color.RGBA{0xff, 0x00, 0x00, 0xff},
	color.RGBA{0x00, 0xff, 0x00, 0xff},
	color.RGBA{0x00, 0x00, 0xff, 0xff},
	color.RGBA{0xff, 0xff, 0x00, 0xff},
	color.RGBA{0x00, 0xff, 0xff, 0xff},
}

type Params struct {
	Cycles  int
	Res     float64
	Size    float64
	Nframes int
	Delay   int
}

const (
	whiteIndex = 0
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	http.HandleFunc("/", lissajous)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func getFormValue(r *http.Request, p *Params) error {
	var err error
	v := r.FormValue("cycles")
	if len(v) != 0 {
		p.Cycles, err = strconv.Atoi(v)
		if err != nil {
			return err
		}
	}
	v = r.FormValue("res")
	if len(v) != 0 {
		res, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		p.Res = float64(res)
	}
	v = r.FormValue("size")
	if len(v) != 0 {
		size, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		p.Size = float64(size)
	}
	v = r.FormValue("nframes")
	if len(v) != 0 {
		p.Nframes, err = strconv.Atoi(v)
		if err != nil {
			return err
		}
	}
	v = r.FormValue("delay")
	if len(v) != 0 {
		p.Delay, err = strconv.Atoi(v)
		if err != nil {
			return err
		}
	}
	return nil
}
func lissajous(w http.ResponseWriter, r *http.Request) {
	p := Params{
		Cycles:  5,
		Res:     0.001,
		Size:    100.0,
		Nframes: 64,
		Delay:   8,
	}
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "FormParse Error: %v\n", err)
		return
	}
	err = getFormValue(r, &p)
	if err != nil {
		fmt.Fprintf(w, "getFormValue Error: %v\n", err)
		return
	}
	// if len(r.FormValue("cyles")) != 0 {
	// 	cycles, err = strconv.Atoi(r.Form.Get("cycles"))
	// 	if err != nil {
	// 		fmt.Fprintf(w, "FormParse Error: %v\n", err)
	// 		return
	// 	}
	// }

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: p.Nframes}
	phase := 0.0
	for i := 0; i < p.Nframes; i++ {
		rect := image.Rect(0, 0, int(2*p.Size+1), int(2*p.Size+1))
		img := image.NewPaletted(rect, palette)
		index := uint8(rand.Int()%(len(palette)-1) + 1)
		for t := 0.0; t < float64(p.Cycles)*2*math.Pi; t += p.Res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(int(p.Size+x*p.Size+0.5), int(p.Size+y*p.Size+0.5), index)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, p.Delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(w, &anim)
}
