package day00

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type AocDay0 struct{}

const DIR = "day00/"

func (d AocDay0) Puzzle1(useSample bool) {

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
		line string
	)

	for scanner.Scan() {

		line = scanner.Text()

		// Just a template to be copied...
		fmt.Println(line)

	}

}

func (d AocDay0) Puzzle2(useSample bool) {

}
