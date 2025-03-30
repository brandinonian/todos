package main

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	path := "."

	if len(os.Args) > 2 {
		panic("Usage: todos <path>")
	}

	if len(os.Args) == 2 {
		path = os.Args[1]
	}

	err, results := SearchDir(path)
	if err != nil {
		println("Error searching directory...")
		panic(err)
	}

	p := tea.NewProgram(initialModel(results), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		println("Error running tea program...")
		panic(err)
	}
}
