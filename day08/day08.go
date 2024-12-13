package day08

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type AocDay8 struct{}

const DIR = "day08/"

// Lets try to get it right this time
// x - Horizontal, y - Vertical
type Point struct {
	x, y int
}

func (d AocDay8) Puzzle1(useSample int) {

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
		ants                            []Point
		p1, p2, node                    Point
		line                            string
		w, h, l, a, b, n, dx, dy, count int
		ok                              bool
	)

	antennas := make(map[byte][]Point)
	grid := make([][]byte, 0)

	for scanner.Scan() {

		line = scanner.Text()
		grid = append(grid, []byte(line))

		if w == 0 {
			w = len(line)
		}

		for l = 0; l < len(line); l++ {

			if line[l] == '.' {
				continue
			}

			if _, ok = antennas[line[l]]; !ok {
				// Does it default to a cap of 1 anyway, or higher? Hmm...
				antennas[line[l]] = make([]Point, 0, 1)
			}

			antennas[line[l]] = append(antennas[line[l]], Point{l, h})

		}

		h++

	}

	placed := make([][]bool, h)
	for l = range placed {
		placed[l] = make([]bool, w)
	}

	for _, ants = range antennas {

		n = len(ants)

		for a = 0; a < n-1; a++ {

			p1 = ants[a]

			for b = a + 1; b < n; b++ {

				p2 = ants[b]

				dx, dy = p1.x-p2.x, p1.y-p2.y

				node = Point{p1.x + dx, p1.y + dy}

				if inBounds(node, w, h) {
					if grid[node.y][node.x] == '.' {
						grid[node.y][node.x] = '#'
					}
					if !placed[node.y][node.x] {
						placed[node.y][node.x] = true
						count++
					}
				}

				node = Point{p2.x - dx, p2.y - dy}

				if inBounds(node, w, h) {
					if grid[node.y][node.x] == '.' {
						grid[node.y][node.x] = '#'
					}
					if !placed[node.y][node.x] {
						placed[node.y][node.x] = true
						count++
					}
				}

			}
		}

	}

	for _, row := range grid {
		fmt.Println(string(row))
	}

	fmt.Println("")
	fmt.Println("Count: ", count)

}

func (d AocDay8) Puzzle2(useSample int) {

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
		ants                            []Point
		p1, p2, node                    Point
		line                            string
		w, h, l, a, b, n, dx, dy, count int
		ok                              bool
	)

	antennas := make(map[byte][]Point)
	grid := make([][]byte, 0)

	for scanner.Scan() {

		line = scanner.Text()
		grid = append(grid, []byte(line))

		if w == 0 {
			w = len(line)
		}

		for l = 0; l < len(line); l++ {

			if line[l] == '.' {
				continue
			}

			if _, ok = antennas[line[l]]; !ok {
				// Does it default to a cap of 1 anyway, or higher? Hmm...
				antennas[line[l]] = make([]Point, 0, 1)
			}

			antennas[line[l]] = append(antennas[line[l]], Point{l, h})

		}

		h++

	}

	placed := make([][]bool, h)
	for l = range placed {
		placed[l] = make([]bool, w)
	}

	for _, ants = range antennas {

		n = len(ants)
		if n < 2 {
			continue
		}

		for a = 0; a < n-1; a++ {

			p1 = ants[a]

			if !placed[p1.y][p1.x] {
				placed[p1.y][p1.x] = true
				count++
			}

			for b = a + 1; b < n; b++ {

				p2 = ants[b]

				if !placed[p2.y][p2.x] {
					placed[p2.y][p2.x] = true
					count++
				}

				dx, dy = p1.x-p2.x, p1.y-p2.y

				node = Point{p1.x, p1.y}
				for {

					node.x += dx
					node.y += dy

					if inBounds(node, w, h) {
						if grid[node.y][node.x] == '.' {
							grid[node.y][node.x] = '#'
						}
						if !placed[node.y][node.x] {
							placed[node.y][node.x] = true
							count++
						}
					} else {
						break
					}

				}

				node = Point{p1.x, p1.y}
				for {

					node.x -= dx
					node.y -= dy

					if inBounds(node, w, h) {
						if grid[node.y][node.x] == '.' {
							grid[node.y][node.x] = '#'
						}
						if !placed[node.y][node.x] {
							placed[node.y][node.x] = true
							count++
						}
					} else {
						break
					}

				}

			}
		}

	}

	for _, row := range grid {
		fmt.Println(string(row))
	}

	fmt.Println("")
	fmt.Println("Count: ", count)

}

func inBounds(p Point, w, h int) bool {
	if p.x < 0 || p.x >= w || p.y < 0 || p.y >= h {
		return false
	}
	return true
}
