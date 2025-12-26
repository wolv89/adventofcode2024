package day17

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type AocDay17 struct{}

const DIR = "day17/"

func (d AocDay17) Puzzle1(useSample int) {

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
	scanner.Split(bufio.ScanLines)

	var (
		line          string
		l, num        int
		readRegisters bool
	)

	program := Program{
		make([]int64, 0),
		make([]int, 0),
		[3]int64{0, 0, 0},
		0,
		false,
	}

	for scanner.Scan() {

		line = scanner.Text()

		// "Read" like "Red", as in past tense...
		if len(line) == 0 {
			readRegisters = true
			continue
		}

		if !readRegisters {

			parts := strings.Split(line, ": ")
			if len(parts) < 2 {
				fmt.Println("Unable to parse register line: ", parts)
				continue
			}

			num, err = strconv.Atoi(parts[1])

			if err != nil {
				fmt.Println("Unable to parse register value: ", err)
				continue
			}

			program.Register[l] = int64(num)
			l++

		} else {

			parts := strings.Split(line, ": ")
			if len(parts) < 2 {
				fmt.Println("Unable to parse program line: ", parts)
				continue
			}

			for l = 0; l < len(parts[1]); l++ {
				if parts[1][l] == ',' {
					continue
				}
				program.Instructions = append(program.Instructions, int(parts[1][l]-'0'))
			}

		}

	}

	fmt.Println("")

	program.Run()

	fmt.Println(program.Render())

}

/*
 * Finally got this done, with some help from Gemini, suggesting the backtracking approach
 */
func (d AocDay17) Puzzle2(useSample int) {

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

	target := make([]uint64, 0)

	for scanner.Scan() {

		line := scanner.Text()
		if len(line) < 10 {
			continue
		}

		if line[:7] != "Program" {
			continue
		}

		for i := 9; i < len(line); i++ {
			if line[i] >= '0' && line[i] <= '9' {
				target = append(target, uint64(line[i]-'0'))
			}
		}

	}

	fmt.Println("Target:", target)

	var backtrack func(uint64, int) bool
	backtrack = func(A uint64, ind int) bool {

		fmt.Println(A, ind)

		if ind < 0 {
			fmt.Println("Result:", A)
			return true
		}

		for i := uint64(0); i < 8; i++ {

			next := (A << 3) | i

			B := (next % 8) ^ 5
			C := next >> B
			B = B ^ C ^ 6

			if B%8 == target[ind] {
				if backtrack(next, ind-1) {
					return true
				}
			}

		}

		return false

	}

	backtrack(0, len(target)-1)

}
