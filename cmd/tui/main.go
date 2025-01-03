package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ethansaxenian/budgeting/tui"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	m, err := tui.NewModel()
	if err != nil {
		log.Fatal(err)
	}

	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		log.Fatal(err)
	}
}
