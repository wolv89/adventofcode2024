package day18

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

	priorityQueue "github.com/emirpasic/gods/queues/priorityqueue"
)

type AocDay18 struct{}

const DIR = "day18/"

type Point struct {
	x, y int
}

type Step struct {
	coord Point
	score int
	path  map[Point]int
}

var (
	SIZE  = 70
	BYTES = 1024

	DIRS = [4]Point{
		{0, -1},
		{0, 1},
		{-1, 0},
		{1, 0},
	}
)

func (d AocDay18) Puzzle1(useSample int) {

	datafile := DIR + "data.txt"
	if useSample == 1 {
		datafile = DIR + "sample.txt"
		SIZE = 6
		BYTES = 12
	}

	f, err := os.Open(datafile)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var (
		coords     []string
		line       string
		l, g, x, y int
	)

	grid := make([][]byte, SIZE+1)
	for g = range grid {
		grid[g] = bytes.Repeat([]byte("."), SIZE+1)
	}

	for scanner.Scan() {

		line = scanner.Text()

		coords = strings.Split(line, ",")
		if len(coords) < 2 {
			fmt.Println("PROBLEM splitting line: ", line)
			continue
		}

		x, y = NumToInt(coords[0]), NumToInt(coords[1])

		if x < 0 || y < 0 {
			fmt.Println("PROBLEM parsing line: ", line)
			continue
		}

		grid[y][x] = '#'

		l++

		if l >= BYTES {
			break
		}

	}

	score, path := dijkstras(grid)

	fmt.Println("")

	for p, _ := range path {
		grid[p.y][p.x] = 'O'
	}

	fmt.Println("")

	for g = range grid {
		fmt.Println(string(grid[g]))
	}

	fmt.Println("")

	fmt.Println("Steps: ", score)

	fmt.Println("")

}

func (d AocDay18) Puzzle2(useSample int) {

}

func dijkstras(grid [][]byte) (int, map[Point]int) {

	pq := priorityQueue.NewWith(func(a, b interface{}) int {
		return a.(Step).score - b.(Step).score
	})

	pq.Enqueue(Step{
		Point{0, 0}, // Hard coded start, top left
		0,
		make(map[Point]int),
	})
	visited := make(map[Point]struct{})

	var (
		element interface{}
		next    []Step
		node, n Step
		ok      bool
	)

	for !pq.Empty() {

		element, _ = pq.Dequeue()
		node = element.(Step)

		if _, ok = visited[node.coord]; ok {
			continue
		}

		node.path[node.coord] = node.score

		// Bottom right - exit
		if node.coord.x == SIZE && node.coord.y == SIZE {
			return node.score, node.path
		}

		next = getNext(node, grid, visited)
		for _, n = range next {
			pq.Enqueue(n)
		}

		visited[node.coord] = struct{}{}

	}

	return -1, make(map[Point]int)

}

func getNext(node Step, grid [][]byte, visited map[Point]struct{}) []Step {

	next := make([]Step, 0)

	var (
		newCoord Point
		newScore int
		ok       bool
	)

	for _, dir := range DIRS {

		newCoord = Point{
			node.coord.x + dir.x,
			node.coord.y + dir.y,
		}

		if newCoord.x < 0 || newCoord.y < 0 || newCoord.x > SIZE || newCoord.y > SIZE || grid[newCoord.y][newCoord.x] == '#' {
			continue
		}

		if _, ok = visited[newCoord]; ok {
			continue
		}

		newScore = node.score + 1

		next = append(next, Step{
			newCoord,
			newScore,
			copyMap(node.path),
		})

	}

	return next

}

func NumToInt(num string) int {

	if len(num) > 2 {
		return -1
	} else if len(num) == 1 {
		return int(num[0] - '0')
	}

	return int(num[1]-'0') + int(num[0]-'0')*10

}

func copyMap(path map[Point]int) map[Point]int {
	new := make(map[Point]int, len(path))
	for key, value := range path {
		new[key] = value
	}
	return new
}
