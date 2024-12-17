package day15

import (
	"bufio"
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type AocDay15 struct{}

const DIR = "day15/"

const (
	ROBOT = '@'
	BOX   = 'O'
	WALL  = '#'
	SPACE = '.'

	SPEED = 10

	MOVEUP    = '^'
	MOVEDOWN  = 'v'
	MOVELEFT  = '<'
	MOVERIGHT = '>'
)

type Point struct {
	x, y int
}

func (d AocDay15) Puzzle1(useSample int) {

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
		rob           Point
		line          string
		x, y, w, h    int
		readWarehouse bool
	)

	warehouse := make([][]byte, 0)
	moves := make([]byte, 0)

	for scanner.Scan() {

		line = scanner.Text()

		if !readWarehouse {

			if len(line) == 0 {
				readWarehouse = true
				continue
			}

			warehouse = append(warehouse, []byte(line))

		} else {

			moves = append(moves, []byte(line)...)

		}

	}

	w, h = len(warehouse[0]), len(warehouse)

findstart:
	for y = 0; y < h; y++ {
		for x = 0; x < w; x++ {
			if warehouse[y][x] == ROBOT {
				rob = Point{x, y}
				break findstart
			}
		}
	}

	p := tea.NewProgram(model{
		sub:       make(chan struct{}),
		w:         w,
		h:         h,
		robot:     rob,
		warehouse: warehouse,
		moves:     moves,
		result:    "",
	})

	if _, err = p.Run(); err != nil {
		fmt.Println("Quitting...")
		os.Exit(1)
	}

}

func (d AocDay15) Puzzle2(useSample int) {

}
