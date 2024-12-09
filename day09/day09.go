package day09

import (
	"fmt"
	"io"
	"log"
	"os"
)

type AocDay9 struct{}

type FileBlock struct {
	id int16
}

const DIR = "day09/"
const READ = 512

func (d AocDay9) Puzzle1(useSample bool) {

	datafile := DIR + "data.txt"
	if useSample {
		datafile = DIR + "sample.txt"
	}

	f, err := os.Open(datafile)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	var (
		buf          []byte
		checksum     int64
		n, i, j, num int
		id           int16
		space        bool
	)

	seq := make([]FileBlock, 0)

	for {

		buf = make([]byte, READ)

		n, err = io.ReadAtLeast(f, buf, READ)

		if err == io.ErrUnexpectedEOF {
			f.Seek(int64(-n), io.SeekCurrent)
			n, _ = io.ReadAtLeast(f, buf, n)
		}

		if n == 0 {
			break
		}

		for i = 0; i < n; i++ {

			num = int(buf[i] - '0')

			if !space {

				for j = 0; j < num; j++ {
					seq = append(seq, FileBlock{id})
				}
				id++

			} else {

				for j = 0; j < num; j++ {
					seq = append(seq, FileBlock{-1})
				}

			}

			space = !space

		}

	}

	n = len(seq)
	j = n - 1

	if useSample {
		RenderSeq(seq)
	}

	for i = 0; i < n; i++ {

		if seq[i].id >= 0 {
			continue
		}

		for seq[j].id < 0 {
			j--
		}

		if i >= j {
			break
		}

		seq[i], seq[j] = seq[j], seq[i]

		if useSample {
			RenderSeq(seq)
		}

	}

	for i = 0; i < n; i++ {

		if seq[i].id < 0 {
			continue
		}

		checksum += int64(i * int(seq[i].id))

	}

	fmt.Println("")
	fmt.Println("Checksum: ", checksum)

}

func (d AocDay9) Puzzle2(useSample bool) {

}

func RenderSeq(seq []FileBlock) {

	for _, s := range seq {
		if s.id >= 0 {
			fmt.Printf("%d", s.id)
		} else {
			fmt.Print(".")
		}
	}

	fmt.Print("\n")

}
