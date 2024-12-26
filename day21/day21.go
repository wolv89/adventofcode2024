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

}
