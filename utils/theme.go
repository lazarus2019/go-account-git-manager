package utils

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/bubbles/list"
)

func OrangeBold() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(lipgloss.Color("208")).Bold(true)
}

func OrangeFaint() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(lipgloss.Color("208")).Faint(true)
}

func NewDelegate() list.DefaultDelegate {
	d := list.NewDefaultDelegate()
	d.Styles.SelectedTitle = OrangeBold()
	d.Styles.SelectedDesc = OrangeFaint()
	return d
}
