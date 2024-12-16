package day13

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

type AocDay13 struct{}

const DIR = "day13/"

type Point struct {
	x, y int
}

func (d AocDay13) Puzzle1(useSample int) {

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
		found       [][]string
		BA, BB, PR  Point
		line        string
		cost, total int64
		x, y, a, b  int
		ready       bool
	)

	btregx := regexp.MustCompile(`^Button (A|B): X\+([0-9]+), Y\+([0-9]+)$`)
	prregx := regexp.MustCompile(`^Prize: X=([0-9]+), Y=([0-9]+)$`)

	for scanner.Scan() {

		line = scanner.Text()

		if len(line) == 0 {
			continue
		}

		switch line[:5] {
		case "Butto":
			found = btregx.FindAllStringSubmatch(line, -1)
			x, _ = strconv.Atoi(found[0][2])
			y, _ = strconv.Atoi(found[0][3])
			if found[0][1] == "A" {
				BA = Point{x, y}
			} else {
				BB = Point{x, y}
			}
		case "Prize":
			found = prregx.FindAllStringSubmatch(line, -1)
			x, _ = strconv.Atoi(found[0][1])
			y, _ = strconv.Atoi(found[0][2])
			PR = Point{x, y}
			ready = true
		}

		if ready {
			// fmt.Println("Btn A: {", BA.x, ",", BA.y, "} | Btn B: {", BB.x, ",", BB.y, "} | Prz {", PR.x, ",", PR.y, "}")
			a, b = Calculate(BA, BB, PR)
			if a >= 0 || b >= 0 {
				cost = TokenCost(a, b)
				fmt.Println("Btn A: {", BA.x, ",", BA.y, "} | Btn B: {", BB.x, ",", BB.y, "} | Prz {", PR.x, ",", PR.y, "} | ", a, "*", b, " | ", cost, " | ", total)
				total += cost
			} else {
				fmt.Println("Btn A: {", BA.x, ",", BA.y, "} | Btn B: {", BB.x, ",", BB.y, "} | Prz {", PR.x, ",", PR.y, "} | Unreachable")
			}
			ready = false
			fmt.Println("")
			fmt.Println("--------------------------------")
			fmt.Println("")
		}

	}

	fmt.Println("")
	fmt.Println("Total: ", total)

}

func Calculate(BA, BB, PR Point) (int, int) {

	var (
		cost       int64
		maxa, maxb int
		a, b, x, y int
	)

	cheapest := int64(math.MaxInt64)
	chpa, chpb := math.MaxInt, math.MaxInt

	// Find maximum possible presses for each button
	// Ie the lowest presses that keeps us under (or equal?) to the prize target
	// By definition if we press more than that, then we'll overshoot...

	maxa = PR.x / BA.x
	maxa = min(maxa, PR.y/BA.y)

	maxb = PR.x / BB.x
	maxb = min(maxb, PR.y/BB.y)

	// B presses are cheaper, so start with them maximised and work backwards
	// But should this be a binary search...?
	for b = maxb; b >= 0; b-- {
		for a = maxb - b; a < maxa; a++ {

			x = BB.x*b + BA.x*a
			y = BB.y*b + BA.y*a

			// fmt.Println(maxa, ",", maxb, " | ", a, ",", b, " | ", x, ",", y, " | ", PR.x, ",", PR.y)

			if x == PR.x && y == PR.y {
				cost = TokenCost(a, b)
				if cost < cheapest {
					chpa, chpb = a, b
					cheapest = cost
				}
			} else if x > PR.x || y > PR.y {
				break
			}

		}
	}

	if chpa > maxa || chpb > maxb {
		return -1, -1
	}

	return chpa, chpb

}

func TokenCost(a, b int) int64 {
	return int64(a*3 + b)
}

func (d AocDay13) Puzzle2(useSample int) {

}
