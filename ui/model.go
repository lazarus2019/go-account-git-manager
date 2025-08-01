// ui/model.go
package ui

import (
	"ghaccount/utils"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
)

type modeType int
const (
	modeView modeType = iota
	modeAdd
	modeEdit
	modeConfirmDelete
)

type keyMap struct {
	Tab      key.Binding
	ShiftTab key.Binding
}

var keys = keyMap{
	Tab: key.NewBinding(key.WithKeys("tab"), key.WithHelp("tab", "next field")),
	ShiftTab: key.NewBinding(key.WithKeys("shift+tab"), key.WithHelp("shift+tab", "prev field")),
}

type model struct {
	list       list.Model
	accounts   []utils.Account
	quitting   bool
	help       string
	mode       modeType
	editingIdx int
	msg        string
	input      []textinput.Model
	focusIdx   int
}

func NewModel(accounts []utils.Account) model {
	items := make([]list.Item, len(accounts))
	for i, a := range accounts {
		items[i] = a
	}
	l := list.New(items, utils.NewDelegate(), 80, 20)
	l.Title = "GitHub Accounts"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(true)

	inputs := utils.NewAccountInputs()
	inputs[0].Focus()
	for i := 1; i < len(inputs); i++ {
		inputs[i].Blur()
	}

	return model{
		list:     l,
		help:     helpText(),
		accounts: accounts,
		mode:     modeView,
		input:    inputs,
		focusIdx: 0,
	}
}

func helpText() string {
	return "[↑/↓] Navigate • [Enter] Use • [a] Add • [e] Edit • [d] Delete • [Tab] Next Field • [Shift+Tab] Prev Field • [q] Quit"
}
