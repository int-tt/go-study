package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", render)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
func render(w http.ResponseWriter, r *http.Request) {
	var (
		scale  = 100
		xx, yy = 0, 0
	)
	var err error

	if v := r.FormValue("scale"); len(v) != 0 {
		scale, err = strconv.Atoi(v)
		if err != nil {
			fmt.Fprintf(w, "FormParse Error: %v\n", err)
			return
		}
	}

	if v := r.FormValue("x"); len(v) != 0 {
		xx, err = strconv.Atoi(v)
		if err != nil {
			fmt.Fprintf(w, "FormParse Error: %v\n", err)
			return
		}
	}

	if v := r.FormValue("y"); len(v) != 0 {
		yy, err = strconv.Atoi(v)
		if err != nil {
			fmt.Fprintf(w, "FormParse Error: %v\n", err)
			return
		}
	}
	var (
		xmin, ymin, xmax, ymax = -2 - xx, -2 - yy, +2 + xx, +2 + yy
		width, height          = 1024 * scale / 100, 1024 * scale / 100
	)
	img := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
	for py := 0; py < int(height); py++ {
		y := float64(py)/float64(height)*(float64(ymax)-float64(ymin)) + float64(ymin)
		for px := 0; px < int(width); px++ {
			x := float64(px)/float64(width)*(float64(xmax)-float64(xmin)) + float64(xmin)
			z := complex(x, y)
			// 画像の点(px,py)は複素数値zを表している。
			img.Set(px, py, newton(z))
		}
	}
	png.Encode(w, img)
}
func newton(z complex128) color.Color {
	const iterations = 37
	const constrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		switch {
		case cmplx.Abs((1+0i)-z) < 1e-6:
			return color.RGBA{255 - constrast*1, 0, 0, 0xff}
		case cmplx.Abs((-1+0i)-z) < 1e-6:
			return color.RGBA{0, 255 - constrast*1, 0, 0xff}
		case cmplx.Abs((0+1i)-z) < 1e-6:
			return color.RGBA{0, 0, 255 - constrast*i, 0xff}
		case cmplx.Abs((0-1i)-z) < 1e-6:
			return color.RGBA{255 - constrast*i, 255 - constrast*i, 0, 0xff}
		}
	}
	return color.Black
}
