package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func (m menu) toggleCasing() tea.Model {
	m.items[m.index].text = "hello"
	return m
}

func (m menu) Init() tea.Cmd {
	return tea.SetWindowTitle("My app")
}

func (m menu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	last := len(m.items) - 1
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "w":
			if m.index > 0 {
				m.index--
			} else if m.index == 0 {
				m.index = last
			}
			return m, nil
		case "down", "s":
			if m.index < last {
				m.index++
			} else if m.index == last {
				m.index = 0
			}
			return m, nil
		case "ctrl+c":
			return m, tea.Quit
		case "enter", " ":
			// return m, m.items[m.index].fn
			return newSecondScreen(m).Update(msg)
		}
		// case toggle:
		// 	switch m.index {
		// 	case 0:
		// 		return m.toggleCasing(), nil
		// 	}
	}
	return m, cmd
}

func (m menu) View() string {
	var options []string
	for i, o := range m.items {
		if i == m.index {
			options = append(options, selectedOptionStyle.Render(fmt.Sprintf(" > %s", o.text)))
		} else {
			options = append(options, fmt.Sprintf("   %s", o.text))
		}
	}
	return lipgloss.JoinVertical(lipgloss.Left, headerStyle.Render("   Some Recipe App\n"), strings.Join(options, "\n"), helpMenuStyle.Render("\n   Enter/space to select, arrow keys to navigate, or Ctrl+C to exit."))
}

var conn *pgx.Conn

func newInitalScreen(prevIndex int) menu {
	m := menu{
		items: []menuItem{
			menuItem{
				text: "See a list of all recipes by name",
				fn:   func() tea.Msg { return toggle{} },
			},
			menuItem{
				text: "Search for a recipe by name",
				fn:   func() tea.Msg { return toggle{} },
			},
			menuItem{
				text: "Search for a list of recipes by ingredient",
				fn:   func() tea.Msg { return toggle{} },
			},
			menuItem{
				text: "View your favourites list",
				fn:   func() tea.Msg { return toggle{} },
			},
			menuItem{
				text: "Edit your favourites list",
				fn:   func() tea.Msg { return toggle{} },
			},
			menuItem{
				text: "See a list of the most popular recipes",
				fn:   func() tea.Msg { return toggle{} },
			},
		},
		index: prevIndex,
	}
	return m
}

func main() {
	// defer conn.Close(context.Background())
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer f.Close()

	p := tea.NewProgram(newInitalScreen(0), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}

	envFile, _ := godotenv.Read(".env")
	conn, err = pgx.Connect(context.Background(), envFile["DATABASE_URL"])
	if err != nil {
		log.Fatal(err)
	}
	query, _ := conn.Query(context.Background(), "select * from recipe")
	if err != nil {
		log.Fatal(err)
	}
	rows, _ := pgx.CollectRows(query, pgx.RowToStructByName[recipe])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(rows)
}
