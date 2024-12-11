package day09

import (
	"container/list"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
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

// FS for File System ... yup
type FsFile struct {
	id, size int16
}

type FsSpace struct {
	size int16
}

type SpaceRef struct {
	id, size int16
	elm      *list.Element
}

func (d AocDay9) Puzzle2(useSample bool) {

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
		buf                     []byte
		checksum                int64
		n, i, s, b              int
		id, num                 int16
		space                   bool
		newFile, fsf            FsFile
		fileElm, spaceElm       *list.Element
		newSpace, oldSpace, fsp FsSpace
		spRef                   SpaceRef
	)

	seq := list.New()

	files := make([]*list.Element, 0)
	spaces := make([]SpaceRef, 0)

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

			if !space {

				newFile = FsFile{
					id,
					num,
				}
				fileElm = seq.PushBack(newFile)
				files = append(files, fileElm)
				id++

			} else if num > 0 {

				newSpace = FsSpace{
					num,
				}
				spaceElm = seq.PushBack(newSpace)
				spRef = SpaceRef{
					id,
					num,
					spaceElm,
				}
				spaces = append(spaces, spRef)

			}

			space = !space

		}

	}

	var bfore strings.Builder

	for e := seq.Front(); e != nil; e = e.Next() {

		switch e.Value.(type) {
		case FsFile:
			fsf = e.Value.(FsFile)
			bfore.WriteString(fmt.Sprintf("{%d x%d}", fsf.id, fsf.size))
		case FsSpace:
			fsp = e.Value.(FsSpace)
			bfore.WriteString(strings.Repeat(".", int(fsp.size)))
		}

	}

	fmt.Println(bfore.String())
	fmt.Println("")
	fmt.Println("------------")
	fmt.Println("")

	n = len(files) - 1

	for i = n; i > 0; i-- {

		fileElm = files[i]

		for s = 0; s < i; s++ {

			if spaces[s].id >= fileElm.Value.(FsFile).id {
				break
			}

			if spaces[s].size >= fileElm.Value.(FsFile).size {

				// Create a new "space" where the file is
				newSpace = FsSpace{
					fileElm.Value.(FsFile).size,
				}
				seq.InsertAfter(newSpace, fileElm)

				// Move the file to lowest free space (that's big enough)
				seq.MoveBefore(fileElm, spaces[s].elm)

				// Decrease size in slice
				spaces[s].size -= fileElm.Value.(FsFile).size

				// Decrease size in actual list node
				oldSpace = spaces[s].elm.Value.(FsSpace)
				oldSpace.size -= fileElm.Value.(FsFile).size
				spaces[s].elm.Value = oldSpace

				break

			}

		}

	}

	var bout strings.Builder

	for e := seq.Front(); e != nil; e = e.Next() {

		switch e.Value.(type) {
		case FsFile:
			fsf = e.Value.(FsFile)
			for s = 0; s < int(fsf.size); s++ {
				checksum += int64(b * int(fsf.id))
				b++
			}
			bout.WriteString(fmt.Sprintf("{%d x%d}", fsf.id, fsf.size))
		case FsSpace:
			fsp = e.Value.(FsSpace)
			b += int(fsp.size)
			bout.WriteString(strings.Repeat(".", int(fsp.size)))
		}

	}

	fmt.Println(bout.String())

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
