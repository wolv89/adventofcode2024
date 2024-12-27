package day22

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type AocDay22 struct{}

const DIR = "day22/"

const MOD = 16777216

func (d AocDay22) Puzzle1(useSample int) {

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
		line                 string
		start, secret, total int64
		s                    int
	)

	for scanner.Scan() {

		line = scanner.Text()

		start, err = strconv.ParseInt(line, 10, 64)
		if err != nil {
			fmt.Println("Error parsing: ", line, " | ", err)
			continue
		}

		secret = start

		for s = 0; s < 2000; s++ {
			secret = NextSecret(secret)
		}

		fmt.Printf("%d: %d\n", start, secret)

		total += secret

	}

	fmt.Println("")
	fmt.Println("Total: ", total)

}

func NextSecret(sec int64) int64 {

	m1 := sec * 64
	sec ^= m1
	sec %= MOD

	d1 := sec / 32
	sec ^= d1
	sec %= MOD

	m2 := sec * 2048
	sec ^= m2
	sec %= MOD

	return sec

}

/*** ------------------------------ PART 2 ------------------------------ ***/

type Cycle struct {
	a, b, c, d int8
}

type Product struct {
	secret        int64
	price, change int8
}

type Seller struct {
	sequence   []Product
	bestPrices map[Cycle]int8
}

type Market struct {
	sellers []Seller
	cycles  map[Cycle]struct{}
}

func (d AocDay22) Puzzle2(useSample int) {

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
		line  string
		start int64
	)

	market := Market{
		make([]Seller, 0),
		make(map[Cycle]struct{}),
	}

	for scanner.Scan() {

		line = scanner.Text()

		start, err = strconv.ParseInt(line, 10, 64)
		if err != nil {
			fmt.Println("Error parsing: ", line, " | ", err)
			continue
		}

		market.AddSeller(start)

	}

	// Works, but takes 4 or 5 seconds to run on full data
	fmt.Println("Best: ", market.GetBestPrice())

}

func (m *Market) AddSeller(secret int64) {

	slr := Seller{
		make([]Product, 0),
		make(map[Cycle]int8),
	}

	pr := Product{
		secret,
		int8(secret % 10),
		0,
	}

	slr.sequence = append(slr.sequence, pr)

	var (
		cyc           Cycle
		i             int
		price, change int8
		ok            bool
	)

	// Starting at 1 because we already appended the starting secret/price above
	// So that serves as our 0 index
	for i = 1; i <= 2000; i++ {

		secret, price, change = NextSecretWithPrice(secret, price, change)

		pr = Product{
			secret,
			price,
			change,
		}

		slr.sequence = append(slr.sequence, pr)

		if i > 3 {
			cyc = Cycle{
				slr.sequence[i-3].change,
				slr.sequence[i-2].change,
				slr.sequence[i-1].change,
				slr.sequence[i].change,
			}
			if _, ok = slr.bestPrices[cyc]; !ok {
				slr.bestPrices[cyc] = price
				if _, ok = m.cycles[cyc]; !ok {
					m.cycles[cyc] = struct{}{}
				}
			}
		}

	}

	m.sellers = append(m.sellers, slr)

}

func NextSecretWithPrice(sec int64, price, change int8) (int64, int8, int8) {

	m1 := sec * 64
	sec ^= m1
	sec %= MOD

	d1 := sec / 32
	sec ^= d1
	sec %= MOD

	m2 := sec * 2048
	sec ^= m2
	sec %= MOD

	newPrice := int8(sec % 10)
	newChange := newPrice - price

	return sec, newPrice, newChange

}

func (m Market) GetBestPrice() int {

	var (
		slr         Seller
		total, best int
		ok          bool
	)

	for cycle := range m.cycles {

		total = 0

		for _, slr = range m.sellers {
			if _, ok = slr.bestPrices[cycle]; ok {
				total += int(slr.bestPrices[cycle])
			}
		}

		best = max(best, total)

	}

	return best

}
