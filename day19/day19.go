package day19

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/wolv89/adventofcode2024/structures"
)

type AocDay19 struct{}

type AsyncCounter struct {
	count int64
	mu    sync.RWMutex
}

type Cache struct {
	saved map[string]int
	mu    sync.RWMutex
}

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
		rawPatterns []string
		line        string
		wg          sync.WaitGroup
	)

	total := AsyncCounter{
		0,
		sync.RWMutex{},
	}

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

	allPatterns := make([]string, 0)

	for scanner.Scan() {

		line = scanner.Text()

		allPatterns = append(allPatterns, line)

	}

	const THREADS = 32
	var start int

	ch := &Cache{
		make(map[string]int),
		sync.RWMutex{},
	}

	// Honestly the cache is doing all the work
	// The Goroutines might not be helping much
	// Could even be extra overhead?!
	for {

		mx := min(THREADS, len(allPatterns[start:]))

		for _, pat := range allPatterns[start : start+mx] {

			wg.Add(1)

			go func(l string) {
				defer wg.Done()

				count := countValidPatterns(l, pt, ch)

				if count > 0 {
					fmt.Printf("[%5d] %s\n", count, l)
					total.mu.Lock()
					total.count += int64(count)
					total.mu.Unlock()
				} else {
					fmt.Printf("[x] %s\n", l)
				}
			}(pat)

		}

		wg.Wait()

		start += THREADS
		if start >= len(allPatterns) {
			break
		}

	}

	fmt.Println("")
	fmt.Println("Total: ", total.count)

}

func isValidPattern(pattern string, pt structures.Trie) bool {

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

func countValidPatterns(pattern string, pt structures.Trie, ch *Cache) int {

	p := len(pattern)

	m := min(pt.Longest(), p)

	lookup := ch.Get(pattern)
	if lookup >= 0 {
		return lookup
	}

	matches := pt.FindSubmatches(pattern[:m])

	n := len(matches)

	if n == 0 {
		return 0
	}

	var sn, count int

	for i := n - 1; i >= 0; i-- {

		sn = len(matches[i])

		// Matched the full remaining part
		if sn == p {
			count++
			continue
		}

		// Check the remaining string, minus current match
		count += countValidPatterns(pattern[sn:], pt, ch)

	}

	ch.Save(pattern, count)

	return count

}

func (c *Cache) Get(s string) int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if _, ok := c.saved[s]; ok {
		return c.saved[s]
	}
	return -1
}

func (c *Cache) Save(s string, cn int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.saved[s] = cn
}
