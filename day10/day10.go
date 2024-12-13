package day10

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
)

type AocDay10 struct{}

const DIR = "day10/"

// Left, Down, Right, Up
var DIRS = [4][2]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

// x - Horizontal, y - Vertical
type Point struct {
	x, y int
}

func (d AocDay10) Puzzle1(useSample int) {

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
		line           string
		h, w, x, count int
		c              byte
		wg             sync.WaitGroup
	)

	grid := make([][]byte, 0)
	trails := make([]Point, 0)

	for scanner.Scan() {

		line = scanner.Text()
		grid = append(grid, []byte(line))

		if w == 0 {
			w = len(line)
		}

		for x, c = range grid[h] {
			if c == '0' {
				trails = append(trails, Point{x, h})
			}
		}

		h++

	}

	for _, pt := range trails {

		// fmt.Println("# START ", pt, " ------")

		wg.Add(1)

		go func() {
			defer wg.Done()
			tc := StartWalking(grid, pt)
			count += tc
		}()

	}

	wg.Wait()

	fmt.Println("")
	fmt.Println("Trails: ", count)

}

func StartWalking(grid [][]byte, pt Point) int {

	vis := make([][]bool, len(grid))
	for v := range vis {
		vis[v] = make([]bool, len(grid[0]))
	}

	vis[pt.y][pt.x] = true
	count := 0

	WalkTrail(grid, pt, '0', &count, &vis)

	return count

}

func WalkTrail(grid [][]byte, pt Point, last byte, count *int, vis *[][]bool) {

	next := pt

	for _, d := range DIRS {

		next.x = pt.x + d[0]
		next.y = pt.y + d[1]

		if next.x < 0 || next.y < 0 || next.x >= len(grid[0]) || next.y >= len(grid) || (*vis)[next.y][next.x] {
			continue
		}

		// fmt.Println(next, " >> ", string(grid[next.y][next.x]))

		if grid[next.y][next.x] == last+1 {

			// fmt.Println(next, " | ", string(grid[next.y][next.x]))

			(*vis)[next.y][next.x] = true

			if grid[next.y][next.x] == '9' {
				*count++
			} else {
				WalkTrail(grid, next, last+1, count, vis)
			}

		}

	}

}

func (d AocDay10) Puzzle2(useSample int) {

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
		line           string
		h, w, x, count int
		c              byte
		wg             sync.WaitGroup
	)

	grid := make([][]byte, 0)
	trails := make([]Point, 0)

	for scanner.Scan() {

		line = scanner.Text()
		grid = append(grid, []byte(line))

		if w == 0 {
			w = len(line)
		}

		for x, c = range grid[h] {
			if c == '0' {
				trails = append(trails, Point{x, h})
			}
		}

		h++

	}

	// Near 600 starting points, wondering if this should be divied up a little differently
	// Ie not a new Goroutine for every start, but a new GR for every 10 or something?
	for _, pt := range trails {

		wg.Add(1)

		go func() {
			defer wg.Done()
			tc := StartWalking2(grid, pt)
			count += tc
		}()

	}

	wg.Wait()

	fmt.Println("")
	fmt.Println("Trails: ", count)

}

// Same as above, without the visited check... !
func StartWalking2(grid [][]byte, pt Point) int {

	count := 0

	WalkTrail2(grid, pt, '0', &count)

	return count

}

func WalkTrail2(grid [][]byte, pt Point, last byte, count *int) {

	next := pt

	for _, d := range DIRS {

		next.x = pt.x + d[0]
		next.y = pt.y + d[1]

		if next.x < 0 || next.y < 0 || next.x >= len(grid[0]) || next.y >= len(grid) {
			continue
		}

		if grid[next.y][next.x] == last+1 {

			if grid[next.y][next.x] == '9' {
				*count++
			} else {
				WalkTrail2(grid, next, last+1, count)
			}

		}

	}

}
