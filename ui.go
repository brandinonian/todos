package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	width    int
	height   int
	files    []file
	selected int
	preview  []string
	err      error
}

type file struct {
	name  string
	lines []string
}

func initialModel(items []file) model {
	return model{
		files:    items,
		selected: 0,
		preview:  items[0].lines,
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "j", "down":
			if m.selected+1 < len(m.files) {
				m.selected++
				m.preview = m.files[m.selected].lines
			}
		case "k", "up":
			if m.selected > 0 {
				m.selected--
				m.preview = m.files[m.selected].lines
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	return lipgloss.JoinHorizontal(lipgloss.Top, m.listView(), m.previewView())
}

func (m model) listView() string {
	s := ""
	for i, file := range m.files {
		if i == m.selected {
			s += ">"
		} else {
			s += " "
		}

		s += fmt.Sprintf("%d: %s\n", i, file.name)
	}
	for len(s) < m.width/2 {
		s += " "
	}
	return s
}

func (m model) previewView() string {
	s := ""
	for i, text := range m.preview {
		s += fmt.Sprintf("%d. %s\n", i+1, text)
	}
	return s
}

func (m model) Init() tea.Cmd {
	return tea.SetWindowTitle("TODO's")
}
