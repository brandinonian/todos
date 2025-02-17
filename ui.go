package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	files    []string
	selected int
	preview  []string
}

func initialModel(items []string) model {
	return model{
		files:    items,
		selected: 0,
		preview:  nil,
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "j", "down":
			if m.selected+1 < len(m.files) {
				m.selected++
			}
		case "k", "up":
			if m.selected > 0 {
				m.selected--
			}
			/* case "e", "enter":
			cmd := exec.Command("nvim", m.files[m.selected])
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Start(); err != nil {
				println("Error running neovim...")
			}
			cmd.Wait()
			return m, tea.Quit */
		}
	}
	return m, nil
}

func (m model) View() string {
	s := ""
	for i, file := range m.files {
		if i == m.selected {
			s += ">"
		} else {
			s += " "
		}

		s += fmt.Sprintf("%d: %s\n", i, file)
	}
	return s
}

func (m model) Init() tea.Cmd {
	return tea.SetWindowTitle("TODO's")
}
