package day11

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type AocDay11 struct{}

const DIR = "day11/"

func (d AocDay11) Puzzle1(useSample int) {

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
	scanner.Scan()
	line := scanner.Text()

	start := strings.Split(line, " ")
	stones := make([]int64, 0, len(start))

	var (
		st        string
		stone, mp int64
		b, dg     int
	)

	for _, st = range start {

		stone, err = strconv.ParseInt(st, 10, 64)
		if err != nil {
			fmt.Println("Unable to convert... ", st, " | ", err)
			continue
		}

		stones = append(stones, stone)

	}

	fmt.Println(stones)
	fmt.Println("---")

	const BLINKS = 50

	for b = 0; b < BLINKS; b++ {

		next := make([]int64, 0, len(stones))

		for _, stone = range stones {

			if stone == 0 {
				next = append(next, 1)
				continue
			}

			dg = countDigits(stone)

			if dg%2 == 0 {
				mp = int64(math.Pow10(dg / 2))
				next = append(next, stone/mp)
				next = append(next, stone%mp)
				continue
			}

			next = append(next, stone*2024)

		}

		stones = next

		if BLINKS <= 6 {
			fmt.Println(stones)
		} else {
			fmt.Println(b+1, ": ", len(stones))
		}

	}

	fmt.Println("")

}

func countDigits(num int64) int {

	d, m := 0, int64(1)

	for m <= num {
		m *= 10
		d++
	}

	return d

}

func (d AocDay11) Puzzle2(useSample int) {

}
