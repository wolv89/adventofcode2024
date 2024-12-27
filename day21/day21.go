package day21

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type AocDay21 struct{}

const DIR = "day21/"

func (d AocDay21) Puzzle1(useSample int) {

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
		line, seq  string
		b          strings.Builder
		total      int64
		l, num     int
		last, next byte
	)

	for scanner.Scan() {

		line = scanner.Text()

		// Always start pointing to the A
		last = 'A'
		b.Reset()

		// NUM PAD
		for l = 0; l < len(line); l++ {
			next = line[l]
			b.WriteString(NUMPAD[last][next]) // Safety's off...
			b.WriteByte('A')
			last = next
		}

		seq = b.String()
		fmt.Println(line, " | (", len(seq), ") ", seq)

		b.Reset()
		last = 'A'

		// CON PAD 1
		for l = 0; l < len(seq); l++ {
			next = seq[l]
			b.WriteString(CONPAD[last][next])
			b.WriteByte('A')
			last = next
		}

		seq = b.String()
		fmt.Println(line, " | (", len(seq), ") ", seq)

		b.Reset()
		last = 'A'

		// CON PAD 2
		for l = 0; l < len(seq); l++ {
			next = seq[l]
			b.WriteString(CONPAD[last][next])
			b.WriteByte('A')
			last = next
		}

		seq = b.String()
		fmt.Println(line, " | (", len(seq), ") ", seq)

		num, err = strconv.Atoi(line[:len(line)-1])
		if err != nil {
			fmt.Println("Unable to parse: ", line, " | ", err)
			continue
		}

		fmt.Println("> ", num)

		total += int64(num * len(seq))

		fmt.Println("")

	}

	fmt.Println("")
	fmt.Println("Total: ", total)

}

func (d AocDay21) Puzzle2(useSample int) {

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
		line   string
		total  int64
		l, num int
	)

	cache := make(map[OpSeq]int)

	for scanner.Scan() {

		line = scanner.Text()

		l = Solve(line, 0, &cache)

		num, err = strconv.Atoi(line[:len(line)-1])
		if err != nil {
			fmt.Println("Unable to parse: ", line, " | ", err)
			continue
		}

		fmt.Println(l, " * ", num)

		total += int64(num * l)

		fmt.Println("")

	}

	fmt.Println("")
	fmt.Println("Total: ", total)

}

type OpSeq struct {
	seq   string
	depth int
}

const ROBOTS = 25

/*
 * (Part 2)
 * CREDIT: Boojum
 * https://www.reddit.com/r/adventofcode/comments/1hj2odw/comment/m3482ai/?utm_source=share&utm_medium=web3x&utm_name=web3xcss&utm_term=1&utm_content=share_button
 *
 * Huge thanks for this post - was able to roughly translate the python iterators into my clunky Go code below, but it definitely helped me understand the branching
 * Ie how to actually split this problem into solvable (and repeatable/cacheable) steps...
 */

func Solve(seq string, depth int, cache *map[OpSeq]int) int {

	ops := OpSeq{seq, depth}

	if _, ok := (*cache)[ops]; ok {
		return (*cache)[ops]
	}

	if depth > ROBOTS {
		return len(seq)
	}

	seq = "A" + seq

	var l int

	for s := 1; s < len(seq); s++ {
		l += Solve(Path(seq[s-1], seq[s], depth == 0), depth+1, cache)
	}

	(*cache)[ops] = l

	return l

}

func Path(from, to byte, numpad bool) string {

	var b strings.Builder

	if numpad {
		b.WriteString(NUMPAD[from][to])
	} else {
		b.WriteString(CONPAD[from][to])
	}
	b.WriteByte('A')

	return b.String()

}
