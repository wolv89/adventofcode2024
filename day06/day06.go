package day06

import (
	"bufio"
	"container/ring"
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
		line                           string
		x, y, dir, h, w, turns, px, py int
		found                          bool
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

	trail := make([][]byte, h)
	for t := range trail {
		trail[t] = make([]byte, w)
	}

	fmt.Println(guard)
	fmt.Println("")

	const cn = 3
	corners := ring.New(cn)

	poss := Point{-1, -1}

	tr := Point{guard.x, guard.y}
	for {
		tr.x++
		if tr.x >= h {
			break
		}
		trail[tr.x][tr.y] = '^'
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

			corners.Value = Point{guard.x, guard.y}
			corners = corners.Next()

			if turns < 2 {
				turns++
			} else {

				px, py = 0, 0

				corners.Do(func(cnr any) {
					px ^= cnr.(Point).x
					py ^= cnr.(Point).y
				})

				poss.x, poss.y = px, py

				fmt.Println("> ", poss)
			}

		} else {
			if grid[next.x][next.y] == '.' {
				spaces++
			}
			guard.x, guard.y = next.x, next.y

			if guard.x == poss.x && guard.y == poss.y {

				next.x = guard.x + dirs[dir][0]
				next.y = guard.y + dirs[dir][1]

				if next.x >= 0 || next.x < h || next.y >= 0 || next.y < w && grid[next.x][next.y] == '.' {
					grid[next.x][next.y] = 'O'
					obstacles++
				}

				next.x, next.y = guard.x, guard.y

			}

		}

	}

	fmt.Println("")

	for _, row := range grid {
		for _, cell := range row {
			fmt.Printf("%s ", string(cell))
		}
		fmt.Println("")
	}

	fmt.Println("----------")

	for _, row := range trail {
		for _, cell := range row {
			if cell != 0 {
				fmt.Printf("%s ", string(cell))
			} else {
				fmt.Printf(". ")
			}
		}
		fmt.Println("")
	}

	fmt.Println("")
	fmt.Println("Obstacles: ", obstacles)

}
