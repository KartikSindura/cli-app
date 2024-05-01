package main

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type toggle struct{}

type menu struct {
	items []menuItem
	index int
}

type menuItem struct {
	text string
	msg  tea.Msg
}

type Recipe struct {
	Recipeid     int
	Name         string
	Description  string
	Instructions string
	Preptime     time.Time
	Cooktime     time.Time
	Totaltime    time.Time
}

type lenRecipe struct {
	lenid    int
	lenname  int
	lendesc  int
	lenins   int
	lenprep  int
	lencook  int
	lentotal int
}
