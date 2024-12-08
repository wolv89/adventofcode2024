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

var OPS = [2]byte{
	'*',
	'+',
}

func (d AocDay7) Puzzle1(useSample bool) {

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
		parts                   []string
		nums                    []int64
		line                    string
		target, num, sum, total int64
		p, lim                  int
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

		sum = Calc(nums, nums[0], target, 1)

		if sum == target {
			// fmt.Println("")
			fmt.Println(lim, " | VALID - ", sum)
			total += sum
		} else {
			fmt.Println(lim, " | Invalid - ", target)
		}

		fmt.Println("--------------")

		lim++
		// if lim > 10 {
		// 	break
		// }

	}

	// fmt.Println("")
	// for v, vld := range valids {
	// 	fmt.Println(v, " > ", vld)
	// }
	fmt.Println("")
	fmt.Println("Total: ", total)

}

func Calc(nums []int64, sum, target int64, n int) int64 {

	if n >= len(nums) {
		return sum
	}

	s := sum

	// fmt.Println(n, " | ", sum, " | Multiply ", nums[n])

	sum *= nums[n]
	if sum == target {
		return sum
	} else if sum < target {
		sum = Calc(nums, sum, target, n+1)
		if sum == target {
			return sum
		}
	}

	sum = s
	// fmt.Println(n, " | ", sum, " | Add ", nums[n])

	sum += nums[n]
	if sum == target {
		return sum
	} else if sum > target {
		return s
	}

	return Calc(nums, sum, target, n+1)

}

func (d AocDay7) Puzzle2(useSample bool) {

}
