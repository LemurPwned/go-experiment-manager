package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

// Rule3 determines how to transform generations
type Rule3 struct {
	Name string
	// 3 elements => 1 output element
	RuleMap map[string]string
}

func (r *Rule3) nextGeneration(cellGroup string) string {
	// parse rule
	return r.RuleMap[cellGroup]
}

func generateRule(ruleNumber int) Rule3 {
	ruleMap := make(map[string]string)
	ruleOutput := fmt.Sprintf("%08b", ruleNumber)
	log.Println(ruleOutput)
	for i := 0; i < 8; i++ {
		ruleMap[fmt.Sprintf("%03b", i)] = string(ruleOutput[len(ruleOutput)-1-i])
	}
	log.Println(ruleMap)
	return Rule3{
		Name:    fmt.Sprintf("Rule %d", ruleNumber),
		RuleMap: ruleMap,
	}
}

// AutomataDriver drives the automata
type AutomataDriver struct {
	r        Rule3
	rowSize  int
	ruleSize int
}

// path intensity
// generate N images with N different initialisations
// every path will be a option -- dynamic programming
func (driver *AutomataDriver) transformGeneration(generation []string, boundary bool) []string {
	// we wrap on boundary conditions
	left := -1
	right := 1 // we start with 1 right neighbour
	// rule is based off the length
	modifier := 0
	if !boundary {
		left = 0
		right = 2
		modifier = 1
	}

	genLen := len(generation)
	newGeneration := make([]string, driver.rowSize)
	for {
		// construct neighbours
		indices := []int{
			(left + genLen) % genLen,
			(left + 1 + genLen) % genLen,
			(right + genLen) % genLen,
		}
		inp := fmt.Sprintf("%s%s%s", generation[indices[0]], generation[indices[1]], generation[indices[2]])
		newGeneration[(left+1+genLen)%genLen] = driver.r.nextGeneration(inp)
		log.Printf("Input %s => %s\n", inp, newGeneration[(left+1+genLen)%genLen])
		if (left + modifier) == genLen-2 {
			break
		}
		left++
		right++
	}
	return newGeneration
}

func (driver *AutomataDriver) runNGenerations(generation []int, N int, boundary bool) {

	log.Println("Converting to string input...")
	genStr := make([]string, len(generation))

	for i := 0; i < len(generation); i++ {
		genStr[i] = fmt.Sprintf("%d", generation[i])
	}
	log.Println(genStr)

	log.Printf("Running for %d generations \n", N)
	generationList := make([][]string, N+1)
	generationList[0] = genStr
	for n := 0; n < N; n++ {
		generationList[n+1] = make([]string, len(genStr))
		generationList[n+1] = driver.transformGeneration(generationList[n], boundary)
		log.Printf("Gen: %d: %v\n", n, generationList[n+1])
	}

	log.Println("Saving the image")
	// hopefully save the image
	makeImage(generationList, N+1, len(generation), driver.r.Name)
}

const (
	pos = "1"
	neg = "0"
)

func makeImage(ruleMap [][]string, N, rowSize int, name string) {
	img := image.NewRGBA(image.Rect(0, 0, rowSize, N))
	f, _ := os.Create(name + ".png")

	// older gens grow UP
	// row is read left-right
	for i := range ruleMap {
		for j := range ruleMap[i] {

			switch ruleMap[i][j] {
			case "0":
				img.Set(j, i, color.RGBA{255, 255, 255, 0xff})
			case "1":
				img.Set(j, i, color.RGBA{0, 0, 0, 0xff})
			}
		}
	}

	png.Encode(f, img)
}
func main() {

	a := AutomataDriver{
		r:        generateRule(30),
		ruleSize: 3,
		rowSize:  100,
	}
	firstRow := make([]int, a.rowSize)
	// bias := 0.3
	// for i := range firstRow {
	// 	if rand.Float64() > bias {
	// 		firstRow[i] = 1
	// 	}
	// }
	firstRow[a.rowSize/2] = 1
	a.runNGenerations(firstRow, 30, true)
	// a.runNGenerations([]int{
	// 	0, 1, 1, 0, 1, 1, 1, 0, 0, 1, 1, 0, 0, 0, 1, 0,
	// }, 30, true)
}
