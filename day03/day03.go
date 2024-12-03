package day03

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
)

type AocDay3 struct{}

const DIR = "day03/"
const READ = 16

func (d AocDay3) Puzzle1(useSample bool) {

	datafile := DIR + "data.txt"
	if useSample {
		datafile = DIR + "sample.txt"
	}

	f, err := os.Open(datafile)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	var (
		search                     [][]string
		buf                        []byte
		str                        string
		n, b, r, num1, num2, total int
	)

	regx := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)

	for {

		buf = make([]byte, READ)
		n, err = io.ReadAtLeast(f, buf, READ)

		if n == 0 {
			break
		}

		if err == io.ErrUnexpectedEOF {
			f.Seek(int64(-n), io.SeekCurrent)
			n, _ = io.ReadAtLeast(f, buf, n)
		}

		str = string(buf)

		// If we read the full compliment of bytes (not end of file, as above), then
		// walk back the last 12 characters (if there are 12) and look for a closing bracket.
		// We don't want to have a split that breaks a mul(###,###) - 12 chars
		// (Although this doesn't exactly protect against something like mul(###,###))))
		if n == READ {
			r = min(0, n-12)
			for b = n - 1; b > r; b-- {
				if buf[b] == ')' {
					b++
					break
				}
			}
			if b > r {
				str = string(buf[:b])
				f.Seek(int64(b-n), io.SeekCurrent)
			}
		}

		search = regx.FindAllStringSubmatch(str, -1)

		if len(search) > 0 {
			for _, found := range search {

				num1, err = strconv.Atoi(found[1])
				if err != nil {
					fmt.Println(err)
					continue
				}

				num2, err = strconv.Atoi(found[2])
				if err != nil {
					fmt.Println(err)
					continue
				}

				fmt.Println("Multiplying ", num1, " * ", num2, " | ", total)
				total += num1 * num2

			}
		}

	}

	fmt.Println("TOTAL: ", total)
	fmt.Println("")

}

func (d AocDay3) Puzzle2(useSample bool) {

	datafile := DIR + "data.txt"
	if useSample {
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
		search, nums      [][]string
		line              string
		num1, num2, total int
	)

	do := true

	regx := regexp.MustCompile(`(mul|do|don\'t)\(([0-9,]{3,7})?\)`)
	numx := regexp.MustCompile(`([0-9]{1,3}),([0-9]{1,3})`)

	// Reverted to reading in full lines as I could not think of a clean method
	// to read in buffered chunks like in Puzzle1(), and not break some of the input
	// Although as I type this I wonder if I could have a slice of "special" characters
	// and iterate back from the end of each buffered read to make sure it doesn't end
	// with any of those characters, so we know we're splitting on a "dead" character
	for scanner.Scan() {

		line = scanner.Text()

		search = regx.FindAllStringSubmatch(line, -1)

		if len(search) > 0 {
			for _, found := range search {

				if !do {
					if found[1] == "do" {
						fmt.Println("# Enabling...")
						do = true
					}
					continue
				}

				if found[1] == "mul" {

					nums = numx.FindAllStringSubmatch(found[2], -1)

					for _, nms := range nums {

						num1, err = strconv.Atoi(nms[1])
						if err != nil {
							fmt.Println(err)
							continue
						}

						num2, err = strconv.Atoi(nms[2])
						if err != nil {
							fmt.Println(err)
							continue
						}

						fmt.Println("Multiplying ", num1, " * ", num2, " | ", total)
						total += num1 * num2

					}

				} else if found[1] == "don't" {
					fmt.Println("Disabling...")
					do = false
				}

			}
		}

	}

	fmt.Println("TOTAL: ", total)
	fmt.Println("")

}
