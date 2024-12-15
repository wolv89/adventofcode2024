package day12

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
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

	datafile := DIR + "data.txt"
	if useSample == 1 {
		datafile = DIR + "sample.txt"
	} else if useSample == 2 {
		datafile = DIR + "sample2.txt"
	} else if useSample == 3 {
		datafile = DIR + "sample3.txt"
	} else if useSample == 4 {
		datafile = DIR + "sample4.txt"
	} else if useSample == 5 {
		datafile = DIR + "sample5.txt"
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
		i, j, total, sides int
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
		sides = Trace(rg, useSample > 0)
		if useSample == 0 {
			fmt.Println("#", string(rg.id), " | ", total, " | ", len(rg.points), " * ", sides)
		}
		total += len(rg.points) * sides
	}

	fmt.Println("")
	fmt.Println("Total: ", total)

}

func Trace(reg Region, dbg bool) int {

	var (
		pt                      Point
		maxh, maxw, s, t, sides int
		top, bot, lft, rgt      bool
	)

	if dbg {
		fmt.Println(string(reg.id), "---")
	}

	minw, minh := math.MaxInt, math.MaxInt

	for _, pt = range reg.points {
		minw = min(minw, pt.x)
		maxw = max(maxw, pt.x)
		minh = min(minh, pt.y)
		maxh = max(maxh, pt.y)
	}

	w, h := maxw-minw+1, maxh-minh+1

	shape := make([][]bool, h)
	for s = range shape {
		shape[s] = make([]bool, w)
	}

	for _, pt = range reg.points {
		shape[pt.y-minh][pt.x-minw] = true
	}

	// Scan left to right, checking top/bottom
	for s = range shape {
		top, bot = false, false
		for t = range shape[s] {
			if shape[s][t] {
				// Check top
				if s-1 < 0 || !shape[s-1][t] {
					if !top {
						top = true
						sides++
					}
				} else {
					top = false
				}
				// Check bottom
				if s+1 >= h || !shape[s+1][t] {
					if !bot {
						bot = true
						sides++
					}
				} else {
					bot = false
				}
			} else {
				top, bot = false, false
			}
		}
	}

	// Scan top to bottom, checking left/right
	for t = range shape[0] {
		lft, rgt = false, false
		for s = range shape {
			if shape[s][t] {
				// Check left
				if t-1 < 0 || !shape[s][t-1] {
					if !lft {
						lft = true
						sides++
					}
				} else {
					lft = false
				}
				// Check right
				if t+1 >= w || !shape[s][t+1] {
					if !rgt {
						rgt = true
						sides++
					}
				} else {
					rgt = false
				}
			} else {
				lft, rgt = false, false
			}
		}
	}

	if dbg {
		for s = range shape {
			var b strings.Builder
			for t = range shape[s] {
				if shape[s][t] {
					b.WriteByte(reg.id)
				} else {
					b.WriteByte('.')
				}
			}
			fmt.Println(b.String())
		}
		fmt.Println(len(reg.points), " * ", sides)
		fmt.Println("")
	}

	return sides

}
