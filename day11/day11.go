package day11

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type AocDay11 struct{}

const DIR = "day11/"

func (d AocDay11) Puzzle1(useSample bool) {

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
	scanner.Scan()
	line := scanner.Text()

	start := strings.Split(line, " ")
	stones := make([]int64, 0, len(start))

	for _, st := range start {

		stone, err := strconv.ParseInt(st, 10, 64)
		if err != nil {
			fmt.Println("Unable to convert... ", st, " | ", err)
			continue
		}

		stones = append(stones, stone)

	}

	fmt.Println(stones)

	fmt.Println("")

}

func (d AocDay11) Puzzle2(useSample bool) {

}
