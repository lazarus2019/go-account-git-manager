// This file generate by AI
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

type Account struct {
	Username string
	Email    string
	Name     string
	Scope    string
}

func (a Account) Title() string       { return a.Username }
func (a Account) Description() string { return fmt.Sprintf("%s | %s | %s", a.Email, a.Name, a.Scope) }
func (a Account) FilterValue() string { return a.Username }

const dbFile = "accounts.txt"

func loadAccounts() []list.Item {
	file, err := os.Open(dbFile)
	if err != nil {
		return []list.Item{}
	}
	defer file.Close()

	var items []list.Item
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "|")
		if len(parts) == 4 {
			items = append(items, Account{
				Username: parts[0],
				Email:    parts[1],
				Name:     parts[2],
				Scope:    parts[3],
			})
		}
	}
	return items
}

func setGitConfig(acc Account) error {
	cmds := []*exec.Cmd{
		exec.Command("git", "config", "--local", "user.name", acc.Name),
		exec.Command("git", "config", "--local", "user.email", acc.Email),
	}
	for _, cmd := range cmds {
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return err
		}
	}
	return nil
}

type model struct {
	list     list.Model
	quitting bool
	msg      string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		case "enter":
			selected := m.list.SelectedItem().(Account)
			if err := setGitConfig(selected); err != nil {
				m.msg = "Failed to set git config"
			} else {
				m.msg = fmt.Sprintf("Set git config to: %s <%s>", selected.Name, selected.Email)
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.quitting {
		return "Goodbye!\n"
	}
	return lipgloss.NewStyle().Padding(1).Render(m.list.View() + "\n" + m.msg)
}

func main() {
	items := loadAccounts()
	delegate := list.NewDefaultDelegate()
	l := list.New(items, delegate, 80, 20)
	l.Title = "GitHub Accounts"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(true)

	m := model{list: l}
	p := tea.NewProgram(m)

	if err := p.Start(); err != nil {
		log.Fatalf("Error running program: %v", err)
	}
}
