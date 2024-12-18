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

	BOXLEFT  = '['
	BOXRIGHT = ']'

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

	datafile := DIR + "data.txt"
	if useSample == 1 {
		datafile = DIR + "sample.txt"
	} else if useSample == 2 {
		datafile = DIR + "sample2.txt"
	} else if useSample == 3 {
		datafile = DIR + "sample3.txt"
	}

	f, err := os.Open(datafile)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var (
		row           []byte
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

			row = make([]byte, 0, len(line)*2)

			for x = 0; x < len(line); x++ {
				if line[x] == BOX {
					row = append(row, '[', ']')
				} else if line[x] == ROBOT {
					row = append(row, '@', '.')
					rob = Point{x * 2, y}
				} else {
					row = append(row, line[x], line[x])
				}
			}

			warehouse = append(warehouse, row)
			y++

		} else {

			moves = append(moves, []byte(line)...)

		}

	}

	w, h = len(warehouse[0]), len(warehouse)

	// Decided to add the log because my solution wasn't working
	// While adding the log I fixed a bug, which fixed the solution.... faaaar out

	/*
		lf, lerr := tea.LogToFile("log/debug-"+time.Now().Format("20060102")+".log", "log")
		if lerr != nil {
			fmt.Println("Fatal:", lerr)
			os.Exit(1)
		}
		defer lf.Close()

		log.Println("+---+ = +---+ = +---+ = +---+ = +---+")
		log.Println("--           NEW SESSION           --")
		log.Println("+---+ = +---+ = +---+ = +---+ = +---+")
	*/

	p := tea.NewProgram(modelt{
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
