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

func (d AocDay17) Puzzle2(useSample int) {

}
