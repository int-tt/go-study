package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"net/http"
)

var (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = float64(width) / 2 / xyrange
	zscale        = float64(height) * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

// TODO: 高さ、幅、色対応
func main() {
	http.HandleFunc("/", renderSVG)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
func renderSVG(w http.ResponseWriter, r *http.Request) {
	str := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; storke-width: 0.7' "+
		"width='%d' height='%d' >", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, err := corner(i+1, j)
			if err != nil {
				fmt.Println(err)
				continue
			}
			bx, by, err := corner(i, j)
			if err != nil {
				fmt.Println(err)
				continue
			}
			cx, cy, err := corner(i, j+1)
			if err != nil {
				fmt.Println(err)
				continue
			}
			dx, dy, err := corner(i+1, j+1)
			if err != nil {
				fmt.Println(err)
				continue
			}

			str = str + fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g' />\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	str = str + "</svg>"
	w.Header().Set("Content-Type", "image/svg+xml")
	w.Write([]byte(str))
}

func corner(i, j int) (float64, float64, error) {
	x := xyrange * (float64(i)/float64(cells) - 0.5)
	y := xyrange * (float64(j)/float64(cells) - 0.5)
	z := f(x, y)
	if math.IsNaN(z) {
		return 0, 0, errors.New("Outside the float")

	}
	sx := float64(width)/2 + (x-y)*cos30*xyscale
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, nil
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
