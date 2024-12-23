package day17

import (
	"fmt"
	"math"
	"strings"
)

type Program struct {
	Output       []int64
	Instructions []int
	Register     [3]int64
	Ptr          int
	WithChecks   bool
}

const (
	A = 0
	B = 1
	C = 2
)

func (p *Program) Reset(rega int64) {

	p.Register[A] = rega
	p.Register[B] = 0
	p.Register[C] = 0
	p.Ptr = 0
	p.Output = make([]int64, 0)

}

func (p *Program) Adv(operand int) {

	p.Register[A] = p.Register[A] / Pow64(2, p.ComboOperand(operand))

}

func (p *Program) Bxl(operand int) {

	p.Register[B] = p.Register[B] ^ int64(operand)

}

func (p *Program) Bst(operand int) {

	p.Register[B] = p.ComboOperand(operand) % 8

}

func (p *Program) Jnz(operand int) bool {

	if p.Register[A] == 0 {
		return false
	}

	p.Ptr = operand
	return true

}

func (p *Program) Bxc(operand int) {

	p.Register[B] = p.Register[B] ^ p.Register[C]

}

func (p *Program) Out(operand int) int {

	p.Output = append(p.Output, p.ComboOperand(operand)%8)

	if !p.WithChecks {
		return 0
	}

	return p.Check()

}

func (p *Program) Bdv(operand int) {

	p.Register[B] = p.Register[A] / Pow64(2, p.ComboOperand(operand))

}

func (p *Program) Cdv(operand int) {

	p.Register[C] = p.Register[A] / Pow64(2, p.ComboOperand(operand))

}

func (p *Program) Run() bool {

	n := len(p.Instructions)

	var (
		op, chk int
		jumped  bool
	)

	for p.Ptr < n-1 {

		op = p.Ptr + 1
		jumped = false

		switch p.Instructions[p.Ptr] {
		case 0:
			p.Adv(p.Instructions[op])
		case 1:
			p.Bxl(p.Instructions[op])
		case 2:
			p.Bst(p.Instructions[op])
		case 3:
			jumped = p.Jnz(p.Instructions[op])
		case 4:
			p.Bxc(p.Instructions[op])
		case 5:
			chk = p.Out(p.Instructions[op])
		case 6:
			p.Bdv(p.Instructions[op])
		case 7:
			p.Cdv(p.Instructions[op])
		}

		if p.WithChecks {
			if chk < 0 {
				return false
			} else if chk > 0 {
				return true
			}
		}

		if !jumped {
			p.Ptr += 2
		}

	}

	return false

}

func (p Program) Render() string {

	n := len(p.Output)

	if n == 0 {
		return ""
	} else if n == 1 {
		return fmt.Sprintf("%d", p.Instructions[0])
	}

	var b strings.Builder

	b.WriteByte(byte(p.Output[0] + '0'))

	for o := 1; o < n; o++ {
		b.WriteByte(',')
		b.WriteByte(byte(p.Output[o] + '0'))
	}

	return b.String()

}

func (p Program) Check() int {

	in, on := len(p.Instructions), len(p.Output)
	var i int

	for i = 0; i < in && i < on; i++ {
		if int64(p.Instructions[i]) != p.Output[i] {
			return -1
		}
	}

	// Partial match
	if i != in {
		return 0
	}

	return 1

}

func (p Program) ComboOperand(op int) int64 {

	// Should not be used?
	if op >= 7 {
		return 0
	}

	co := int64(op)

	if op > 3 {
		co = p.Register[op-4]
	}

	return co

}

func Pow64(base, exp int64) int64 {
	return int64(math.Pow(float64(base), float64(exp)))
}
