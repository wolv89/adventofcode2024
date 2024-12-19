package day16

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
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
	path   [][]int
	maze   [][]byte
	goal   Point
	w, h   int
	chains [][]Point
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
		nil,
	}

	oly.Walk(start, 1)

	if useSample > 0 {
		fmt.Println("")
		for _, pt := range oly.path {
			for _, tp := range pt {
				fmt.Printf("%7d ", tp)
			}
			fmt.Print("\n\n")
		}
		fmt.Println("")
	}

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
		make([][]Point, 0),
	}

	oly.Walk(start, 1)

	/*
		if useSample > 0 {
			fmt.Println("")
			for _, pt := range oly.path {
				for _, tp := range pt {
					fmt.Printf("%7d ", tp)
				}
				fmt.Print("\n\n")
			}
			fmt.Println("")
		}
	*/

	fmt.Println("")
	fmt.Println("Score: ", oly.path[end.y][end.x]-1)

	// There's an "off by one" in my scoring here... I think I'm counting the start space as a 1, incorrectly...
	oly.Observe(start, 1, oly.path[end.y][end.x], make([]Point, 0), make([]bool, oly.w*oly.h))

	fmt.Println("")
	fmt.Println("---")
	fmt.Println("")

	bestmaze := make([][]byte, h)
	for b := range bestmaze {
		bestmaze[b] = make([]byte, w)
		copy(bestmaze[b], maze[b])
	}

	var count int

	for _, ch := range oly.chains {
		for _, pt := range ch {
			if bestmaze[pt.y][pt.x] != 'O' {
				bestmaze[pt.y][pt.x] = 'O'
				count++
			}
		}
	}

	if useSample > 0 {
		fmt.Println(strings.Repeat("#", w+2))
		for b := range bestmaze {
			fmt.Println("#" + string(bestmaze[b]) + "#")
		}
		fmt.Println(strings.Repeat("#", w+2))
	}

	fmt.Println("")
	fmt.Println("Spots: ", count)

}

func (o *Olympics) Observe(spot Vec, score, target int, chain []Point, seen []bool) {

	if score > target {
		return
	}

	chain = append(chain, Point{spot.x, spot.y})
	seen[o.h*spot.y+spot.x] = true

	if len(chain)%100 == 0 {
		fmt.Println("Working... ", len(chain))
	}

	// Made it...
	if spot.x == o.goal.x && spot.y == o.goal.y {
		o.chains = append(o.chains, chain)
		return
	}

	nextdirs := make([]Vec, 0)

	for d, dir := range DIRS {

		next := Vec{spot.x, spot.y, d}

		next.x += dir.x
		next.y += dir.y

		if next.x < 0 || next.y < 0 || next.x >= o.w || next.y >= o.h || o.maze[next.y][next.x] == '#' || seen[o.h*next.y+next.x] {
			continue
		}

		nextdirs = append(nextdirs, next)

	}

	// Continue current chain
	if len(nextdirs) == 1 {

		dif := abs(spot.dir - nextdirs[0].dir)
		if dif > 2 {
			dif = 1
		}

		o.Observe(nextdirs[0], score+1+(1_000*dif), target, chain, seen)
		return

	}

	// Else, copy slices and branch to new chains

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		for _, next := range nextdirs {

			newchain := make([]Point, len(chain))
			copy(newchain, chain)

			newseen := make([]bool, len(seen))
			copy(newseen, seen)

			dif := abs(spot.dir - next.dir)
			if dif > 2 {
				dif = 1
			}

			o.Observe(next, score+1+(1_000*dif), target, newchain, newseen)

		}
	}()

	wg.Wait()

}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}
