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
	} else if useSample == 4 {
		datafile = DIR + "sample4.txt"
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

	const BLINKS = 10

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

		if BLINKS <= 10 {
			fmt.Println(len(stones), " ", stones)
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

type StoneKey struct {
	stone int64
	depth int
}

/*
 * CREDIT: tlareg
 * https://www.reddit.com/r/adventofcode/comments/1hbm0al/comment/m1xsgsa/?utm_source=share&utm_medium=web3x&utm_name=web3xcss&utm_term=1&utm_content=share_button
 *
 * Thought Goroutines might just do the magic for me, but they didn't really help at all!
 * This one is a bit hefty to brute force. Adapted from a caching/recursive solution by tlareg
 */
func (d AocDay11) Puzzle2(useSample int) {

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
		st           string
		stone, total int64
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

	const BLINKS = 75

	cache := make(map[StoneKey]int64)

	for _, stone = range stones {
		total += countStones(stone, BLINKS, &cache)
	}

	fmt.Println("Total: ", total)

}

func countStones(stone int64, depth int, cache *map[StoneKey]int64) int64 {

	if depth == 0 {
		return 1
	}

	key := StoneKey{
		stone,
		depth,
	}

	if _, ok := (*cache)[key]; ok {
		return (*cache)[key]
	}

	save := func(res int64) int64 {
		(*cache)[key] = res
		return res
	}

	if stone == 0 {
		return save(countStones(1, depth-1, cache))
	}

	dg := countDigits(stone)

	if dg%2 == 0 {
		mp := int64(math.Pow10(dg / 2))
		return save(countStones(stone/mp, depth-1, cache) + countStones(stone%mp, depth-1, cache))
	}

	return save(countStones(stone*2024, depth-1, cache))

}
