package day24

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/emirpasic/gods/maps/treemap"
)

type AocDay24 struct{}

const DIR = "day24/"

type Gate struct {
	Val, Set bool
}

type Rule struct {
	inp1, inp2, out string
	op              byte
}

const (
	OP_AND = '&'
	OP_OR  = '|'
	OP_XOR = '^'
)

func (d AocDay24) Puzzle1(useSample int) {

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
		parts            []string
		line             string
		zgates           int
		op               byte
		readingRules, ok bool
	)

	gates := treemap.NewWithStringComparator()
	rules := list.New()

	for scanner.Scan() {

		line = scanner.Text()

		if len(line) == 0 {
			readingRules = true
			continue
		}

		if !readingRules {
			// Read gates/initial values

			parts = strings.Split(line, ": ")
			if len(parts) < 2 {
				fmt.Println("Error parsing line: ", line)
				continue
			}

			gate := Gate{parts[1] == "1", true}
			gates.Put(parts[0], gate)

			if parts[0][0] == 'z' {
				zgates++
			}

		} else {
			// Read rules

			parts = strings.Split(line, " ")
			if len(parts) < 5 {
				fmt.Println("Error parsing line: ", line)
				continue
			}

			// Add all gates mentioned in this rule to the gate map, with no initial value
			if _, ok = gates.Get(parts[0]); !ok {
				gates.Put(parts[0], Gate{false, false})
				if parts[0][0] == 'z' {
					zgates++
				}
			}

			if _, ok = gates.Get(parts[2]); !ok {
				gates.Put(parts[2], Gate{false, false})
				if parts[2][0] == 'z' {
					zgates++
				}
			}

			if _, ok = gates.Get(parts[4]); !ok {
				gates.Put(parts[4], Gate{false, false})
				if parts[4][0] == 'z' {
					zgates++
				}
			}

			switch parts[1] {
			case "AND":
				op = OP_AND
			case "OR":
				op = OP_OR
			case "XOR":
				op = OP_XOR
			}

			// Create rule
			rule := Rule{
				inp1: parts[0],
				inp2: parts[2],
				out:  parts[4],
				op:   op,
			}

			rules.PushBack(rule)

		}

	}

	// INITIAL STATE
	if useSample > 0 {

		fmt.Println("GATES")
		fmt.Println("-----")
		gates.Each(func(key, value interface{}) {
			fmt.Println(key, ": ", value)
		})

		fmt.Println("")

		fmt.Println("RULES")
		fmt.Println("-----")
		for r := rules.Front(); r != nil; r = r.Next() {
			rl := r.Value.(Rule)
			fmt.Println(rl.inp1, string(rl.op), rl.inp2, " -> ", rl.out)
		}

		fmt.Println("")
		fmt.Println("Z Gates: ", zgates)

	}

	bin := make([]byte, zgates)
	z := zgates

	var (
		rl            Rule
		g1i, g2i, goi interface{}
		g1, g2, gout  Gate
		last          *list.Element
		deleteLast    bool
	)

	for zgates > 0 {

		for r := rules.Front(); r != nil; r = r.Next() {

			if deleteLast {
				rules.Remove(last)
			}

			deleteLast = false
			last = r

			rl = r.Value.(Rule)

			g1i, _ = gates.Get(rl.inp1)
			g2i, _ = gates.Get(rl.inp2)

			// Puzzled about this separation
			// "Get"ing interface, then casting to gate
			// Can't seem to do in one line...?
			g1, g2 = g1i.(Gate), g2i.(Gate)

			if !g1.Set || !g2.Set {
				continue
			}

			goi, _ = gates.Get(rl.out)
			gout = goi.(Gate)

			switch rl.op {
			case OP_AND:
				gout.Val = g1.Val && g2.Val
			case OP_OR:
				gout.Val = g1.Val || g2.Val
			case OP_XOR:
				if (g1.Val || g2.Val) && (g1.Val != g2.Val) {
					gout.Val = true
				} else {
					gout.Val = false
				}
			}

			gout.Set = true
			gates.Put(rl.out, gout)

			if rl.out[0] == 'z' {
				zgates--
			}

			deleteLast = true

		}

		if deleteLast {
			rules.Remove(last)
		}

	}

	// ACTIONED STATE
	if useSample > 0 {

		fmt.Println("")
		fmt.Println("~")
		fmt.Println("")

	}

	var (
		gname string
	)

	fmt.Println("RESULT")
	fmt.Println("------")
	gates.Each(func(key, value interface{}) {

		gname = key.(string)
		g1 = value.(Gate)

		if useSample < 1 {
			if gname[0] != 'z' {
				return
			}
		}

		op = ' '
		if g1.Set {
			if g1.Val {
				op = '1'
			} else {
				op = '0'
			}
		}

		if gname[0] == 'z' {
			z--
			bin[z] = op
		}

		fmt.Println(key, ":", string(op))
	})

	fmt.Println("")

	res, cerr := strconv.ParseInt(string(bin), 2, 64)
	if cerr != nil {
		fmt.Println("Error parsing binary string: ", cerr)
	}

	fmt.Println("Raw: ", string(bin))
	fmt.Println("Result: ", res)

}

func (d AocDay24) Puzzle2(useSample int) {

}
