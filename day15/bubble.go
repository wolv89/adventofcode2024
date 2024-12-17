package day15

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	robotStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#edb479"))
	boxStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#3764c9"))
	wallStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#404040"))
	spaceStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#404040"))
	moveStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#ffffff"))
)

type updateMsg struct{}

func listenForUpdate(sub chan struct{}) tea.Cmd {
	return func() tea.Msg {
		for {
			time.Sleep(time.Millisecond * SPEED)
			sub <- struct{}{}
		}
	}
}

func waitForUpdate(sub chan struct{}) tea.Cmd {
	return func() tea.Msg {
		return updateMsg(<-sub)
	}
}

type model struct {
	sub       chan struct{}
	w, h      int
	robot     Point
	warehouse [][]byte
	moves     []byte
	result    string
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		listenForUpdate(m.sub),
		waitForUpdate(m.sub),
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m model) View() string {

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
			case BOX:
				s += boxStyle.Render("O")
			case WALL:
				s += wallStyle.Render("#")
			case ROBOT:
				s += robotStyle.Render("@")
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

func (m *model) Action() {

	if len(m.moves) > 0 {
		m.Move()
		return
	}

	if len(m.result) == 0 {
		m.Sum()
		return
	}

}

func (m *model) Move() {

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
		m.warehouse[m.robot.y][m.robot.x], m.warehouse[next.y][next.x] = m.warehouse[next.y][next.x], m.warehouse[m.robot.y][m.robot.x]
		m.robot.x, m.robot.y = next.x, next.y
		return
	}

	// Else, it's a box <(^_^)>

	pulse := Point{next.x, next.y}
	canMove := false

	for {

		pulse.x += shift.x
		pulse.y += shift.y

		pl = m.warehouse[pulse.y][pulse.x]

		if pl == BOX {
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

	// Swap box with space
	m.warehouse[pulse.y][pulse.x], m.warehouse[next.y][next.x] = m.warehouse[next.y][next.x], m.warehouse[pulse.y][pulse.x]

	// Then do normal move
	m.warehouse[m.robot.y][m.robot.x], m.warehouse[next.y][next.x] = m.warehouse[next.y][next.x], m.warehouse[m.robot.y][m.robot.x]
	m.robot.x, m.robot.y = next.x, next.y

}

func (m *model) Sum() {

	var x, y, sum int

	for y = 1; y < m.h-1; y++ {
		for x = 1; x < m.w-1; x++ {
			if m.warehouse[y][x] == BOX {
				sum += 100*y + x
			}
		}
	}

	m.result = fmt.Sprintf("GPS Sum: %d", sum)

}
