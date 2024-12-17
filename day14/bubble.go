package day14

import (
	"container/ring"
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

/*
 * https://github.com/charmbracelet/bubbletea/blob/main/examples/realtime/main.go
 * Based on this example
 */

var (
	baubleStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#e73007"))
	treeStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#16df16"))
	trailStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#327b32"))
	fadeStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#283b28"))
	dotStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#808380"))
)

type updateMsg struct{}

func listenForUpdate(sub chan struct{}, factory *Factory) tea.Cmd {
	return func() tea.Msg {
		for {
			time.Sleep(time.Millisecond * 500)
			factory.Run()
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
	sub             chan struct{}
	secs, w, h, cyc int
	data            *ring.Ring
	factory         *Factory
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		listenForUpdate(m.sub, m.factory),
		waitForUpdate(m.sub),
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		return m, tea.Quit
	case updateMsg:
		m.secs += 1
		m.cyc = (m.cyc + 1) % TRAIL
		return m, waitForUpdate(m.sub)
	default:
		return m, nil
	}
}

func (m model) View() string {

	var (
		x, y int
		s    string
	)

	refs := make([][][]uint8, TRAIL)
	cyc := make([]int, TRAIL)

	for x = 0; x < TRAIL; x++ {
		refs[x] = m.data.Value.([][]uint8)
		m.data = m.data.Next()
		cyc[x] = (m.cyc + x) % TRAIL
	}

	for y = 0; y < m.h; y++ {
		for x = 0; x < m.w; x++ {
			if refs[cyc[0]][y][x] > 0 {
				if refs[cyc[0]][y][x] > 1 {
					s += baubleStyle.Render("@")
				} else {
					s += treeStyle.Render(string('0' + refs[cyc[0]][y][x]))
				}
			} else if refs[cyc[1]][y][x] > 0 {
				if refs[cyc[1]][y][x] > 1 {
					s += trailStyle.Render("@")
				} else {
					s += trailStyle.Render(string('0' + refs[cyc[1]][y][x]))
				}
			} else if refs[cyc[2]][y][x] > 0 {
				if refs[cyc[2]][y][x] > 1 {
					s += fadeStyle.Render("@")
				} else {
					s += fadeStyle.Render(string('0' + refs[cyc[2]][y][x]))
				}
			} else {
				s += dotStyle.Render(".")
			}
		}
		s += "\n"
	}

	s += "\n"
	s += fmt.Sprintln(m.secs)

	return s

}
