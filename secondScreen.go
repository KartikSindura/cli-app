package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
)

type screenTwoModel struct {
	menu  menu
	table table.Model
}

func makeTable(index int) table.Model {
	var tbl table.Model
	switch index {
	case 0:
		tbl = getAllRecipesByName()
	// case 1:
	// 	tbl = getRecipeByName()
	// case 2:
	// 	tbl = gethRecipeByIngredient()
	// case 3:
	// 	tbl = getFavList()
	// case 4:
	// tbl = editFavList()
	// case 5:
	// 	tbl = getPopularRecipes()
	}
	return tbl
}

func newSecondScreen(m menu) screenTwoModel {
	return screenTwoModel{
		menu:  m,
		// table: makeTable(m.index),
	}
}

func (s screenTwoModel) Init() tea.Cmd { return nil }

func (s screenTwoModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "q":
			return newInitalScreen(s.menu.index).Update(msg)
		case "ctrl+c":
			return s, tea.Quit
		}
		// case toggle:
		// 	switch s.menu.index {
		// 	case 0:
		// 		return s.menu.toggleCasing(), nil
		// 	}
	}
	return s, cmd
}

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

func (s screenTwoModel) View() string {
	return lipgloss.JoinVertical(lipgloss.Center, baseStyle.Render(s.table.View()))
}
