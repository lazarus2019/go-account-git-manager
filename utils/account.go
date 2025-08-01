// utils/account.go
package utils

import "fmt"

type Account struct {
	Username string
	Email    string
	Name     string
	Scope    string
}

func (a Account) Title() string       { return a.Username }
func (a Account) Description() string { return fmt.Sprintf("%s | %s | %s", a.Email, a.Name, a.Scope) }
func (a Account) FilterValue() string { return a.Username }
