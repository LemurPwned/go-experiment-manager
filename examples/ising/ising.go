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
	// localMetripolis("sim", true)
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
		for _, nn := range neighbourMap[spinIdx] {
			subE += spinMap[nn]
		}
		E0 -= float64(spinMap[spinIdx] * subE)
	}
	return 0.5 * E0
}

func localMetripolis(folder string, saveImg bool) {
	L := 6
	N := L * L
	T := 2.0
	beta := 1.0 / T
	maxSteps := 100000

	// construct neighbour list
	neighbourMap := make(map[int][]int)
	spinMap := make([]int, N, N)
	for i := 0; i < N; i++ {
		neighbourMap[i] = []int{
			(i/L)*L + (i+1)%L,
			(i + L) % N,
			(i/L)*L + ((i-1)%L+L)%L,
			((i-L)%N + N) % N,
		}
		// fill the spin
		spinMap[i] = 2*rand.Intn(2) - 1 // 1 or -1
	}
	Etot := calcEnergy(spinMap, neighbourMap)
	energies := make([]float64, 0, maxSteps)
	for step := 0; step < maxSteps; step++ {
		k := rand.Intn(N)
		deltaE := 0.0
		for _, spinNIdx := range neighbourMap[k] {
			deltaE += float64(spinMap[spinNIdx])
		}
		deltaE *= 2.0 * float64(spinMap[k])
		if rand.Float64() < math.Exp(-beta*deltaE) {
			spinMap[k] *= -1
			Etot += deltaE
		}
		energies = append(energies, Etot)
	}
	energySum := 0.0
	for _, val := range energies {
		energySum += val
	}
	fmt.Printf("Mean energy %f per spin", energySum/float64(maxSteps*N))

}

func clusterWolff(folder string, saveImg bool) {
	maxSteps := 10000
	L := 100
	N := L * L
	T := 2.0
	p := 1.0 - math.Exp(-2.0/T)

	// construct neighbour list
	neighbourMap := make(map[int][]int)
	spinMap := make([]int, N, N)
	for i := 0; i < N; i++ {
		neighbourMap[i] = []int{
			(i/L)*L + (i+1)%L,
			(i + L) % N,
			(i/L)*L + ((i-1)%L+L)%L,
			((i-L)%N + N) % N,
		}
		// fill the spin
		spinMap[i] = 2*rand.Intn(2) - 1 // 1 or -1
	}

	E := calcEnergy(spinMap, neighbourMap)
	energies := make([]float64, 0, maxSteps+1)
	energies = append(energies, E)
	fmt.Printf("Step: %d, Energy: %.3f\n", 0, E)
	for step := 1; step < maxSteps; step++ {
		// take a random spin
		k := rand.Intn(N)
		pocket := make([]int, 0, N)
		cluster := make(map[int]int)

		pocket = append(pocket, k)
		cluster[k] = k
		// constructing the large cluster
		for {

			if len(pocket) == 0 {
				break
			}
			j := pocket[rand.Intn(len(pocket))]
			// log.Println(pocket, j, cluster, neighbourMap[j])
			for _, l := range neighbourMap[j] {
				_, ok := cluster[l]
				if !ok && (spinMap[l] == spinMap[j]) && (rand.Float64() < p) {
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
		energies = append(energies, E)
		if step%1 == 0 {
			// fmt.Printf("Step: %d, Energy: %.3f\n", step, E)
		}
	}

	meanEnergy := 0.0
	meanEnergy2 := 0.0
	for _, eng := range energies {
		meanEnergy += eng
		meanEnergy2 += (eng * eng)
	}
	meanEnergy = meanEnergy / float64(maxSteps)
	meanEnergy2 = meanEnergy2 / float64(maxSteps)
	specificHeat := (meanEnergy2 - math.Pow(meanEnergy, 2)) / float64(N) / (math.Pow(T, 2))
	fmt.Printf("Mean energy: %.4f\nSpecific Heat: %.4f\nEnergy per spin %.4f\n", meanEnergy, specificHeat, meanEnergy/float64(N))
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
