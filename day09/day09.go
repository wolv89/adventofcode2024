package day09

import (
	"container/list"
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

func (d AocDay9) Puzzle1(useSample int) {

	datafile := DIR + "data.txt"
	if useSample == 1 {
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

	if useSample == 1 {
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

		if useSample == 1 {
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

type Item struct {
	id, size    int16
	file, moved bool
}

/*
 * NOT WORKING :/
 * Close, but not close enough...
 */
func (d AocDay9) Puzzle2(useSample int) {

	datafile := DIR + "data.txt"
	if useSample == 1 {
		datafile = DIR + "sample.txt"
	}

	f, err := os.Open(datafile)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	var (
		buf        []byte
		checksum   int64
		n, i, s, b int
		id, num    int16
		isSpace    bool
	)

	seq := list.New()

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

			num = int16(buf[i] - '0')

			if !isSpace {

				seq.PushBack(Item{
					id,
					num,
					true, // File
					false,
				})
				id++

			} else if num > 0 {

				seq.PushBack(Item{
					id,
					num,
					false, // Space
					false,
				})

			}

			isSpace = !isSpace

		}

	}

	var (
		item, space Item
		it, sp, nsp *list.Element
		// b4, aft     strings.Builder
	)

	/*
		for it = seq.Front(); it != nil; it = it.Next() {

			item = it.Value.(Item)

			if item.file {
				b4.WriteString(fmt.Sprintf("{%d x%d}", item.id, item.size))
			} else {
				b4.WriteString(strings.Repeat(".", int(item.size)))
			}

		}

		fmt.Println(b4.String())
		fmt.Println("")
	*/

	for it = seq.Back(); it != nil; it = it.Prev() {

		item = it.Value.(Item)

		if item.moved || !item.file {
			continue
		}

		for sp = seq.Front(); sp != nil; sp = sp.Next() {

			space = sp.Value.(Item)

			if space.file || space.size < item.size {
				continue
			}

			if space.id >= item.id {
				break
			}

			nsp = seq.InsertAfter(Item{
				item.id,
				item.size,
				false,
				false,
			}, it)

			seq.MoveBefore(it, sp)

			space.size -= item.size
			sp.Value = space

			item.moved = true
			it.Value = item

			// Reset pointer to continue loop
			it = nsp
			break

		}

	}

	for it = seq.Front(); it != nil; it = it.Next() {

		item = it.Value.(Item)

		if item.file {
			for s = 0; s < int(item.size); s++ {
				checksum += int64(b * int(item.id))
				b++
			}
			// aft.WriteString(fmt.Sprintf("{%d x%d}", item.id, item.size))
		} else {
			b += int(item.size)
			// aft.WriteString(strings.Repeat(".", int(item.size)))
		}

	}

	// fmt.Println(aft.String())
	// fmt.Println("")

	fmt.Println("")
	fmt.Println("Checksum: ", checksum)

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
