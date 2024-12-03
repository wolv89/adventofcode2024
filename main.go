package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/wolv89/adventofcode2024/day01"
	"github.com/wolv89/adventofcode2024/day02"
	"github.com/wolv89/adventofcode2024/day03"
)

var (
	flagday, flagpuzzle int
	flagsample          bool
)

type AocDay interface {
	Puzzle1(bool)
	Puzzle2(bool)
}

func init() {

	flag.IntVar(&flagday, "day", 0, "Advent of Code day to run (Between 1 and 25)")
	flag.IntVar(&flagday, "d", 0, "Advent of Code day to run (Between 1 and 25)")

	flag.IntVar(&flagpuzzle, "puzzle", 1, "Which puzzle to run on the given day (1 or 2, defaults to 1)")
	flag.IntVar(&flagpuzzle, "p", 1, "Which puzzle to run on the given day (1 or 2, defaults to 1)")

	flag.BoolVar(&flagsample, "sample", false, "Use sample data, instead of full data? (Defaults to false)")
	flag.BoolVar(&flagsample, "s", false, "Use sample data, instead of full data? (Defaults to false)")

}

func validateFlags() {

	if flagday < 1 || flagday > 25 {
		log.Fatalf("Please enter a day between 1 and 25")
	}

	if flagpuzzle < 1 || flagpuzzle > 2 {
		log.Fatalf("Puzzle step can only be 1 or 2")
	}

}

func main() {

	flag.Parse()
	validateFlags()

	var day AocDay

	switch flagday {
	case 1:
		day = day01.AocDay1{}
	case 2:
		day = day02.AocDay2{}
	case 3:
		day = day03.AocDay3{}
	}

	if flagpuzzle == 1 {
		day.Puzzle1(flagsample)
	} else {
		day.Puzzle2(flagsample)
	}

	fmt.Println("")

}
