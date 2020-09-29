package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/rand"
	"os"
	"strconv"
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
		// log.Printf("Input %s => %s\n", inp, newGeneration[(left+1+genLen)%genLen])
		if (left + modifier) == genLen-2 {
			break
		}
		left++
		right++
	}
	return newGeneration
}

func (driver *AutomataDriver) runNGenerations(generation []int, N int, boundary bool) [][]string {

	// log.Println("Converting to string input...")
	genStr := make([]string, len(generation))

	for i := 0; i < len(generation); i++ {
		genStr[i] = fmt.Sprintf("%d", generation[i])
	}
	// log.Println(genStr)

	// log.Printf("Running for %d generations \n", N)
	generationList := make([][]string, N+1)
	generationList[0] = genStr
	for n := 0; n < N; n++ {
		generationList[n+1] = make([]string, len(genStr))
		generationList[n+1] = driver.transformGeneration(generationList[n], boundary)
		// log.Printf("Gen: %d: %v\n", n, generationList[n+1])
	}
	// hopefully save the image
	return generationList
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

func randomRuns(rowNumber, initNums int) [][]int {

	inits := make([][]int, initNums)

	for i := 0; i < initNums; i++ {
		inits[i] = make([]int, rowNumber)
		for j := 0; j < rowNumber; j++ {
			n := 0
			if rand.Float64() > 0.5 {
				n = 1
			}
			inits[i][j] = n
		}
	}
	return inits
}

func simulateIntensity() {

	runs := 1000
	rows := 100
	generations := 60
	inits := randomRuns(rows, runs)

	rule := 42
	a := AutomataDriver{
		r:        generateRule(rule),
		ruleSize: 3,
		rowSize:  rows,
	}

	outputPlaceholder := make([][]int, generations+1)
	for i := 0; i < generations+1; i++ {
		outputPlaceholder[i] = make([]int, rows)
		for j := 0; j < rows; j++ {
			outputPlaceholder[i][j] = 0
		}
	}
	for i := 0; i < runs; i++ {
		// a.runNGenerations(inits[i], generations, true)
		res := a.runNGenerations(inits[i], generations, true)
		for g := range res {
			for r := range res[g] {
				num, _ := strconv.Atoi(res[g][r])
				outputPlaceholder[g][r] += num
			}
		}
	}
	name := fmt.Sprintf("Rule_%d_intensity", rule)
	drawInitialisations(outputPlaceholder, rows, generations, name)
}

func drawInitialisations(initGroups [][]int, rowSize, N int, name string) {
	img := image.NewGray(image.Rect(0, 0, rowSize, N))
	f, _ := os.Create(name + ".png")

	// older gens grow UP
	// row is read left-right
	// find max
	max := 0

	for i := range initGroups {
		for j := range initGroups[i] {
			if initGroups[i][j] > max {
				max = initGroups[i][j]
			}
		}
	}

	for i := range initGroups {
		for j := range initGroups[i] {
			if initGroups[i][j] == 0 {
				img.Set(j, i, color.Gray{0})
				continue
			}
			grayNormalised := uint8(255 * max / initGroups[i][j])
			img.Set(j, i, color.Gray{grayNormalised})
		}
	}

	png.Encode(f, img)
}

func main() {
	simulateIntensity()
	// a := AutomataDriver{
	// 	r:        generateRule(30),
	// 	ruleSize: 3,
	// 	rowSize:  100,
	// }
	// firstRow := make([]int, a.rowSize)
	// // bias := 0.3
	// // for i := range firstRow {
	// // 	if rand.Float64() > bias {
	// // 		firstRow[i] = 1
	// // 	}
	// // }
	// firstRow[a.rowSize/2] = 1
	// N := 30
	// gList := a.runNGenerations(firstRow, N, true)
	// makeImage(gList, N+1, len(gList), a.r.Name)

	// a.runNGenerations([]int{
	// 	0, 1, 1, 0, 1, 1, 1, 0, 0, 1, 1, 0, 0, 0, 1, 0,
	// }, 30, true)
}
