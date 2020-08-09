package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
	"sync"
	"time"
)

var colourArray = []color.NRGBA{
	{66, 30, 15, 0xff},
	{25, 7, 26, 0xff},
	{9, 1, 47, 0xff},
	{4, 4, 73, 0xff},
	{0, 7, 100, 0xff},
	{12, 44, 138, 0xff},
	{24, 82, 177, 0xff},
	{57, 125, 209, 0xff},
	{134, 181, 229, 0xff},
	{211, 236, 248, 0xff},
	{241, 233, 191, 0xff},
	{248, 201, 95, 0xff},
	{255, 170, 0, 0xff},
	{204, 128, 0, 0xff},
	{153, 87, 0, 0xff},
	{106, 52, 3, 0xff},
}

type point struct {
	x float64
	y float64
}

type cell struct {
	x        int
	y        int
	drawable uint32
}

func (p *point) abs() float64 {

	return math.Sqrt(p.x*p.x + p.y*p.y)
}

func generateFractalImg(fn string) {

	start := time.Now()
	imgX := 640
	imgY := 640
	m := image.NewRGBA(image.Rect(0, 0, imgX, imgY))

	xStart := -1.0
	yStart := -1.0
	xStop := 1.0
	yStop := 1.0
	maxIter := 10000
	var wg sync.WaitGroup
	for x := 0; x < imgX; x++ {
		wg.Add(1)
		go func(x int, wg *sync.WaitGroup) {
			defer wg.Done()
			for y := 0; y < imgY; y++ {
				p := point{
					x: float64(xStart + (float64(x)/float64(imgX))*(xStop-xStart)),
					y: float64(yStart + (float64(y)/float64(imgY))*(yStop-yStart))}

				px := mandelbrotPixel(maxIter, p)
				cx := int(px) % len(colourArray)
				c := colourArray[cx]
				m.Set(x, y, c)
			}
		}(x, &wg)
	}
	wg.Wait()

	f, _ := os.Create("image.png")
	png.Encode(f, m)

	elapsed := time.Since(start)
	log.Printf("Generated an image %dx%d in %s\n", imgX, imgY, elapsed)
}

func mandelbrotPixel(maxIter int, c point) float64 {
	z := point{0, 0}
	escapeRadius := 20.0
	n := 0
	var modulus float64
	for {
		modulus = z.abs()
		if modulus > escapeRadius {
			break
		}
		if n == maxIter {
			return float64(maxIter)
		}
		z = point{
			x: z.x*z.x - z.y*z.y + c.x,
			y: 2*(z.x*z.y) + c.y,
		}
		n++
	}
	return float64(n) - (math.Log10(math.Log10(modulus)) / math.Log10(2))
}
