package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	height int
	width  int
	grid   [][]bool
}

func initalModel() model {
	w, h := 20, 20
	g := make([][]bool, h)
	for i := range g {
		g[i] = make([]bool, w)
	}
	g[1][2], g[2][3], g[3][1], g[3][2], g[3][3] = true, true, true, true, true
	return model{width: w, height: h, grid: g}
}

func (m model) nextGen() [][]bool {
	newGrid := make([][]bool, m.height)
	for y := 0; y < m.height; y++ {
		newGrid[y] = make([]bool, m.width)
		for x := 0; x < m.width; x++ {
			neighbors := m.countNeighbors(x, y)
			if m.grid[y][x] {
				newGrid[y][x] = neighbors == 2 || neighbors == 3
			} else {
				newGrid[y][x] = neighbors == 3
			}
		}
	}
	return newGrid
}

func (m model) countNeighbors(x, y int) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			nx, ny := x+i, y+j
			if nx >= 0 && nx < m.width && ny >= 0 && ny < m.height && m.grid[ny][nx] {
				count++
			}
		}
	}
	return count
}

type tickMsg time.Time

func (m model) Init() tea.Cmd {
	return tea.Tick(time.Millisecond*200, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case tickMsg:
		m.grid = m.nextGen()
		return m, tea.Tick(time.Millisecond*200, func(t time.Time) tea.Msg {
			return tickMsg(t)
		})
	}
	return m, nil
}

func (m model) View() string {
	s := "Game of Life (Press 'q' to quit)\n\n"
	for _, row := range m.grid {
		for _, cell := range row {
			if cell {
				s += "█" // Live cell
			} else {
				s += "·" // Dead cell
			}
		}
		s += "\n"
	}
	return s
}

func main() {
	p := tea.NewProgram(initalModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
