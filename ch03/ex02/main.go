package main

import (
	"fmt"
	"math"
)

const (
	width   = 600
	height  = 320
	cells   = 100
	xyrange = 100.0
	xyscale = width / 2 / xyrange
	zscale  = height * 0.4
	angle   = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func f(x, y float64) (value float64, err bool) {
	r := math.Hypot(x, y)
	result := math.Sin(r) / r
	if math.IsNaN(result) {
		return 0, true
	}
	return result, false
}

func corner(i, j int) (float64, float64, bool) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z, err := f(x, y)
	if err {
		return 0, 0, true
	}

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, false
}

func main() {
	fmt.Printf("<svg xmls='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, err := corner(i+1, j)
			if err {
				continue
			}
			bx, by, err := corner(i, j)
			if err {
				continue
			}
			cx, cy, err := corner(i, j+1)
			if err {
				continue
			}
			dx, dy, err := corner(i+1, j+1)
			if err {
				continue
			}
			fmt.Printf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' stroke='red' />\n", ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}
