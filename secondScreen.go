package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
)

type screenTwoModel struct {
	menu  menu
	table table.Model
	answerField textinput.Model
	showAnswerField bool
}

func newSecondScreen(m menu) screenTwoModel {
	ti := textinput.New()
	ti.Placeholder = "Enter name"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return screenTwoModel{
		menu:  m,
		table: table.Model{},
		answerField: ti,
		showAnswerField: false,
	}
}

func (s screenTwoModel) Init() tea.Cmd { 
	return textinput.Blink
 }

func (s screenTwoModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "q":
			return newInitalScreen(s.menu.index).Update(msg)
		case "ctrl+c":
			return s, tea.Quit
		case "enter":
			if s.showAnswerField {
				val := strings.TrimSpace(s.answerField.Value())
				s.showAnswerField = false
				s.table = getRecipeByName(val)
				return s, nil
			}
		}
	case toggle:
		switch s.menu.index {
		case 0:
			s.table = getAllRecipesByName()
			return s, nil
		case 1:
			s.showAnswerField = true
			return s, nil
		}
	}
	s.answerField, cmd = s.answerField.Update(msg)
	return s, cmd
}

var baseStyle = lipgloss.NewStyle().
	BorderForeground(lipgloss.Color("240"))

func (s screenTwoModel) View() string {
	// return lipgloss.JoinVertical(lipgloss.Center, baseStyle.Render(lipgloss.Place()))
	
	// return fmt.Sprintf(baseStyle.Render(s.table.View()), s.answerField.View())
	if s.showAnswerField {
		return fmt.Sprintf(s.answerField.View())
	} else {
		return baseStyle.Render(s.table.View())
	}
}
