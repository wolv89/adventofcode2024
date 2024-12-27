package day23

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type AocDay23 struct{}

const DIR = "day23/"

func (d AocDay23) Puzzle1(useSample int) {

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
		line, c1, c2                  string
		index, i1, i2, x, y, z, count int
		valid                         byte
		ok                            bool
	)

	comps := make(map[string]int)
	indexedComps := make([]string, 0)
	connections := make([]string, 0)

	for scanner.Scan() {

		line = scanner.Text()

		// All lines in the exact format: ab-xy
		c1, c2 = line[0:2], line[3:5]

		if _, ok = comps[c1]; !ok {
			comps[c1] = index
			indexedComps = append(indexedComps, c1)
			index++
		}

		if _, ok = comps[c2]; !ok {
			comps[c2] = index
			indexedComps = append(indexedComps, c2)
			index++
		}

		connections = append(connections, line)

	}

	network := make([][]bool, index)
	for n := range network {
		network[n] = make([]bool, index)
	}

	for _, line = range connections {

		c1, c2 = line[0:2], line[3:5]
		i1, i2 = comps[c1], comps[c2]

		network[i1][i2] = true
		network[i2][i1] = true

	}

	if useSample > 0 {
		fmt.Println("")

		fmt.Print("    ")
		for _, ic := range indexedComps {
			fmt.Printf("%s  ", ic)
		}
		fmt.Print("\n")

		for x, ic := range indexedComps {
			fmt.Print(ic, "  ")
			for y := range indexedComps {
				if network[x][y] {
					fmt.Print("X   ")
				} else {
					fmt.Print(".   ")
				}
			}
			fmt.Print("\n")
		}

		fmt.Println("")
	}

	for x = 0; x < index; x++ {

		for y = x + 1; y < index; y++ {

			if !network[x][y] {
				continue
			}

			for z = y + 1; z < index; z++ {

				if !network[x][z] || !network[y][z] {
					continue
				}

				valid = ' '

				if indexedComps[x][0] == 't' || indexedComps[y][0] == 't' || indexedComps[z][0] == 't' {
					valid = 'X'
					count++
				}

				fmt.Printf("[%s] %s,%s,%s\n", string(valid), indexedComps[x], indexedComps[y], indexedComps[z])

			}

		}

	}

	fmt.Println("")
	fmt.Println("Count: ", count)

}

type WAN struct {
	network [][]bool
	comps   []string
	lim     int
	lookup  map[string]int
	largest string
}

func (d AocDay23) Puzzle2(useSample int) {

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
		line, c1, c2  string
		index, i1, i2 int
		ok            bool
	)

	comps := make(map[string]int)
	indexedComps := make([]string, 0)
	connections := make([]string, 0)

	for scanner.Scan() {

		line = scanner.Text()

		// All lines in the exact format: ab-xy
		c1, c2 = line[0:2], line[3:5]

		if _, ok = comps[c1]; !ok {
			comps[c1] = index
			indexedComps = append(indexedComps, c1)
			index++
		}

		if _, ok = comps[c2]; !ok {
			comps[c2] = index
			indexedComps = append(indexedComps, c2)
			index++
		}

		connections = append(connections, line)

	}

	network := make([][]bool, index)
	for n := range network {
		network[n] = make([]bool, index)
	}

	for _, line = range connections {

		c1, c2 = line[0:2], line[3:5]
		i1, i2 = comps[c1], comps[c2]

		network[i1][i2] = true
		network[i2][i1] = true

	}

	if useSample > 0 {
		fmt.Println("")

		fmt.Print("    ")
		for _, ic := range indexedComps {
			fmt.Printf("%s  ", ic)
		}
		fmt.Print("\n")

		for x, ic := range indexedComps {
			fmt.Print(ic, "  ")
			for y := range indexedComps {
				if network[x][y] {
					fmt.Print("X   ")
				} else {
					fmt.Print(".   ")
				}
			}
			fmt.Print("\n")
		}

		fmt.Println("")
		fmt.Println(indexedComps)
		fmt.Println("")

	}

	wan := WAN{
		network,
		indexedComps,
		index,
		comps,
		"",
	}

	wan.Map("", 0)

	fmt.Println("")
	fmt.Println(wan.largest)

	fmt.Println("")
	fmt.Println("Password: ", wan.GetPassword())

}

func (w *WAN) Map(network string, n int) {

	if n >= w.lim {
		if len(network) > 2 {
			// w.found = append(w.found, network)
			if len(network) > len(w.largest) {
				w.largest = network
			}
		}
		return
	}

	var (
		x int
	)

	valid := true

	for i := 0; i < len(network); i += 2 {

		x = w.lookup[network[i:i+2]]

		if !w.network[n][x] {
			valid = false
			break
		}

	}

	if valid {
		w.Map(network+w.comps[n], n+1)
	}

	w.Map(network, n+1)

}

func (w WAN) GetPassword() string {

	parts := make([]string, 0, len(w.largest)/2)

	for i := 0; i < len(w.largest); i += 2 {
		parts = append(parts, w.largest[i:i+2])
	}

	slices.Sort(parts)

	return strings.Join(parts, ",")

}
