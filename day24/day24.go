package day24

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"math/bits"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/emirpasic/gods/maps/treemap"
)

type AocDay24 struct{}

const DIR = "day24/"

// Evidently I have used the wrong names for these
type Gate struct {
	Val, Set bool
}

type Rule struct {
	Inp1 string `json:"inp1"`
	Inp2 string `json:"inp2"`
	Out  string `json:"out"`
	Op   byte   `json:"op"`
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
				Inp1: parts[0],
				Inp2: parts[2],
				Out:  parts[4],
				Op:   op,
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
			fmt.Println(rl.Inp1, string(rl.Op), rl.Inp2, " -> ", rl.Out)
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

			g1i, _ = gates.Get(rl.Inp1)
			g2i, _ = gates.Get(rl.Inp2)

			// Puzzled about this separation
			// "Get"ing interface, then casting to gate
			// Can't seem to do in one line...?
			// NB - I get it now, multiple return values, can't type-assert all of them
			g1, g2 = g1i.(Gate), g2i.(Gate)

			if !g1.Set || !g2.Set {
				continue
			}

			goi, _ = gates.Get(rl.Out)
			gout = goi.(Gate)

			switch rl.Op {
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
			gates.Put(rl.Out, gout)

			if rl.Out[0] == 'z' {
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

type Adder struct {
	gates     []Rule
	registers treemap.Map
}

/*
 * CREDIT: LxsterGames
 * https://www.reddit.com/r/adventofcode/comments/1hla5ql/2024_day_24_part_2_a_guide_on_the_idea_behind_the/
 *
 * Thanks to this great write-up and Kotlin solution, roughly adapted to Go below...
 */
func (d AocDay24) Puzzle2(useSample int) {

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
		parts        []string
		line         string
		op           byte
		readingRules bool
	)

	adder := Adder{
		make([]Rule, 0),
		*treemap.NewWithStringComparator(),
	}

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

			adder.registers.Put(parts[0], parts[1] == "1")

		} else {
			// Read rules

			parts = strings.Split(line, " ")
			if len(parts) < 5 {
				fmt.Println("Error parsing line: ", line)
				continue
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
				Inp1: parts[0],
				Inp2: parts[2],
				Out:  parts[4],
				Op:   op,
			}

			// Oh boy...
			adder.gates = append(adder.gates, rule)

		}

	}

	swappedGates := adder.swapGates()

	xval := adder.getWiresVal('x')
	yval := adder.getWiresVal('y')

	zval := adder.Add()

	res := (xval + yval) ^ zval

	fmt.Println("X:   ", xval)
	fmt.Println("Y: + ", yval)
	fmt.Println("Z: = ", zval)
	fmt.Println("---")
	fmt.Println("xor ", res)
	fmt.Printf("%b\n", res)
	fmt.Println("---")

	tail := fmt.Sprintf("%02d", bits.TrailingZeros64(uint64(res)))
	extraGates := adder.findGatesEndingWith(tail)

	var g int

	finalNames := make([]string, 0, len(swappedGates)+len(extraGates))

	for _, g = range swappedGates {
		finalNames = append(finalNames, adder.gates[g].Out)
	}

	for _, g = range extraGates {
		finalNames = append(finalNames, adder.gates[g].Out)
	}

	slices.Sort(finalNames)

	fmt.Println(strings.Join(finalNames, ","))

}

func (a *Adder) swapGates() []int {

	filter1, filter2 := make([]int, 0), make([]int, 0)

	var (
		gate             Rule
		g, h, swap, comp int
	)

	for g, gate = range a.gates {

		if gate.Out[0] == 'z' && gate.Out != "z45" && gate.Op != OP_XOR {
			filter1 = append(filter1, g)
		}

		if (gate.Inp1[0] != 'x' && gate.Inp1[0] != 'y') && (gate.Inp2[0] != 'x' && gate.Inp2[0] != 'y') && gate.Out[0] != 'z' && gate.Op == OP_XOR {
			filter2 = append(filter2, g)
		}

	}

	for _, g = range filter2 {
		comp = a.findZGateUsing(g)
		for _, h = range filter1 {
			if h == comp {
				swap = h
			}
		}
		a.gates[g].Out, a.gates[swap].Out = a.gates[swap].Out, a.gates[g].Out
	}

	filter1 = append(filter1, filter2...)
	return filter1

}

func (a Adder) findGatesEndingWith(tail string) []int {

	found := make([]int, 0)

	for h, gate := range a.gates {
		if strings.HasSuffix(gate.Inp1, tail) && strings.HasSuffix(gate.Inp2, tail) {
			found = append(found, h)
		}
	}

	return found

}

func (a Adder) findZGateUsing(g int) int {

	filter := make([]int, 0)

	c := a.gates[g].Out

	for h, gate := range a.gates {
		if gate.Inp1 == c || gate.Inp2 == c {
			filter = append(filter, h)
		}
	}

	if len(filter) == 0 {
		return -1
	}

	for _, fl := range filter {
		if a.gates[fl].Out[0] == 'z' {
			return a.findZGatePrevious(a.gates[fl].Out)
		}
	}

	return a.findZGateUsing(filter[0])

}

func (a Adder) findZGatePrevious(name string) int {

	x, err := strconv.Atoi(name[1:])
	if err != nil {
		log.Fatal("Uh oh: ", err)
	}

	search := fmt.Sprintf("z%02d", x-1)

	for g := range a.gates {
		if a.gates[g].Out == search {
			return g
		}
	}

	return -1

}

func (a Adder) getWiresVal(c byte) int64 {

	var b strings.Builder

	a.registers.Each(func(key, value interface{}) {

		name := key.(string)
		val := value.(bool)

		if name[0] == c {
			if val {
				b.WriteByte('1')
			} else {
				b.WriteByte('0')
			}
		}

	})

	num := b.String()

	b.Reset()
	b.Grow(len(num))

	for i := len(num) - 1; i >= 0; i-- {
		b.WriteByte(num[i])
	}

	num = b.String()

	x, err := strconv.ParseInt(num, 2, 64)
	if err != nil {
		log.Fatal("Oh sh*t: ", err)
	}

	return x

}

func (a *Adder) Add() int64 {

	rules := list.New()

	var (
		rl                           Rule
		g1i, g2i                     interface{}
		last                         *list.Element
		zgates                       int
		deleteLast, ok, g1, g2, gout bool
	)

	for _, rl = range a.gates {
		rules.PushBack(rl)
		if rl.Out[0] == 'z' {
			zgates++
		}
	}

	for zgates > 0 {

		for r := rules.Front(); r != nil; r = r.Next() {

			if deleteLast {
				rules.Remove(last)
			}

			deleteLast = false
			last = r

			rl = r.Value.(Rule)

			if _, ok = a.registers.Get(rl.Inp1); !ok {
				continue
			}

			if _, ok = a.registers.Get(rl.Inp2); !ok {
				continue
			}

			g1i, _ = a.registers.Get(rl.Inp1)
			g2i, _ = a.registers.Get(rl.Inp2)

			g1, g2 = g1i.(bool), g2i.(bool)

			switch rl.Op {
			case OP_AND:
				gout = g1 && g2
			case OP_OR:
				gout = g1 || g2
			case OP_XOR:
				if (g1 || g2) && (g1 != g2) {
					gout = true
				} else {
					gout = false
				}
			}

			a.registers.Put(rl.Out, gout)

			if rl.Out[0] == 'z' {
				zgates--
			}

			deleteLast = true

		}

		if deleteLast {
			rules.Remove(last)
		}

	}

	return a.getWiresVal('z')

}
