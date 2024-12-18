package day15

/*
 * Puzzle 2 version of "Bubble"
 */

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type modelt struct {
	sub       chan struct{}
	w, h      int
	robot     Point
	warehouse [][]byte
	moves     []byte
	result    string
}

func (m modelt) Init() tea.Cmd {
	return tea.Batch(
		listenForUpdate(m.sub),
		waitForUpdate(m.sub),
	)
}

func (m modelt) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		return m, tea.Quit
	case updateMsg:
		m.Action()
		return m, waitForUpdate(m.sub)
	default:
		return m, nil
	}
}

func (m modelt) View() string {

	var (
		x, y int
	)

	s := "\n" + wallStyle.Render(strings.Repeat("#", m.w)) + "\n"

	// Skipping top/bottom as we render them as single string above/below
	for y = 1; y < m.h-1; y++ {
		for x = 0; x < m.w; x++ {
			switch m.warehouse[y][x] {
			case SPACE:
				s += spaceStyle.Render(".")
			case WALL:
				s += wallStyle.Render("#")
			case ROBOT:
				s += robotStyle.Render("@")
			default:
				s += boxStyle.Render(string(m.warehouse[y][x]))
			}

		}
		s += "\n"
	}

	s += wallStyle.Render(strings.Repeat("#", m.w)) + "\n\n"

	if len(m.moves) > 0 {
		mn := min(m.w, len(m.moves))
		s += moveStyle.Render(string(m.moves[0])) + spaceStyle.Render(string(m.moves[1:mn])) + "\n"
	} else {
		s += moveStyle.Render("Complete :)") + "\n"
	}

	if len(m.result) > 0 {
		s += "\n" + moveStyle.Render(m.result) + "\n"
	}

	return s

}

func (m *modelt) Action() {

	if len(m.moves) > 0 {
		m.Move()
		return
	}

	if len(m.result) == 0 {
		m.Sum()
		return
	}

}

func (m *modelt) Move() {

	var (
		move, spot, pl byte
	)

	move, m.moves = m.moves[0], m.moves[1:]

	next := Point{m.robot.x, m.robot.y}
	shift := Point{0, 0}

	switch move {
	case MOVEUP:
		next.y--
		shift.y--
	case MOVEDOWN:
		next.y++
		shift.y++
	case MOVELEFT:
		next.x--
		shift.x--
	case MOVERIGHT:
		next.x++
		shift.x++
	}

	spot = m.warehouse[next.y][next.x]

	if spot == WALL {
		return
	}

	if spot == SPACE {
		m.Swap(m.robot, next)
		m.robot.x, m.robot.y = next.x, next.y
		return
	}

	// Else, it's a box <(^_^)>

	pulse := Point{next.x, next.y}
	canMove := false

	// Moving left or right
	if shift.y == 0 {

		for {

			pulse.x += shift.x
			pl = m.warehouse[pulse.y][pulse.x]

			if pl == BOXLEFT || pl == BOXRIGHT {
				continue
			}
			if pl == WALL {
				break
			}
			if pl == SPACE {
				canMove = true
				break
			}

		}

		if !canMove {
			return
		}

		// "Bubble" back...
		tail := Point{pulse.x, pulse.y}

		for pulse.x != next.x {
			tail.x -= shift.x
			m.Swap(pulse, tail)
			pulse.x -= shift.x
		}

		m.Swap(m.robot, next)
		m.robot.x, m.robot.y = next.x, next.y

	} else {

		// Moving up or down, the harder bit... :')

		box := Point{next.x, next.y}
		if spot == BOXRIGHT {
			box.x--
		}

		canMove = m.CanMoveVertical(box, shift.y == -1)

		if !canMove {
			return
		}

		m.DoMoveVertical(box, shift.y == -1)

		m.Swap(m.robot, next)
		m.robot.x, m.robot.y = next.x, next.y

	}

}

// Pass left point of box, e.g. [] coord of the '['
func (m modelt) CanMoveVertical(box Point, up bool) bool {

	l, r := box.x, box.x+1
	y, d := box.y, 0

	if up {
		y--
		d--
	} else {
		y++
		d++
	}

	// log.Println("Checking box ", box, " | ", up)

	// Both points above (or below) the box are spaces? Then we're clear...
	if m.warehouse[y][l] == SPACE && m.warehouse[y][r] == SPACE {
		// log.Println("[/] Clear above")
		return true
	}

	// Either point above is a wall? Then we're super(man ie, imovable)
	if m.warehouse[y][l] == WALL || m.warehouse[y][r] == WALL {
		// log.Println("[x] Wall above")
		return false
	}

	upLeft := Point{box.x, box.y + d}
	upRight := Point{box.x + 1, box.y + d}

	// Box directly above
	if m.warehouse[y][l] == BOXLEFT {
		return m.CanMoveVertical(upLeft, up)
	}

	leftClear, rightClear := true, true

	if m.warehouse[y][l] == BOXRIGHT {
		boxUpLeft := Point{upLeft.x - 1, upLeft.y}
		leftClear = m.CanMoveVertical(boxUpLeft, up)
	}
	if m.warehouse[y][r] == BOXLEFT {
		rightClear = m.CanMoveVertical(upRight, up)
	}

	return leftClear && rightClear

}

// This one works like the checker above, but in reverse
// Pulse down to move the last box first
func (m *modelt) DoMoveVertical(boxLeft Point, up bool) {

	l, r := boxLeft.x, boxLeft.x+1
	y, d := boxLeft.y, 0

	if up {
		y--
		d--
	} else {
		y++
		d++
	}

	boxRight := Point{boxLeft.x + 1, boxLeft.y}

	upLeft := Point{boxLeft.x, boxLeft.y + d}
	upRight := Point{boxRight.x, boxLeft.y + d}

	// Box directly above
	if m.warehouse[y][l] == BOXLEFT {
		m.DoMoveVertical(upLeft, up)
	}

	// Boxes on either side
	if m.warehouse[y][l] == BOXRIGHT {
		boxUpLeft := Point{upLeft.x - 1, upLeft.y}
		m.DoMoveVertical(boxUpLeft, up)
	}
	if m.warehouse[y][r] == BOXLEFT {
		m.DoMoveVertical(upRight, up)
	}

	// Clear ?!
	if m.warehouse[y][l] == SPACE && m.warehouse[y][r] == SPACE {
		m.Swap(boxLeft, upLeft)
		m.Swap(boxRight, upRight)
	}

}

func (m *modelt) Sum() {

	var x, y, sum int

	for y = 1; y < m.h-1; y++ {
		for x = 1; x < m.w-1; x++ {
			if m.warehouse[y][x] == BOXLEFT {
				sum += 100*y + x
			}
		}
	}

	m.result = fmt.Sprintf("GPS Sum: %d", sum)

}

func (m *modelt) Swap(a, b Point) {
	m.warehouse[a.y][a.x], m.warehouse[b.y][b.x] = m.warehouse[b.y][b.x], m.warehouse[a.y][a.x]
}
