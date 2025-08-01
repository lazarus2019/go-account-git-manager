// main.go
package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"ghaccount/ui"
	"ghaccount/utils"
)

func main() {
	accounts := utils.LoadAccounts()
	m := ui.NewModel(accounts)
	p := tea.NewProgram(m)

	if err := p.Start(); err != nil {
		log.Fatalf("Error running program: %v", err)
	}
}
