// utils/storage.go
package utils

import (
	"bufio"
	"os"
	"os/exec"
	"strings"
	"github.com/charmbracelet/bubbles/textinput"
)

const dbFile = "accounts.txt"

func LoadAccounts() []Account {
	file, err := os.Open(dbFile)
	if err != nil {
		return []Account{}
	}
	defer file.Close()

	var accounts []Account
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "|")
		if len(parts) == 4 {
			accounts = append(accounts, Account{
				Username: parts[0],
				Email:    parts[1],
				Name:     parts[2],
				Scope:    parts[3],
			})
		}
	}
	return accounts
}

func SaveAccounts(accounts []Account) {
	f, err := os.Create(dbFile)
	if err != nil {
		return
	}
	defer f.Close()
	for _, a := range accounts {
		line := strings.Join([]string{a.Username, a.Email, a.Name, a.Scope}, "|") + "\n"
		f.WriteString(line)
	}
}

func SetGitConfig(acc Account) error {
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

func NewAccountInputs() []textinput.Model {
	fields := []string{"Username", "Email", "Name", "Scope"}
	inputs := make([]textinput.Model, len(fields))
	for i, ph := range fields {
		ti := textinput.New()
		ti.Placeholder = ph
		ti.Focus()
		inputs[i] = ti
	}
	inputs[0].Focus()
	return inputs
}

func PrefillInputs(a Account) []textinput.Model {
	inputs := NewAccountInputs()
	inputs[0].SetValue(a.Username)
	inputs[1].SetValue(a.Email)
	inputs[2].SetValue(a.Name)
	inputs[3].SetValue(a.Scope)
	return inputs
}

func InputToAccount(inputs []textinput.Model) Account {
	return Account{
		Username: inputs[0].Value(),
		Email:    inputs[1].Value(),
		Name:     inputs[2].Value(),
		Scope:    inputs[3].Value(),
	}
}
