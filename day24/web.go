package day24

import (
	"bufio"
	"container/list"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"slices"
	"strings"

	"github.com/emirpasic/gods/maps/treemap"
)

/*
 * In my head I was going to have this nice layout that I could see and identify clear problems
 * But in reality its just spaghetti, so I'll have to shift to some other approach
 */
func (d AocDay24) Puzzle2Web(useSample int) {

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
	allRules := make([]Rule, 0)

	gatemap := make([][]string, 0)
	col := make([]string, 0)

	logged := make(map[string]struct{})

	for scanner.Scan() {

		line = scanner.Text()

		if len(line) == 0 {
			gatemap = append(gatemap, col)
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

			// Should probably have stored parts[0] in a var like gatename, etc...
			if parts[0][0] == 'z' {
				zgates++
			} else if parts[0][0] == 'x' || parts[0][0] == 'y' {
				col = append(col, parts[0])
				logged[parts[0]] = struct{}{}
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
			allRules = append(allRules, rule)

		}

	}

	var (
		rl         Rule
		last       *list.Element
		gname      string
		deleteLast bool
	)

	for zgates > 0 {

		col := make([]string, 0)

		for r := rules.Front(); r != nil; r = r.Next() {

			if deleteLast {
				rules.Remove(last)
			}

			deleteLast = false
			last = r

			rl = r.Value.(Rule)

			if _, ok = logged[rl.Inp1]; !ok {
				continue
			}

			if _, ok = logged[rl.Inp2]; !ok {
				continue
			}

			col = append(col, rl.Out)

			if rl.Out[0] == 'z' {
				zgates--
			}

			deleteLast = true

		}

		if len(col) > 0 {
			for _, gname = range col {
				logged[gname] = struct{}{}
			}
			slices.Sort(col)
			gatemap = append(gatemap, col)
		}

		if deleteLast {
			rules.Remove(last)
		}

	}

	type Data struct {
		GateMap [][]string
		Rules   []Rule
	}

	tmpl := template.Must(template.ParseFiles(DIR + "layout.html"))

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir(DIR+"assets/"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		dat := Data{
			gatemap,
			allRules,
		}
		tmpl.Execute(w, dat)
	})

	fmt.Println("")
	fmt.Println("Starting web server:")
	fmt.Println("---")
	fmt.Println("")
	fmt.Println("http://localhost:2627")

	http.ListenAndServe(":2627", nil)

}
