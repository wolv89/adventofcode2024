package day20

import (
	"bufio"
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type AocDay20 struct{}

const DIR = "day20/"

const (
	STARTCHAR = 'S'
	ENDCHAR   = 'E'
)

type Point struct {
	x, y int
}

type CheatPair struct {
	count, saving int
}

type CheatPoint struct {
	a, b, c Point
}

// Right, down, left, up
var DIRS = [4][2]int{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

func (d AocDay20) Puzzle1(useSample int) {

	datafile := DIR + "data.txt"
	if useSample > 0 {
		d.Puzzle1Sample(useSample)
		return
	}

	f, err := os.Open(datafile)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var (
		start, loc, next, cheat Point
		line                    string
		w, h, dr, step, save    int
		foundStart              bool
	)

	path := make([][]byte, 0)

	for scanner.Scan() {

		line = scanner.Text()

		path = append(path, []byte(line))

		if !foundStart {
			for w = 0; w < len(line); w++ {
				if line[w] == STARTCHAR {
					start = Point{w, h}
					foundStart = true
				}
			}
			h++
		}

	}

	w, h = len(path[0]), len(path)
	loc = Point{start.x, start.y}
	next, cheat = Point{0, 0}, Point{0, 0}

	dist := make([][]int, h)
	vis := make([][]bool, h)

	for dr = range dist {
		dist[dr] = make([]int, w)
		vis[dr] = make([]bool, w)
	}

	// --- CYCLE 1: Count steps to exit ---
	for {

		for dr = range DIRS {

			next.x, next.y = loc.x+DIRS[dr][0], loc.y+DIRS[dr][1]

			if next.x < 0 || next.y < 0 || next.x >= w || next.y >= h || vis[next.y][next.x] || path[next.y][next.x] == '#' {
				continue
			} else {
				break
			}

		}

		vis[loc.y][loc.x] = true

		step++
		loc.x, loc.y = next.x, next.y
		dist[loc.y][loc.x] = step

		if path[loc.y][loc.x] == ENDCHAR {
			break
		}

	}

	for dr = range vis {
		clear(vis[dr])
	}

	loc = Point{start.x, start.y}
	superCheats := 0

	// --- CYCLE 2: Count cheats ---
	for {

		for dr = range DIRS {

			next.x, next.y = loc.x+DIRS[dr][0], loc.y+DIRS[dr][1]
			cheat.x, cheat.y = next.x+DIRS[dr][0], next.y+DIRS[dr][1]

			// Next check is saying the cell inbewtween must be a wall, can't just jump two spaces ahead in an unobstructed path
			if cheat.x < 0 || cheat.y < 0 || cheat.x >= w || cheat.y >= h || path[next.y][next.x] != '#' || dist[cheat.y][cheat.x] == 0 {
				continue
			}

			save = dist[cheat.y][cheat.x] - dist[loc.y][loc.x] - 2
			if save >= 100 {
				superCheats++
			}

		}

		if path[loc.y][loc.x] == ENDCHAR {
			break
		}

		for dr = range DIRS {

			next.x, next.y = loc.x+DIRS[dr][0], loc.y+DIRS[dr][1]

			if next.x < 0 || next.y < 0 || next.x >= w || next.y >= h || vis[next.y][next.x] || path[next.y][next.x] == '#' {
				continue
			} else {
				break
			}

		}

		vis[loc.y][loc.x] = true
		loc.x, loc.y = next.x, next.y

	}

	fmt.Println("")
	fmt.Println("SUPER CHEATS")
	fmt.Println("------------")
	fmt.Println("(Save 100 or more picoseconds)")
	fmt.Println(">", superCheats)

}

/*
 * Copy pasted function
 * Adds a lot of extra logging output, easier to just split
 */
func (d AocDay20) Puzzle1Sample(useSample int) {

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
		start, loc, next, cheat Point
		line                    string
		w, h, dr, step, save    int
		foundStart              bool
	)

	path := make([][]byte, 0)

	for scanner.Scan() {

		line = scanner.Text()

		path = append(path, []byte(line))

		if !foundStart {
			for w = 0; w < len(line); w++ {
				if line[w] == STARTCHAR {
					start = Point{w, h}
					foundStart = true
				}
			}
			h++
		}

	}

	if useSample > 0 {
		for _, row := range path {
			fmt.Println(string(row))
		}
		fmt.Println("")
	}

	w, h = len(path[0]), len(path)
	loc = Point{start.x, start.y}
	next, cheat = Point{0, 0}, Point{0, 0}

	dist := make([][]int, h)
	vis := make([][]bool, h)

	for dr = range dist {
		dist[dr] = make([]int, w)
		vis[dr] = make([]bool, w)
	}

	// --- CYCLE 1: Count steps to exit ---
	for {

		for dr = range DIRS {

			next.x, next.y = loc.x+DIRS[dr][0], loc.y+DIRS[dr][1]

			if next.x < 0 || next.y < 0 || next.x >= w || next.y >= h || vis[next.y][next.x] || path[next.y][next.x] == '#' {
				continue
			} else {
				break
			}

		}

		vis[loc.y][loc.x] = true

		step++
		loc.x, loc.y = next.x, next.y
		dist[loc.y][loc.x] = step

		if path[loc.y][loc.x] == ENDCHAR {
			break
		}

	}

	fmt.Println("")

	if useSample > 0 {
		for dr = range dist {
			for _, cell := range dist[dr] {
				if cell > 0 {
					fmt.Printf("%2d ", cell)
				} else {
					fmt.Print(" . ")
				}
			}
			fmt.Print("\n")
		}
	}

	fmt.Println("")

	for dr = range vis {
		clear(vis[dr])
	}

	cheats := make(map[int]int)
	cheatPoints := make(map[int][]CheatPoint)

	loc = Point{start.x, start.y}

	// --- CYCLE 2: Count cheats ---
	for {

		for dr = range DIRS {

			next.x, next.y = loc.x+DIRS[dr][0], loc.y+DIRS[dr][1]
			cheat.x, cheat.y = next.x+DIRS[dr][0], next.y+DIRS[dr][1]

			// Next check is saying the cell inbewtween must be a wall, can't just jump two spaces ahead in an unobstructed path
			if cheat.x < 0 || cheat.y < 0 || cheat.x >= w || cheat.y >= h || path[next.y][next.x] != '#' || dist[cheat.y][cheat.x] == 0 {
				continue
			}

			// Minus two because you still have to actually take two steps
			save = dist[cheat.y][cheat.x] - dist[loc.y][loc.x] - 2
			if save > 0 {
				cheats[save]++
				cheatPoints[save] = append(cheatPoints[save], CheatPoint{loc, next, cheat})
			}

		}

		if path[loc.y][loc.x] == ENDCHAR {
			break
		}

		for dr = range DIRS {

			next.x, next.y = loc.x+DIRS[dr][0], loc.y+DIRS[dr][1]

			if next.x < 0 || next.y < 0 || next.x >= w || next.y >= h || vis[next.y][next.x] || path[next.y][next.x] == '#' {
				continue
			} else {
				break
			}

		}

		vis[loc.y][loc.x] = true
		loc.x, loc.y = next.x, next.y

	}

	cps := make([]CheatPair, 0)

	for key, val := range cheats {
		cps = append(cps, CheatPair{val, key})
	}

	slices.SortFunc(cps, func(a, b CheatPair) int {
		return cmp.Compare(a.saving, b.saving)
	})

	for _, cp := range cps {
		if cp.count > 1 {
			fmt.Println(cp.count, "cheats that save", cp.saving, "picoseconds")
		} else {
			fmt.Println("1 cheat that saves", cp.saving, "picoseconds")
		}
	}

	cmap := make([][]byte, h)
	for dr = range cmap {
		cmap[dr] = []byte(strings.Repeat(".", w))
	}

	for _, cpoint := range cheatPoints[4] {
		cmap[cpoint.a.y][cpoint.a.x] = 'X'
		cmap[cpoint.b.y][cpoint.b.x] = 'O'
		cmap[cpoint.c.y][cpoint.c.x] = 'X'
	}

	fmt.Println("")
	for _, row := range cmap {
		fmt.Println(string(row))
	}

}

func (d AocDay20) Puzzle2(useSample int) {

}
