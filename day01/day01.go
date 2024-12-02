package day01

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/wolv89/adventofcode2024/structures"
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

	var (
		line                   string
		num1, num2, dif, total int
	)

	list1, list2 := &structures.MinHeap{}, &structures.MinHeap{}
	heap.Init(list1)
	heap.Init(list2)

	for scanner.Scan() {

		line = scanner.Text()

		rawNumbers := strings.Fields(line)
		if len(rawNumbers) < 2 {
			continue // Log ?
		}

		num1, err = strconv.Atoi(rawNumbers[0])
		if err != nil {
			continue // Log ?
		} else {
			heap.Push(list1, num1)
		}

		num2, err = strconv.Atoi(rawNumbers[1])
		if err != nil {
			continue // Log ?
		} else {
			heap.Push(list2, num2)
		}

	}

	if list1.Len() != list2.Len() {
		fmt.Println("Something has gone wrong here...")
		return
	}

	for list1.Len() > 0 {

		num1, num2 = heap.Pop(list1).(int), heap.Pop(list2).(int)

		dif = structures.AbsInt(num1 - num2)
		total += dif

		fmt.Printf("%d - %d | %d | %d\n", num1, num2, dif, total)

	}

	fmt.Println("------")
	fmt.Printf("TOTAL: %d\n", total)

}

func (d AocDay1) Puzzle2(useSample bool) {

	fmt.Printf("Running puzzle TWO - %v", useSample)

}
