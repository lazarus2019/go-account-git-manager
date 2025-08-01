// ui/view.go
package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m model) View() string {
	if m.quitting {
		return "Goodbye!\n"
	}

	switch m.mode {
	case modeView:
		// return lipgloss.NewStyle().Padding(1).Render(m.list.View() + "\n" + m.help + "\n" + m.msg)
		return lipgloss.NewStyle().Padding(1).Render(m.list.View() + "\n" + m.help + "\n" + m.msg)
	case modeAdd, modeEdit:
		b.WriteString("Account Details:\n\n")
		labels := []string{"Username", "Email", "Name", "Scope"}
		for i, input := range m.input {
			prefix := "  "
			if i == m.focusIdx {
				prefix = "➤ "
			}
			b.WriteString(fmt.Sprintf("%s%-9s: %s\n", prefix, labels[i], input.View()))
		}
		b.WriteString("\n[Tab] Next • [Shift+Tab] Prev • [Enter] Submit • [Esc] Cancel")
		return lipgloss.NewStyle().Padding(1).Render(b.String())
	
	case modeConfirmDelete:
		return lipgloss.NewStyle().Padding(1).Render("Are you sure you want to delete this account?\n[y] Yes • [n] No")
	default:
		return ""
	}
}
