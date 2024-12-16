package day13

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type AocDay13 struct{}

const DIR = "day13/"

type Point struct {
	x, y int
}

type Point64 struct {
	x, y int64
}

const OFFSET int64 = 10_000_000_000_000

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
			a, b = Calculate(BA, BB, PR)
			if a >= 0 || b >= 0 {
				cost = TokenCost(a, b)
				fmt.Println("Btn A: {", BA.x, ",", BA.y, "} | Btn B: {", BB.x, ",", BB.y, "} | Prz {", PR.x, ",", PR.y, "} | ", a, "*", b, " | ", cost, " | ", total)
				total += cost
			} else {
				fmt.Println("Btn A: {", BA.x, ",", BA.y, "} | Btn B: {", BB.x, ",", BB.y, "} | Prz {", PR.x, ",", PR.y, "} | Unreachable")
			}
			ready = false
		}

	}

	fmt.Println("")
	fmt.Println("Total: ", total)

}

/*
 * Credit: ThunderChaser
 * https://www.reddit.com/r/adventofcode/comments/1hd7irq/2024_day_13_an_explanation_of_the_mathematics/
 *
 * Simplified to use Cramer's rule, as illustrated in the post above <3
 */
func Calculate(BA, BB, PR Point) (int, int) {

	a := (PR.x*BB.y - PR.y*BB.x) / (BA.x*BB.y - BA.y*BB.x)
	b := (PR.y*BA.x - PR.x*BA.y) / (BA.x*BB.y - BA.y*BB.x)

	if a*BA.x+b*BB.x != PR.x || a*BA.y+b*BB.y != PR.y {
		return -1, -1
	}

	return a, b

}

func TokenCost(a, b int) int64 {
	return int64(a*3 + b)
}

func (d AocDay13) Puzzle2(useSample int) {

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
		BA, BB, PR  Point64
		line        string
		cost, total int64
		x, y, a, b  int64
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
			x, _ = strconv.ParseInt(found[0][2], 10, 64)
			y, _ = strconv.ParseInt(found[0][3], 10, 64)
			if found[0][1] == "A" {
				BA = Point64{x, y}
			} else {
				BB = Point64{x, y}
			}
		case "Prize":
			found = prregx.FindAllStringSubmatch(line, -1)
			x, _ = strconv.ParseInt(found[0][1], 10, 64)
			y, _ = strconv.ParseInt(found[0][2], 10, 64)
			PR = Point64{x + OFFSET, y + OFFSET}
			ready = true
		}

		if ready {
			a, b = Calculate64(BA, BB, PR)
			if a >= 0 || b >= 0 {
				cost = a*3 + b
				fmt.Println("Btn A: {", BA.x, ",", BA.y, "} | Btn B: {", BB.x, ",", BB.y, "} | Prz {", PR.x, ",", PR.y, "} | ", a, "*", b, " | ", cost, " | ", total)
				total += cost
			} else {
				fmt.Println("Btn A: {", BA.x, ",", BA.y, "} | Btn B: {", BB.x, ",", BB.y, "} | Prz {", PR.x, ",", PR.y, "} | Unreachable")
			}
			ready = false
		}

	}

	fmt.Println("")
	fmt.Println("Total: ", total)

}

func Calculate64(BA, BB, PR Point64) (int64, int64) {

	a := (PR.x*BB.y - PR.y*BB.x) / (BA.x*BB.y - BA.y*BB.x)
	b := (PR.y*BA.x - PR.x*BA.y) / (BA.x*BB.y - BA.y*BB.x)

	if a*BA.x+b*BB.x != PR.x || a*BA.y+b*BB.y != PR.y {
		return -1, -1
	}

	return a, b

}
