package day04

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type AocDay4 struct{}

const DIR = "day04/"

func (d AocDay4) Puzzle1(useSample int) {

	datafile := DIR + "data.txt"
	if useSample == 1 {
		datafile = DIR + "sample.txt"
	}

	f, err := os.Open(datafile)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	grid := make([]string, 0)

	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	// Technically we know it's a grid, same width by height, but eh...
	h, w := len(grid), len(grid[0])
	var i, j, a, b, c, s, total int
	var valid bool

	// Off by 1 in the first run on the sample
	// Leaving in the debug code, but commented....
	// cells := make([][]bool, h)
	// for cl := range cells {
	// 	cells[cl] = make([]bool, w)
	// }

	checks := [8][2]int{
		{0, 1},   // Right
		{1, 1},   // Down/Right
		{1, 0},   // Down
		{1, -1},  // Down/Left
		{0, -1},  // Left
		{-1, -1}, // Up/Left
		{-1, 0},  // Up
		{-1, 1},  // Up/Right
	}

	// checkNames := [8]string{"Right", "Down/Right", "Down", "Down/Left", "Left", "Up/Left", "Up", "Up/Right"}

	// Should constants be ALL CAPS ??
	const (
		lead  = 'X'
		trail = "MAS"
		trlen = 3
	)

	for i = 0; i < h; i++ {
		for j = 0; j < w; j++ {

			if grid[i][j] != lead {
				continue
			}

			// cells[i][j] = true

			// fmt.Println("")
			// fmt.Println("# ", i, j)
			// fmt.Println("---")
			// fmt.Println("")

			for c = range checks {

				// fmt.Println("Checking ", checkNames[c])

				// Pre-flight bounds check
				a = i + (checks[c][0] * trlen)
				b = j + (checks[c][1] * trlen)

				if a < 0 || a >= h || b < 0 || b >= w {
					// fmt.Println("Failed bounds check ", a, b)
					continue
				}

				a, b = i, j
				valid = true

				for s = 0; s < trlen; s++ {

					a += checks[c][0]
					b += checks[c][1]

					if grid[a][b] != trail[s] {
						valid = false
						break
					}
					/* else {
						cells[a][b] = true
					} */

				}

				if valid {
					// fmt.Println("MATCHED")
					total++
				}
				/* else {
					fmt.Println("Missed")
				} */

			}

		}
	}

	/*
		for i = 0; i < h; i++ {
			line := make([]byte, 0, w)
			for j = 0; j < w; j++ {
				if cells[i][j] {
					line = append(line, grid[i][j])
				} else {
					line = append(line, '.')
				}
			}
			fmt.Println(string(line))
		}
	*/

	fmt.Println("Total: ", total)
	fmt.Println("")

}

func (d AocDay4) Puzzle2(useSample int) {

	datafile := DIR + "data.txt"
	if useSample == 1 {
		datafile = DIR + "sample.txt"
	}

	f, err := os.Open(datafile)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	grid := make([]string, 0)

	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	h, w := len(grid), len(grid[0])
	var i, j, total, up, dw, lf, rg int

	for i = 1; i < h-1; i++ {
		for j = 1; j < w-1; j++ {

			if grid[i][j] != 'A' {
				continue
			}

			// Up (x2), Down (x2), Left, Right (x2 ABAB Select Start...)
			up, dw = i-1, i+1
			lf, rg = j-1, j+1

			// fmt.Println(string(grid[up][lf]), ".", string(grid[up][rg]))
			// fmt.Println(".", string(grid[i][j]), ".")
			// fmt.Println(string(grid[dw][lf]), ".", string(grid[dw][rg]))
			// fmt.Println("")

			if (grid[up][lf] == 'S' && grid[dw][rg] == 'M') || (grid[up][lf] == 'M' && grid[dw][rg] == 'S') {
				if (grid[dw][lf] == 'S' && grid[up][rg] == 'M') || (grid[dw][lf] == 'M' && grid[up][rg] == 'S') {
					total++
				}
			}

		}
	}

	fmt.Println("Total: ", total)
	fmt.Println("")

}
