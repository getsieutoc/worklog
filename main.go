package main

// An example demonstrating an application with multiple views.
//
// Note that this example was produced before the Bubbles progress component
// was available (github.com/charmbracelet/bubbles/progress) and thus, we're
// implementing a progress bar from scratch here.

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func main() {
	p := tea.NewProgram(initModel())

	if _, err := p.Run(); err != nil {
		fmt.Println("Could not start program:", err)
	}
}

// General stuff for styling the view
var (
	wordStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("211"))
	subtleStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	itemStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))
	mainStyle     = lipgloss.NewStyle().MarginLeft(2)
)

type model struct {
	choices []string
	cursor  int
	chosen  bool
	quit    bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()
		if k == "q" || k == "esc" || k == "ctrl+c" {
			m.quit = true
			return m, tea.Quit
		}
	}

	// Hand off the message and model to the appropriate update function for the
	// appropriate view based on the current state.
	if !m.chosen {
		return updateChoices(msg, m)
	}

	return m, nil
}

// The main view, which just calls the appropriate sub-view
func (m model) View() string {
	var s string
	if m.quit {
		return "\n  See you later!\n\n"
	}

	if !m.chosen {
		s = renderChoices(m)
	} else {
		s = renderChosenView(m)
	}

	return mainStyle.Render("\n" + s + "\n\n")
}

// Update loop for the first view where you're choosing a task.
func updateChoices(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "k", "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "enter":
			m.chosen = true
			return m, nil
		}
	}
	return m, nil
}

// The first view, where you're choosing a task
func renderChoices(m model) string {
	tpl := "What to do today?\n\n"

	for i, choice := range m.choices {
		tpl += fmt.Sprintf(
			"%s\n",
			item(choice, m.cursor == i),
		)
	}

	tpl += "\n"

	tpl += subtleStyle.Render("j/k, up/down: select") + ", " +
		subtleStyle.Render("enter: choose") + ", " +
		subtleStyle.Render("q, esc: quit")

	tpl += "\n"

	return tpl
}

// The second view, after a task has been chosen
func renderChosenView(m model) string {
	var msg string

	switch m.cursor {
	case 0:
		msg = fmt.Sprintln("Render input")
	case 1:
		msg = fmt.Sprintln("Render list of old logs")
	default:
		msg = fmt.Sprintf("It’s always good to see friends.\n\nFetching %s and %s...", wordStyle.Render("social-skills"), wordStyle.Render("conversationutils"))
	}

	return msg + "\n\n"
}

func item(label string, selected bool) string {
	if selected {
		return itemStyle.Render(" • " + label)
	}
	return fmt.Sprintf("   %s", label)
}

func initModel() model {
	return model{
		choices: []string{"Add new log", "View logs"},
		cursor:  0,
		chosen:  false,
		quit:    false,
	}
}
