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

func (d AocDay6) Puzzle1(useSample bool) {

	datafile := DIR + "data.txt"
	if useSample {
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

func (d AocDay6) Puzzle2(useSample bool) {

	datafile := DIR + "data.txt"
	if useSample {
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
		x, y, dir, h, w, turn, rev int
		found                      bool
	)

	grid := make([][]byte, 0)
	guard := Point{0, 0}
	spaces := 1
	obstacles := 0

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

	trail := make([][][4]bool, h)
	for t := range trail {
		trail[t] = make([][4]bool, w)
	}

	// Starting point is an UP
	trail[guard.x][guard.y][0] = true

	for rev = guard.x + 1; rev < h; rev++ {
		if grid[rev][guard.y] == '#' {
			break
		}
		trail[rev][guard.y][0] = true
	}

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

			// Every time we turn, we want to walk the trail
			// BACKWARDS until we hit bounds, or a #
			// There's probably a nicer way to do this rather than a
			// big ol' if/else block but it's very hot here and I'm
			// getting a bit sick of this puzzle.... :/

			// Turned UP so walk down
			if dir == 0 {
				for rev = guard.x + 1; rev < h; rev++ {
					if grid[rev][guard.y] == '#' {
						break
					}
					trail[rev][guard.y][0] = true
				}
				// Turned RIGHT so walk left
			} else if dir == 1 {
				for rev = guard.y - 1; rev >= 0; rev-- {
					if grid[guard.x][rev] == '#' {
						break
					}
					trail[guard.x][rev][1] = true
				}
				// Turned DOWN so walk up
			} else if dir == 2 {
				for rev = guard.x - 1; rev >= 0; rev-- {
					if grid[rev][guard.y] == '#' {
						break
					}
					trail[rev][guard.y][2] = true
				}
				// Turned LEFT so walk right
			} else if dir == 3 {
				for rev = guard.y + 1; rev < w; rev++ {
					if grid[guard.x][rev] == '#' {
						break
					}
					trail[guard.x][rev][3] = true
				}
			}

		} else {
			if grid[next.x][next.y] == '.' {
				spaces++
			}

			// Move
			guard.x, guard.y = next.x, next.y

			// Check
			turn = (dir + 1) % 4
			if trail[guard.x][guard.y][turn] {

				next.x += dirs[dir][0]
				next.y += dirs[dir][1]

				if next.x >= 0 && next.x < h && next.y >= 0 && next.y < w && grid[next.x][next.y] != '#' {
					obstacles++
					grid[next.x][next.y] = 'O'
				}

				next.x, next.y = guard.x, guard.y

			}

			// Mark
			trail[guard.x][guard.y][dir] = true

		}

	}

	fmt.Println("")

	for _, row := range grid {
		fmt.Println(string(row))
	}

	fmt.Println("")
	fmt.Println("Obstacles: ", obstacles)

}
