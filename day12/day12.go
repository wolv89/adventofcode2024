package day12

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type AocDay12 struct{}

const DIR = "day12/"

type Point struct {
	x, y int
}

type Region struct {
	points []Point
	id     byte
	perim  int
}

func (d AocDay12) Puzzle1(useSample int) {

	datafile := DIR + "data.txt"
	if useSample == 1 {
		datafile = DIR + "sample.txt"
	} else if useSample == 2 {
		datafile = DIR + "sample2.txt"
	} else if useSample == 3 {
		datafile = DIR + "sample3.txt"
	}

	f, err := os.Open(datafile)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	grid := make([][]byte, 0)

	for scanner.Scan() {
		// @TODO: Investigate scanner.Bytes() vs []byte(scanner.Text())
		// Evidentally not the same...??!
		grid = append(grid, []byte(scanner.Text()))
	}

	h, w := len(grid), len(grid[0])
	seen := make([][]bool, h)
	for s := range seen {
		seen[s] = make([]bool, w)
	}

	if useSample > 0 {
		fmt.Println("")
		for _, row := range grid {
			fmt.Println(string(row))
		}
		fmt.Println("")
	}

	var (
		i, j, total int
	)

	regions := make([]Region, 0)

	for j = 0; j < h; j++ {
		for i = 0; i < w; i++ {

			if seen[j][i] {
				continue
			}

			seen[j][i] = true

			start := Point{i, j}
			region := Region{
				[]Point{start},
				grid[j][i],
				0,
			}

			Explore(grid, start, &region, &seen)
			regions = append(regions, region)

		}
	}

	for _, rg := range regions {
		fmt.Println("#", string(rg.id), " | ", total, " | ", len(rg.points), " * ", rg.perim)
		total += len(rg.points) * rg.perim
	}

	fmt.Println("")
	fmt.Println("Total: ", total)

}

func Explore(grid [][]byte, point Point, region *Region, seen *[][]bool) {

	// Up, Right, Down, Left
	var dirs = [4][2]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	var x, y int
	h, w := len(grid), len(grid[0]) // Recalc each time, or pass in func sig hmm

	for _, d := range dirs {

		x = point.x + d[0]
		y = point.y + d[1]

		if x < 0 || y < 0 || x >= w || y >= h || grid[y][x] != region.id {
			region.perim++
			continue
		}

		if (*seen)[y][x] {
			continue
		}

		(*seen)[y][x] = true
		next := Point{x, y}
		region.points = append(region.points, next)

		Explore(grid, next, region, seen)

	}

}

func (d AocDay12) Puzzle2(useSample int) {

}
