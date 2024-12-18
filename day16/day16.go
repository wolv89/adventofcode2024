package day16

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type AocDay16 struct{}

const DIR = "day16/"

type Point struct {
	x, y int
}

type Vec struct {
	x, y, dir int
}

type Olympics struct {
	path [][]int
	maze [][]byte
	goal Point
	w, h int
}

// East, South, West, North
var DIRS = [4]Point{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

func (d AocDay16) Puzzle1(useSample int) {

	datafile := DIR + "data.txt"
	if useSample == 1 {
		datafile = DIR + "sample.txt"
	} else if useSample == 2 {
		datafile = DIR + "sample2.txt"
	}

	f, err := os.Open(datafile)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var (
		start Vec
		end   Point
		line  string
		x, y  int
	)

	maze := make([][]byte, 0)

	for scanner.Scan() {

		line = scanner.Text()

		// Trimming # from sides
		maze = append(maze, []byte(line)[1:len(line)-1])

		// Offset because not including outer border in maze/grid
		for x = 1; x < len(line)-1; x++ {
			if line[x] == 'S' {
				start = Vec{x - 1, y - 1, 0}
			} else if line[x] == 'E' {
				end = Point{x - 1, y - 1}
			}
		}

		y++

	}

	// Trim top/bottom
	maze = maze[1 : len(maze)-1]

	w, h := len(maze[0]), len(maze)

	path := make([][]int, h)
	for y = range path {
		path[y] = make([]int, w)
	}

	fmt.Println(strings.Repeat("#", w+2))
	for _, row := range maze {
		fmt.Println("#" + string(row) + "#")
	}
	fmt.Println(strings.Repeat("#", w+2))

	oly := Olympics{
		path,
		maze,
		end,
		w,
		h,
	}

	oly.Walk(start, 1)

	/*
		fmt.Println("")
		for _, pt := range oly.path {
			for _, tp := range pt {
				fmt.Printf("%7d ", tp)
			}
			fmt.Print("\n\n")
		}

		fmt.Println("")
	*/

	fmt.Println("")
	fmt.Println("Score: ", oly.path[end.y][end.x]-1)

}

func (o *Olympics) Walk(spot Vec, score int) {

	if o.path[spot.y][spot.x] != 0 {
		if o.path[spot.y][spot.x] > score {
			o.path[spot.y][spot.x] = score
		} else {
			return
		}
	}

	o.path[spot.y][spot.x] = score

	// Made it...
	if spot.x == o.goal.x && spot.y == o.goal.y {
		return
	}

	for d, dir := range DIRS {

		next := Vec{spot.x, spot.y, d}

		dif := abs(spot.dir - d)
		if dif > 2 {
			dif = 1
		}

		next.x += dir.x
		next.y += dir.y

		if next.x < 0 || next.y < 0 || next.x >= o.w || next.y >= o.h || o.maze[next.y][next.x] == '#' {
			continue
		}

		o.Walk(next, score+1+(1_000*dif))

	}

}

func (d AocDay16) Puzzle2(useSample int) {

}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}
