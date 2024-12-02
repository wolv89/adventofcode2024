package day02

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type AocDay2 struct{}

const DIR = "day02/"

func (d AocDay2) Puzzle1(useSample bool) {

	datafile := DIR + "data.txt"
	if useSample {
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
		line                     string
		i, last, next, dif, safe int
		inc, isSafe              bool
	)

	for scanner.Scan() {

		line = scanner.Text()
		isSafe = true

		rawNumbers := strings.Fields(line)
		if len(rawNumbers) < 2 {
			// Doesn't seem to be any cases like this
			continue
		}

		last, err = strconv.Atoi(rawNumbers[0])
		if err != nil {
			fmt.Println("Unable to parse ", rawNumbers[0])
			continue
		}

		next, err = strconv.Atoi(rawNumbers[1])
		if err != nil {
			fmt.Println("Unable to parse ", rawNumbers[1])
			continue
		}

		if next > last {
			inc = true
			dif = next - last
		} else {
			inc = false
			dif = last - next
		}

		if dif < 1 || dif > 3 {
			fmt.Println("Unsafe | ", line, " | Dif not in range: ", dif)
			isSafe = false
			continue
		}

		for i = 2; i < len(rawNumbers); i++ {

			last = next
			next, err = strconv.Atoi(rawNumbers[i])
			if err != nil {
				fmt.Println("Unable to parse ", rawNumbers[0])
				break
			}

			if inc {

				if next < last {
					fmt.Println("Unsafe | ", line, " | Expected increase at: ", last, ">", next)
					isSafe = false
					break
				}

				dif = next - last

				if dif < 1 || dif > 3 {
					fmt.Println("Unsafe | ", line, " | Dif not in range: ", dif)
					isSafe = false
					break
				}

			} else {

				if next > last {
					fmt.Println("Unsafe | ", line, " | Expected decrease at: ", last, ">", next)
					isSafe = false
					break
				}

				dif = last - next

				if dif < 1 || dif > 3 {
					fmt.Println("Unsafe | ", line, " | Dif not in range: ", dif)
					isSafe = false
					break
				}

			}

		}

		if isSafe {
			fmt.Println("SAFE   | ", line)
			safe++
		}

	}

	fmt.Println("------")
	fmt.Printf("TOTAL: %d\n", safe)

}

func (d AocDay2) Puzzle2(useSample bool) {

}
