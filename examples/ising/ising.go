package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/rand"
	"os"
)

func main() {
	clusterWolff("sim", true)
}

func makeImage(spinMap []int, L int, name string) {
	img := image.NewRGBA(image.Rect(0, 0, L, L))
	f, _ := os.Create(name + ".png")

	for i := 0; i < L; i++ {
		for j := 0; j < L; j++ {

			switch spinMap[i*L+j] {
			case spinUp:
				img.Set(i, j, spinUpColour)
			case spinDown:
				img.Set(i, j, spinDownColour)
			default:
				img.Set(i, j, color.RGBA{0, 255, 0, 0xff})
			}
		}
	}

	// col := color.RGBA{0, 0, 0, 255}
	// point := fixed.Point26_6{fixed.Int26_6(300),
	// 	fixed.Int26_6(3000)}

	// face := inconsolata.Bold8x16
	// face.Height = 32
	// d := &font.Drawer{
	// 	Dst:  img,
	// 	Src:  image.NewUniform(col),
	// 	Face: face,
	// 	Dot:  point,
	// }
	// d.DrawString(fmt.Sprintf("Energy %.2f", 0.02))

	png.Encode(f, img)
}

// Two spins
// up - up / down - down => energy -1
// down - up => energy +1

var (
	spinUp   = 1
	spinDown = -1

	spinUpColour   = color.RGBA{255, 0, 0, 0xff}
	spinDownColour = color.RGBA{0, 0, 255, 0xff}
)

func calcEnergy(spinMap []int, neighbourMap map[int][]int) float64 {
	E0 := 0.0
	for spinIdx := range spinMap {
		subE := 0
		for spinNIdx := range neighbourMap[spinIdx] {
			subE += spinMap[spinNIdx]
		}
		E0 += float64(spinMap[spinIdx] * subE)
	}
	return -0.5 * E0
}

func clusterWolff(folder string, saveImg bool) {
	maxSteps := 350
	L := 100
	N := L * L
	T := 2.5
	p := 1.0 - math.Exp(-2.0/T)

	// construct neighbour list
	neighbourMap := make(map[int][]int)
	spinMap := make([]int, N, N)
	for i := 0; i < N; i++ {
		neighbourMap[i] = []int{
			(i/L)*L + (i+1)%L,
			(i + L) % N,
			(i - L) % N,
			(i/L)*L + (i-1)%L,
		}
		// fill the spin
		spinMap[i] = 2*rand.Intn(2) - 1 // 1 or -1

	}

	pocket := make([]int, 0)
	E := calcEnergy(spinMap, neighbourMap)
	fmt.Printf("Step: %d, Energy: %.3f\n", 0, E)
	for step := 1; step < maxSteps; step++ {
		// take a random spin
		k := rand.Intn(N)
		pocket = pocket[:0]
		cluster := make(map[int]int)

		pocket = append(pocket, k)
		cluster[k] = k
		// constructing the large cluster
		for {
			if len(pocket) == 0 {
				break
			}
			j := pocket[rand.Intn(len(pocket))]
			for l := range neighbourMap[j] {
				_, ok := cluster[l]
				if ok && (spinMap[l] == spinMap[j]) && (rand.Float64() < p) {
					pocket = append(pocket, l)
					cluster[l] = l
				}
			}
			pocket = removeVal(pocket, j)
		}
		for spinIdx := range cluster {
			spinMap[spinIdx] *= -1 // flip the spins
		}
		if saveImg {
			// save the image if desired
			makeImage(spinMap, L, fmt.Sprintf("%s/Iteration_%d", folder, step))
		}
		E := calcEnergy(spinMap, neighbourMap)
		fmt.Printf("Step: %d, Energy: %.3f\n", step, E)
	}

}

func removeVal(s []int, val int) []int {

	for i, v := range s {
		if v == val {
			return remove(s, i)
		}
	}
	return s
}

func remove(s []int, i int) []int {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}
