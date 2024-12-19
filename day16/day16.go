package day16

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	priorityQueue "github.com/emirpasic/gods/queues/priorityqueue"
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
	best []Point
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

/*
 * CREDIT: ShraddhaAg
 * https://www.reddit.com/r/adventofcode/comments/1hfboft/comment/m2au9qu/?utm_source=share&utm_medium=web3x&utm_name=web3xcss&utm_term=1&utm_content=share_button
 *
 * Gave this a crack with my rough approximations - got it working on the test data but was running forever on the real thing
 * Maybe it would have worked if I left it for an hour (or a few?!) who knows. Anyway, ground up re-write based on this much nicer
 * solution by ShraddhaAg
 */

var (
	Up    = Point{0, -1}
	Down  = Point{0, 1}
	Left  = Point{-1, 0}
	Right = Point{1, 0}
)

type Step struct {
	co, lastDir Point
	score       int
	path        map[Point]int
}

type Vector struct {
	co, lastDir Point
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

	var line string
	maze := make([][]byte, 0)

	for scanner.Scan() {
		line = scanner.Text()
		maze = append(maze, []byte(line))
	}

	score, path := dijkstras(maze)

	fmt.Println(score)
	// fmt.Println(path)
	fmt.Println("")

	// Works on sample, but the answer was off slightly on the full data...
	spots := getUniqueCount(maze, path)
	fmt.Println(spots)

}

func dijkstras(grid [][]byte) (int, map[Point]int) {

	pq := priorityQueue.NewWith(func(a, b interface{}) int {
		return a.(Step).score - b.(Step).score
	})

	// Start pos is always bottom left corner (inset from # border)
	pq.Enqueue(Step{Point{1, len(grid) - 2}, Right, 0, make(map[Point]int)})

	visited := make(map[Vector]struct{})

	for !pq.Empty() {

		element, _ := pq.Dequeue()
		currentNode := element.(Step)

		if _, ok := visited[Vector{currentNode.co, currentNode.lastDir}]; ok {
			continue
		}

		currentNode.path[currentNode.co] = currentNode.score

		if grid[currentNode.co.y][currentNode.co.x] == 'E' {
			return currentNode.score, currentNode.path
		}

		nextSteps := getNextSteps(currentNode, grid, visited)
		for _, n := range nextSteps {
			pq.Enqueue(n)
		}

		visited[Vector{currentNode.co, currentNode.lastDir}] = struct{}{}

	}

	return -1, make(map[Point]int)

}

func getNextSteps(current Step, grid [][]byte, visited map[Vector]struct{}) []Step {

	possibleNext := []Step{}

	checkDirs := make([]Point, 0, 3)

	switch current.lastDir {
	case Up:
		checkDirs = append(checkDirs, Up, Left, Right)
	case Down:
		checkDirs = append(checkDirs, Down, Left, Right)
	case Left:
		checkDirs = append(checkDirs, Up, Left, Down)
	case Right:
		checkDirs = append(checkDirs, Up, Right, Down)
	}

	for _, dir := range checkDirs {

		newPos := Point{current.co.x + dir.x, current.co.y + dir.y}

		if newPos.x < 1 || newPos.y < 1 || newPos.x >= len(grid[0])-1 || newPos.y >= len(grid) || grid[newPos.y][newPos.x] == '#' {
			continue
		}

		if _, ok := visited[Vector{newPos, dir}]; ok {
			continue
		}

		score := current.score + 1
		if dir != current.lastDir {
			score += 1000
		}

		possibleNext = append(possibleNext, Step{
			co:      newPos,
			lastDir: dir,
			score:   score,
			path:    copyMap(current.path),
		})

	}

	return possibleNext

}

func getUniqueCount(grid [][]byte, path map[Point]int) int {

	pq := priorityQueue.NewWith(func(a, b interface{}) int {
		return a.(Step).score - b.(Step).score
	})

	pq.Enqueue(Step{Point{1, len(grid) - 2}, Right, 0, make(map[Point]int)})

	visited := make(map[Vector]struct{})
	safeCoords := make(map[Point]struct{})

	for !pq.Empty() {

		element, _ := pq.Dequeue()
		currentNode := element.(Step)

		if score, ok := path[currentNode.co]; ok && score == currentNode.score {
			for point := range currentNode.path {
				if _, ok := path[point]; !ok {
					safeCoords[point] = struct{}{}
				}
			}
		}

		if _, ok := visited[Vector{currentNode.co, currentNode.lastDir}]; ok {
			continue
		}

		currentNode.path[currentNode.co] = currentNode.score

		if grid[currentNode.co.y][currentNode.co.x] == 'E' {
			continue
		}

		nextSteps := getNextSteps(currentNode, grid, visited)
		for _, n := range nextSteps {
			pq.Enqueue(n)
		}

		visited[Vector{currentNode.co, currentNode.lastDir}] = struct{}{}

	}

	return len(path) + len(safeCoords)

}

func copyMap(path map[Point]int) map[Point]int {
	new := make(map[Point]int, len(path))
	for key, value := range path {
		new[key] = value
	}
	return new
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}
