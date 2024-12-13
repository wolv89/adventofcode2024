package day07

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type AocDay7 struct{}

const DIR = "day07/"

func (d AocDay7) Puzzle1(useSample int) {

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
		parts              []string
		nums               []int64
		line               string
		target, num, total int64
		p                  int
	)

outer:
	for scanner.Scan() {

		line = scanner.Text()

		parts = strings.Split(line, " ")

		nums = make([]int64, 0, len(parts)-1)

		for p = 0; p < len(parts); p++ {

			if p == 0 {
				parts[p] = strings.TrimSuffix(parts[p], ":")
			}

			num, err = strconv.ParseInt(parts[p], 10, 64)

			if err != nil {
				fmt.Println("ERR: Unable to parse ", parts[p], " | ", err)
				continue outer
			}

			if p == 0 {
				target = num
			} else {
				nums = append(nums, num)
			}

		}

		if Calc(nums, nums[0], target, 1) {
			total += target
		}

	}

	fmt.Println("")
	fmt.Println("Total: ", total)

}

/*
 * CREDIT: atrocia6
 * https://www.reddit.com/r/adventofcode/comments/1h8l3z5/comment/m12bjeb/?utm_source=share&utm_medium=web3x&utm_name=web3xcss&utm_term=1&utm_content=share_button
 *
 * My original solution was a bit over-engineered, worked on the test data but not on the full
 * Seems my answer was very close (~1.3 trillion!) but too complex to debug :(
 * Adapted from Atrocia's elegant Python solution
 */
func Calc(nums []int64, sum, target int64, n int) bool {

	if n == len(nums) && sum == target {
		return true
	} else if sum > target || n == len(nums) {
		return false
	}

	return Calc(nums, sum*nums[n], target, n+1) || Calc(nums, sum+nums[n], target, n+1)

}

func (d AocDay7) Puzzle2(useSample int) {

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
		parts              []string
		nums               []int64
		line               string
		target, num, total int64
		p                  int
	)

outer:
	for scanner.Scan() {

		line = scanner.Text()

		parts = strings.Split(line, " ")

		nums = make([]int64, 0, len(parts)-1)

		for p = 0; p < len(parts); p++ {

			if p == 0 {
				parts[p] = strings.TrimSuffix(parts[p], ":")
			}

			num, err = strconv.ParseInt(parts[p], 10, 64)

			if err != nil {
				fmt.Println("ERR: Unable to parse ", parts[p], " | ", err)
				continue outer
			}

			if p == 0 {
				target = num
			} else {
				nums = append(nums, num)
			}

		}

		if CalcTwo(nums, nums[0], target, 1) {
			total += target
		}

	}

	fmt.Println("")
	fmt.Println("Total: ", total)

}

func CalcTwo(nums []int64, sum, target int64, n int) bool {

	if n == len(nums) && sum == target {
		return true
	} else if sum > target || n == len(nums) {
		return false
	}

	return CalcTwo(nums, sum*nums[n], target, n+1) || CalcTwo(nums, IntConcat(sum, nums[n]), target, n+1) || CalcTwo(nums, sum+nums[n], target, n+1)

}

func IntConcat(n1, n2 int64) int64 {

	t := int64(1)

	// Was only doing a < check initially
	// Which meant I was off by a lazy 140 TRILLION in the final answer...
	for t <= n2 {
		t *= 10
	}

	return n1*t + n2

}
