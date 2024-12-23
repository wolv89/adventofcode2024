package day17

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
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
		true,
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

	const (
		RTNS = 64
		SIZE = 5_000_000
		STEP = 5
	)

	pgs := make([]Program, RTNS+1)
	inst := len(program.Instructions)

	for l = 0; l < RTNS; l++ {
		pgs[l] = Program{
			Output:       make([]int64, 0),
			Instructions: make([]int, inst),
			Register:     [3]int64{0, 0, 0},
			Ptr:          0,
			WithChecks:   true,
		}
		copy(pgs[l].Instructions, program.Instructions)
	}

	var (
		outer, inner, off int64
		wg                sync.WaitGroup
		found             atomic.Bool
		alt               bool
	)

	// Checked up to here or so
	// Giving up, for now... I'll be back
	off = 37_488_000_000_000 // Pow64(8, int64(inst)-1) <-- our start

	for outer = 0; outer < 3600; outer++ {

		alt = !alt

		if alt {
			fmt.Println("+++ OUTER LOOP +++ {", outer, "} ", Int64Format(off))
		}

		for inner = 0; inner < RTNS; inner++ {

			// fmt.Println("Inner Loop: ", inner, " | ", off)

			wg.Add(1)

			go func(idx int64, offset int64) {
				defer wg.Done()
				for reg := offset; reg < offset+SIZE; reg += STEP {
					pgs[inner].Reset(reg)

					// fmt.Println(reg, " -------")

					// fmt.Println(pgs[inner])
					res := pgs[inner].Run()

					// fmt.Println(reg, " | ", res)

					if res {
						fmt.Println("MATCH [", reg, "]")
						found.Store(true)
					}
				}

			}(inner, off)
			off += SIZE

		}

		wg.Wait()

		if found.Load() {
			break
		}

	}

	/*
		var (
			reg int64
			res bool
		)

		for reg = 2_000_000_000; reg < 4_000_000_000; reg++ {

			if reg%1_000_000 == 0 {
				fmt.Println("A Registers: ", reg, " - ", reg+999_999)
			}

			program.Reset(reg)
			res = program.Run()

			if res {
				fmt.Println("MATCH [", reg, "]")
				break
			}

		}
	*/

}

// Credit: https://stackoverflow.com/a/31046325
func Int64Format(n int64) string {
	in := strconv.FormatInt(n, 10)
	numOfDigits := len(in)
	if n < 0 {
		numOfDigits-- // First character is the - sign (not a digit)
	}
	numOfCommas := (numOfDigits - 1) / 3

	out := make([]byte, len(in)+numOfCommas)
	if n < 0 {
		in, out[0] = in[1:], '-'
	}

	for i, j, k := len(in)-1, len(out)-1, 0; ; i, j = i-1, j-1 {
		out[j] = in[i]
		if i == 0 {
			return string(out)
		}
		if k++; k == 3 {
			j, k = j-1, 0
			out[j] = ','
		}
	}
}
