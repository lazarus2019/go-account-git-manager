// ui/update.go
package ui

import (
	"ghaccount/utils"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch m.mode {
		case modeView:
			switch msg.String() {
			case "ctrl+c", "q":
				m.quitting = true
				return m, tea.Quit
			case "enter":
				selected := m.list.SelectedItem().(utils.Account)
				if err := utils.SetGitConfig(selected); err != nil {
					m.msg = "Failed to set git config"
				} else {
					m.msg = "Git config set to: " + selected.Name + " <" + selected.Email + ">"
				}
			case "a":
				m.mode = modeAdd
				m.input = utils.NewAccountInputs()
			case "e":
				selected := m.list.SelectedItem().(utils.Account)
				m.mode = modeEdit
				m.editingIdx = m.list.Index()
				m.input = utils.PrefillInputs(selected)
			case "d":
				m.mode = modeConfirmDelete
				m.editingIdx = m.list.Index()
			default:
				var cmd tea.Cmd
				m.list, cmd = m.list.Update(msg)
				return m, cmd
			}
		case modeAdd, modeEdit:
			switch {
			case key.Matches(msg, keys.Tab):
				m.focusIdx = (m.focusIdx + 1) % len(m.input)
			case key.Matches(msg, keys.ShiftTab):
				m.focusIdx = (m.focusIdx - 1 + len(m.input)) % len(m.input)
			case msg.String() == "esc":
				m.mode = modeView
			case msg.String() == "enter":
				acc := utils.InputToAccount(m.input)
				// add or edit logic...
				m.mode = modeView
			}
			// Update focus state for each field
			for i := range m.input {
				if i == m.focusIdx {
					m.input[i].Focus()
				} else {
					m.input[i].Blur()
				}
				m.input[i], _ = m.input[i].Update(msg)
			}				
			default:
				for i := range m.input {
					m.input[i], _ = m.input[i].Update(msg)
				}
		case modeConfirmDelete:
			switch msg.String() {
			case "y":
				m.accounts = append(m.accounts[:m.editingIdx], m.accounts[m.editingIdx+1:]...)
				m.list.RemoveItem(m.editingIdx)
				utils.SaveAccounts(m.accounts)
				m.mode = modeView
			case "n", "esc":
				m.mode = modeView
			}
		}
	}
	// var cmd tea.Cmd
	// m.list, cmd = m.list.Update(msg)
	// return m, cmd
}
