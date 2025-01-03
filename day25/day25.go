package day25

import (
	"bufio"
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
)

type AocDay25 struct{}

const DIR = "day25/"

const (
	BLOCKHEIGHT = 7
	BLOCKMARKER = "#####"
	TUMBLER     = 5
)

type Pins struct {
	heights []int
	height  int
}

type Lock struct {
	Pins
}

type Key struct {
	Pins
}

func (d AocDay25) Puzzle1(useSample int) {

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
		lock Lock
		key  Key
		line string
		b    int
	)

	block := make([][]byte, BLOCKHEIGHT)

	locks := make([]Lock, 0)
	keys := make([]Key, 0)

	for scanner.Scan() {

		line = scanner.Text()

		if len(line) == 0 {
			continue
		}

		block[b] = []byte(line)
		b++

		if b == BLOCKHEIGHT {
			if string(block[0]) == BLOCKMARKER {
				lock = newLock(block)
				locks = append(locks, lock)
			} else {
				key = newKey(block)
				keys = append(keys, key)
			}

			b = 0
		}

	}

	// Sort the keys so we can cut short the number of checks
	// we're doing against the locks, below
	slices.SortFunc(keys, func(a, b Key) int {
		return cmp.Compare(a.height, b.height)
	})

	var (
		fit, count, checks int
	)

	for _, lock = range locks {

		for _, key = range keys {

			fit = lock.Fits(key)
			checks++

			if fit == TUMBLER {
				count++
			} else if fit == 0 {
				break
			}

			if useSample > 0 {
				if fit == TUMBLER {
					fmt.Println("[X] ", lock.heights, " ", key.heights, " ~ ", count)
				} else {
					fmt.Println("[ ] ", lock.heights, " ", key.heights, " ! ", fit)
				}
			}

		}

		if useSample > 0 {
			fmt.Println("")
		}

	}

	fmt.Println("")
	fmt.Println("Count: ", count, " (", checks, ")")

}

func newLock(block [][]byte) Lock {

	// fmt.Println("New LOCK")
	// fmt.Println("--------")
	// fmt.Println("")

	// for _, row := range block {
	// 	fmt.Println(string(row))
	// }

	// fmt.Println("")

	var i, j, height int
	w := len(block[0])

	heights := make([]int, w)

	for i = 0; i < w; i++ {
		j = 1
		for j < BLOCKHEIGHT {
			if block[j][i] == '.' {
				break
			}
			j++
		}
		heights[i] = j - 1
	}

	m := 1
	for i = w - 1; i >= 0; i-- {
		height += heights[i] * m
		m *= 10
	}

	lock := Lock{
		Pins{
			heights: heights,
			height:  height,
		},
	}

	// fmt.Println(lock)
	// fmt.Println("")

	return lock

}

func newKey(block [][]byte) Key {

	// fmt.Println("New KEY")
	// fmt.Println("--------")
	// fmt.Println("")

	// for _, row := range block {
	// 	fmt.Println(string(row))
	// }

	// fmt.Println("")

	var i, j, height int
	w := len(block[0])

	heights := make([]int, w)

	for i = 0; i < w; i++ {
		j = BLOCKHEIGHT - 1
		for j > 0 {
			if block[j][i] == '.' {
				break
			}
			j--
		}
		heights[i] = BLOCKHEIGHT - j - 2
	}

	m := 1
	for i = w - 1; i >= 0; i-- {
		height += heights[i] * m
		m *= 10
	}

	key := Key{
		Pins{
			heights: heights,
			height:  height,
		},
	}

	// fmt.Println(key)
	// fmt.Println("")

	return key

}

func (l Lock) Fits(k Key) int {

	var x int

	for x = 0; x < len(l.heights); x++ {
		if l.heights[x]+k.heights[x] > TUMBLER {
			break
		}
	}

	return x

}

func (d AocDay25) Puzzle2(useSample int) {

}
