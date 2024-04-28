package main

import tea "github.com/charmbracelet/bubbletea"

type toggle struct{}

type menu struct {
	items []menuItem
	index int
}

type menuItem struct {
	text string
	fn   func() tea.Msg
}

type recipe struct {
	Recipeid     int 
	Name         string
	Description  string
	Instructions string
	Preptime     string
	Cooktime     string
	Totaltime    string
}
