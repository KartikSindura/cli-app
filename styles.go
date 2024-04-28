package main

import "github.com/charmbracelet/lipgloss"

var selectedOptionStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FAFAFA")).
	Background(lipgloss.Color("#AF7DA5"))
	
var helpMenuStyle = lipgloss.NewStyle().
	Faint(true)

var headerStyle = lipgloss.NewStyle().
	Bold(true)