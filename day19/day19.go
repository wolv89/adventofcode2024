package day19

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/wolv89/adventofcode2024/structures"
)

type AocDay19 struct{}

const DIR = "day19/"

func (d AocDay19) Puzzle1(useSample int) {

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
		rawPatterns []string
		line        string
		total       int
		valid       bool
	)

	for scanner.Scan() {

		line = scanner.Text()

		if len(line) == 0 {
			break
		}

		rawPatterns = strings.Split(line, ", ")

	}

	// Prefix Tree (Trie)
	pt := structures.NewTrie()

	for _, pattern := range rawPatterns {
		pt.Insert(pattern)
	}

	for scanner.Scan() {

		line = scanner.Text()
		valid = isValidPattern(line, pt)

		if valid {
			fmt.Println("[x]", line)
			total++
		} else {
			fmt.Println("[ ]", line)
		}

	}

	fmt.Println("")
	fmt.Println("Total: ", total)

}

func (d AocDay19) Puzzle2(useSample int) {

}

func isValidPattern(pattern string, pt structures.Trie) bool {

	// fmt.Println("> ", pattern)

	p := len(pattern)

	m := min(pt.Longest(), p)
	matches := pt.FindSubmatches(pattern[:m])

	n := len(matches)

	if n == 0 {
		return false
	}

	var sn int

	for i := n - 1; i >= 0; i-- {

		sn = len(matches[i])

		// Matched the full remaining part
		if sn == p {
			return true
		}

		// Check the remaining string, minus current match
		if isValidPattern(pattern[sn:], pt) {
			return true
		}

	}

	return false

}
