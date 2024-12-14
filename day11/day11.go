package day11

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
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

type Counter struct {
	count int64
	mu    sync.RWMutex
}

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
		st    string
		stone int64
		cnt   Counter
		wg    sync.WaitGroup
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

	// Should probably make this a CLI flag...
	const BLINKS = 45

	wg.Add(1)

	go func() {
		defer wg.Done()
		countStones(stones, 0, BLINKS, &cnt)
	}()

	wg.Wait()

	fmt.Println("")
	fmt.Println("Count: ", cnt.count)

}

func countStones(stones []int64, blink, BLINKS int, cnt *Counter) {

	next := make([]int64, 0, len(stones))

	// fmt.Println(stones, " | ", blink, "/", BLINKS)

	for _, stone := range stones {

		if stone == 0 {
			next = append(next, 1)
			continue
		}

		dg := countDigits(stone)

		if dg%2 == 0 {
			mp := int64(math.Pow10(dg / 2))
			next = append(next, stone/mp)
			next = append(next, stone%mp)
			continue
		}

		next = append(next, stone*2024)

	}

	blink++

	if blink == BLINKS {

		cnt.mu.Lock()
		cnt.count += int64(len(next))
		cnt.mu.Unlock()

		return

	}

	if len(next) <= 10 {
		countStones(next, blink, BLINKS, cnt)
	} else {
		mid := len(next) / 2

		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			countStones(next[:mid], blink, BLINKS, cnt)
		}()

		go func() {
			defer wg.Done()
			countStones(next[mid:], blink, BLINKS, cnt)
		}()

		wg.Wait()
	}

}
