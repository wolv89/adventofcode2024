package day06

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type AocDay6 struct{}

const DIR = "day06/"

// x - Vertical, y - Horizontal :/
type Point struct {
	x, y int
}

func (d AocDay6) Puzzle1(useSample int) {

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

	var (
		line            string
		x, y, dir, h, w int
		found           bool
	)

	grid := make([][]byte, 0)
	guard := Point{0, 0}
	spaces := 1

	// Up, Right, Down, Left
	var dirs = [4][2]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	for scanner.Scan() {

		line = scanner.Text()

		grid = append(grid, []byte(line))

		if !found {
			for y = 0; y < len(line); y++ {
				if line[y] == '^' {
					guard.x = x
					guard.y = y
					found = true
					break
				}
			}
		}

		x++

	}

	h, w = len(grid), len(grid[0])
	next := Point{guard.x, guard.y}

patrol:
	for {

		next.x += dirs[dir][0]
		next.y += dirs[dir][1]

		if next.x < 0 || next.x >= h || next.y < 0 || next.y >= w {
			break patrol
		}

		if grid[next.x][next.y] == '#' {
			next.x, next.y = guard.x, guard.y
			dir = (dir + 1) % 4
		} else {
			if grid[next.x][next.y] == '.' {
				spaces++
				grid[next.x][next.y] = 'X'
			}
			guard.x, guard.y = next.x, next.y
		}

	}

	fmt.Println("")
	fmt.Println("Spaces: ", spaces)

}

func (d AocDay6) Puzzle2(useSample int) {

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

	var (
		line                       string
		x, y, dir, h, w, obstacles int
		found                      bool
	)

	grid := make([][]byte, 0)
	guard := Point{0, 0}

	// Up, Right, Down, Left
	var dirs = [4][2]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	for scanner.Scan() {

		line = scanner.Text()

		grid = append(grid, []byte(line))

		if !found {
			for y = 0; y < len(line); y++ {
				if line[y] == '^' {
					guard.x = x
					guard.y = y
					found = true
					break
				}
			}
		}

		x++

	}

	h, w = len(grid), len(grid[0])
	next, start := Point{guard.x, guard.y}, Point{guard.x, guard.y}

	checked := make([][]bool, h)
	for ch := range checked {
		checked[ch] = make([]bool, w)
	}

	for {

		next.x += dirs[dir][0]
		next.y += dirs[dir][1]

		if next.x < 0 || next.x >= h || next.y < 0 || next.y >= w {
			break
		}

		if grid[next.x][next.y] == '#' {

			next.x, next.y = guard.x, guard.y
			dir = (dir + 1) % 4

		} else {

			guard.x, guard.y = next.x, next.y

			if !checked[guard.x][guard.y] && (guard.x != start.x || guard.y != start.y) {

				grid[guard.x][guard.y] = '#'

				if PatrolLoops(grid, start) {
					obstacles++
				}

				grid[guard.x][guard.y] = '.'
				checked[guard.x][guard.y] = true

			}

		}

	}

	fmt.Println("")
	fmt.Println("Obstacles: ", obstacles)

}

func PatrolLoops(grid [][]byte, guard Point) bool {

	h, w := len(grid), len(grid[0])
	next := Point{guard.x, guard.y}

	hist := make([][][4]bool, h)
	for hi := range hist {
		hist[hi] = make([][4]bool, w)
	}

	// Up, Right, Down, Left
	var dirs = [4][2]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}
	dir := 0

	for {

		next.x += dirs[dir][0]
		next.y += dirs[dir][1]

		// Leaving bounds, does not loop
		if next.x < 0 || next.x >= h || next.y < 0 || next.y >= w {
			return false
		} else if hist[next.x][next.y][dir] {
			return true
		}

		if grid[next.x][next.y] == '#' {
			next.x, next.y = guard.x, guard.y
			dir = (dir + 1) % 4
			hist[guard.x][guard.y][dir] = true
		} else {
			guard.x, guard.y = next.x, next.y
			hist[guard.x][guard.y][dir] = true
		}

	}

}
