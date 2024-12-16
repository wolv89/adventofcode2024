package day14

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type AocDay14 struct{}

const DIR = "day14/"

type Robot struct {
	posx, posy int
	velx, vely int
}

const SECONDS = 100

func (d AocDay14) Puzzle1(useSample int) {

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
		rob  Robot
		line string
		x, y int
	)

	w, h := 101, 103
	if useSample == 1 {
		w, h = 11, 7
	}

	// Using uint8 to penny pinch memory
	// Means we can't have more than 255 robots in one spot...
	// A sensible person would just use a POI (Plain Old Int), or maybe uint
	grid := make([][]uint8, h)
	for g := range grid {
		grid[g] = make([]uint8, w)
	}

	for scanner.Scan() {

		line = scanner.Text()

		rob = AssembleRobot(line)
		rob.Walk(w, h)

		grid[rob.posy][rob.posx]++

	}

	if useSample > 0 {
		for g := range grid {
			var st strings.Builder
			for _, robot := range grid[g] {
				if robot > 0 {
					if robot > 9 {
						st.WriteByte('+')
					} else {
						st.WriteByte('0' + robot)
					}
				} else {
					st.WriteByte('.')
				}
			}
			fmt.Println(st.String())
		}
		fmt.Println("")
		fmt.Println("---")
		fmt.Println("")
	}

	midw, midh := w/2, h/2

	if useSample > 0 {
		for g := range grid {
			if g == midh {
				fmt.Println("")
				continue
			}
			var st strings.Builder
			for r, robot := range grid[g] {
				if r == midw {
					st.WriteByte(' ')
				} else if robot > 0 {
					if robot > 9 {
						st.WriteByte('+')
					} else {
						st.WriteByte('0' + robot)
					}
				} else {
					st.WriteByte('.')
				}
			}
			fmt.Println(st.String())
		}
	}

	// Top left, top right, bottom left, bottom right
	quads := [4]int{0, 0, 0, 0}

	for y = range grid {
		if y == midh {
			continue
		}
		for x = range grid[y] {
			if x == midw || grid[y][x] == 0 {
				continue
			}
			if y < midh {
				if x < midw {
					quads[0] += int(grid[y][x])
				} else {
					quads[1] += int(grid[y][x])
				}
			} else {
				if x < midw {
					quads[2] += int(grid[y][x])
				} else {
					quads[3] += int(grid[y][x])
				}
			}
		}
	}

	if useSample > 0 {
		fmt.Println("")
		fmt.Println(quads)
	}

	total := quads[0] * quads[1] * quads[2] * quads[3]

	fmt.Println("")
	fmt.Println("Total: ", total)

}

func AssembleRobot(str string) Robot {

	parts := strings.Split(str, " ")

	// Trimming the leading p=
	nums := strings.Split(parts[0][2:], ",")

	x, _ := strconv.Atoi(nums[0])
	y, _ := strconv.Atoi(nums[1])

	rob := Robot{x, y, 0, 0}

	// Trimming the leading v=
	nums = strings.Split(parts[1][2:], ",")

	x, _ = strconv.Atoi(nums[0])
	y, _ = strconv.Atoi(nums[1])

	rob.velx, rob.vely = x, y

	return rob

}

func (rob *Robot) Walk(w, h int) {

	newx := (rob.posx + rob.velx*SECONDS) % w
	newy := (rob.posy + rob.vely*SECONDS) % h

	if newx < 0 {
		newx += w
	}
	if newy < 0 {
		newy += h
	}

	rob.posx = newx
	rob.posy = newy

}

func (d AocDay14) Puzzle2(useSample int) {

}
