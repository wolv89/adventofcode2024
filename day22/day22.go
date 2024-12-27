package day22

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type AocDay22 struct{}

const DIR = "day22/"

const MOD = 16777216

func (d AocDay22) Puzzle1(useSample int) {

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
		line                 string
		start, secret, total int64
		s                    int
	)

	for scanner.Scan() {

		line = scanner.Text()

		start, err = strconv.ParseInt(line, 10, 64)
		if err != nil {
			fmt.Println("Error parsing: ", line, " | ", err)
			continue
		}

		secret = start

		for s = 0; s < 2000; s++ {
			secret = NextSecret(secret)
		}

		fmt.Printf("%d: %d\n", start, secret)

		total += secret

	}

	fmt.Println("")
	fmt.Println("Total: ", total)

}

func (d AocDay22) Puzzle2(useSample int) {

}

func NextSecret(sec int64) int64 {

	m1 := sec * 64
	sec ^= m1
	sec %= MOD

	d1 := sec / 32
	sec ^= d1
	sec %= MOD

	m2 := sec * 2048
	sec ^= m2
	sec %= MOD

	return sec

}
