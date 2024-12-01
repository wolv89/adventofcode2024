package day01

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type AocDay1 struct{}

const DIR = "day01/"

func (d AocDay1) Puzzle1(useSample bool) {

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

	var line string

	for scanner.Scan() {

		line = scanner.Text()

		fmt.Println(line)

	}

}

func (d AocDay1) Puzzle2(useSample bool) {

	fmt.Printf("Running puzzle TWO - %v", useSample)

}
