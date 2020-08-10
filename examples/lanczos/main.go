
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

// FilterImage filters the image with a specified kernel
func FilterImage(img image.Image, kernel [][]float64) image.Image{
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()
	ksize := len(kernel)
	pad := ((ksize + 1)/2) - 1
	
	newImg := image.NewGray(image.Rect(0, 0, width+1, height+1))
	var wg = sync.WaitGroup{}
	for x := pad; x < (width - pad); x ++   {
		wg.Add(1)
		go func(x int, wg *sync.WaitGroup) {
			defer wg.Done()
			for y := pad; y < (height - pad); y ++ {
				sum := 0.0
				for i := -pad; i <= pad; i++ {
					for j :=pad; j <= pad; j++{

						grayVal := color.GrayModel.Convert(img.At(x+i, y+j)).(color.Gray).Y
						sum += float64(grayVal)*kernel[i+pad][j+pad]
					}
				}
				color := color.Gray{Y: uint8(math.Max(math.Min(sum, 255), 0))}
				newImg.SetGray(x, y, color)
			}
		}(x, &wg)
	}
	wg.Wait()
	return newImg
}

// Lanczos function
func Lanczos(x float64) float64 {
	if math.Abs(x)  >= 2 {
		return 0.0 
	}
	px := math.Pi*x
	a := math.Sin(px)/(px)
	b := math.Sin(px/2)/(px/2)
	return a*b

}

func lanczosInterpolation(srcImg image.Image) image.Image{
	width := srcImg.Bounds().Dx()
	height := srcImg.Bounds().Dy()
	upScaledImg := image.NewGray(image.Rect(0, 0, width+1, height+1))
	for x := 1; x < srcImg.Bounds().Dx()-3; x ++ {
		for y := 1; y < srcImg.Bounds().Dy()-3; y++ {

			n := 2
			w := 0.0
			v := 0.0
			for i := -n+1; i <= n; i ++ {
				for j := -n+1; j <= n; j++{
					lx := Lanczos(float64(i - x) + math.Ceil(float64(x)))
					ly := Lanczos(float64(i - y) + math.Ceil(float64(y)))
					w += lx*ly
					v += float64(srcImg.At(int(math.Ceil(float64(x))) + i, 
								   int(math.Ceil(float64(y))) + j).(color.Gray).Y)*lx*ly
				}
			}
			sum := (1/w)*v
			// X := [4]int{x -1 ,x + 1,x+2,x+3}
			// Y := [4]int{y -1, y+1, y+2, y+3}
			// I := make([]float64, 4)
			// sum := 0.0
			// for k := 0; k <= 3; k++ {
			// 	for i:= 0; i<= 3; i++ {
			// 		I[k] += Lanczos(float64(x - X[i])) * float64(srcImg.At(X[i], Y[k]).(color.Gray).Y)
			// 	}
			// 	sum += I[k] * Lanczos(float64(y - Y[k]))
			// }
			// log.Println(sum)
			color := color.Gray{Y: uint8(math.Max(math.Min(sum, 255), 0))}
			upScaledImg.SetGray(x, y, color)
		}
	}
	return upScaledImg
}



func readAndFilter(fn string, kernelType string){

	f, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}
	srcImg, _, err := image.Decode(f)
	var img image.Image
	start := time.Now()
	if kernelType == "sobel"{
		kernel := [][]float64 {
			{-1, 0, 1 },
			{-1, 0 , 1},
			{-1, 0 , 1},
		}
		img = FilterImage(srcImg, kernel)
	} else if (kernelType == "lanczos"){
		grayImg := image.NewGray(image.Rect(0, 0, 
				srcImg.Bounds().Dx(), srcImg.Bounds().Dy()))
		for x := 0; x < srcImg.Bounds().Dx(); x++{
			for y := 0; y < srcImg.Bounds().Dy(); y++{
				grayImg.SetGray(x, y, color.GrayModel.Convert(srcImg.At(x, y)).(color.Gray)) 
			}
		}
		
		img = lanczosInterpolation(grayImg)
		// img = grayImg
	}

	f, _ = os.Create(kernelType + "_fiterResult.png")
	png.Encode(f, img)

	elapsed := time.Since(start)
	log.Printf("Generated an image %dx%d in %s\n", img.Bounds().Dx(), 
													img.Bounds().Dy(), elapsed)
}




func main(){
	readAndFilter("misty.png", "lanczos")
}