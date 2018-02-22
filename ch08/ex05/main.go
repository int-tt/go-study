package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"sync"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 2048, 2048
	)
	wg := &sync.WaitGroup{}
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		wg.Add(1)
		//img.Setの順番がばらばらになることはある...?
		go func(py int) {
			y := float64(py)/height*(ymax-ymin) + ymin
			for px := 0; px < width; px++ {
				x := float64(px)/width*(xmax-xmin) + xmin
				z := complex(x, y)
				// 画像の点(px,py)は複素数値zを表している。
				img.Set(px, py, mandelbrot(z))
			}
			wg.Done()
		}(py)
	}
	wg.Wait()
	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			//return color.Gray{255 - contrast*n}
			return color.RGBA{n, contrast * n, 255 - contrast*n, 0xff}
		}
	}
	return color.Black
}
