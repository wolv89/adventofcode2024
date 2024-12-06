package day05

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type AocDay5 struct{}

const DIR = "day05/"

func (d AocDay5) Puzzle1(useSample bool) {

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
		rawNums           []string
		line              string
		n, i, total, mid  int
		num, n1, n2, page uint8
		ok, valid         bool
	)

	allRules := make(map[uint8][]uint8)

	pages := make(map[uint8]struct{})
	rules := make(map[uint8][]uint8)

	seen := make(map[uint8]struct{})

	readingRules := true

	for scanner.Scan() {

		line = scanner.Text()

		if len(line) == 0 {
			readingRules = false
			continue
		}

		if readingRules {

			n1, n2 = RuleToInts(line)
			if n1 == 0 || n2 == 0 {
				fmt.Println("Problem parsing numbers from: ", line)
				continue
			}

			allRules[n1] = append(allRules[n1], n2)

		} else {

			rawNums = strings.Split(line, ",")
			clear(pages)
			clear(rules)
			clear(seen)
			list := make([]uint8, 0)
			valid = true

			for n = 0; n < len(rawNums); n++ {
				num = NumToInt(rawNums[n])
				pages[num] = struct{}{}
				list = append(list, num)
			}

			for page = range pages {
				for _, num = range allRules[page] {
					if _, ok = pages[num]; ok {
						rules[page] = append(rules[page], num)
					}
				}
			}

			n = len(list) - 1

		review:
			for i = n; i >= 0; i-- {
				for _, num = range rules[list[i]] {
					if _, ok = seen[num]; !ok {
						valid = false
						break review
					}
				}
				seen[list[i]] = struct{}{}
			}

			if valid {
				mid = int(list[n/2])
				total += mid
				fmt.Println("[x] |", mid, "|", total, " - ", list)
			} else {
				fmt.Println("[ ] |    |", total, " - ", list)
			}

		}

	}

	fmt.Println("")
	fmt.Println("Total: ", total)

}

func (d AocDay5) Puzzle2(useSample bool) {

}

func RuleToInts(rule string) (uint8, uint8) {

	if len(rule) != 5 || rule[2] != '|' {
		return 0, 0
	}

	return NumToInt(string(rule[0:2])), NumToInt(string(rule[3:5]))

}

func NumToInt(num string) uint8 {

	if len(num) != 2 {
		return 0
	}

	return uint8(num[1]-'0') + uint8(num[0]-'0')*10

}
