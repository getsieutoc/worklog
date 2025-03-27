package main

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

// Styles for the UI
var (
	subtleStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	itemStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))
	mainStyle   = lipgloss.NewStyle().MarginLeft(2)

	// Form styles
	labelStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("69"))
	inputStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("255"))
	focusedStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("212")).Bold(true)
	textareaStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("255"))
)

type model struct {
	choices    []string
	cursor     int
	chosen     string
	quit       bool
	newLogModel *NewLogModel
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
		if k == "b" || k == "left" {
			m.chosen = ""
			return m, nil
		}
	}

	// Hand off the message and model to the appropriate update function for the
	// appropriate view based on the current state.
	if m.chosen == "" {
		return updateChoices(msg, m)
	}

	return updateChosen(msg, m)
}

// The main view, which just calls the appropriate sub-view
func (m model) View() string {
	var s string
	if m.quit {
		return "\n  See you later!\n\n"
	}

	if m.chosen == "" {
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
			m.chosen = m.choices[m.cursor]
			if m.chosen == "Add new log" {
				m.newLogModel = &NewLogModel{
					title:       "",
					description: "",
					focusIndex:  0,
					titleCursor: 0,
					descCursor:  0,
				}
			}
			return m, nil
		}
	}
	return m, nil
}

// Update loop for the chosen view
func updateChosen(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	if m.newLogModel == nil {
		return m, nil
	}

	if msg, ok := msg.(tea.KeyMsg); ok {
		switch msg.String() {
		case "tab":
			m.newLogModel.focusIndex = (m.newLogModel.focusIndex + 1) % 2
		case "shift+tab":
			m.newLogModel.focusIndex = (m.newLogModel.focusIndex - 1)
			if m.newLogModel.focusIndex < 0 {
				m.newLogModel.focusIndex = 1
			}
		case "enter":
			if m.newLogModel.focusIndex == 1 {
				// TODO: Save the log entry
				m.chosen = ""
				m.newLogModel = nil
				return m, nil
			}
		default:
			// Handle typing in the focused field
			if m.newLogModel.focusIndex == 0 {
				m.newLogModel.title = handleInput(msg.String(), m.newLogModel.title, &m.newLogModel.titleCursor)
			} else {
				m.newLogModel.description = handleInput(msg.String(), m.newLogModel.description, &m.newLogModel.descCursor)
			}
		}
	}
	return m, nil
}

func handleInput(key, text string, cursor *int) string {
	if key == "backspace" && len(text) > 0 && *cursor > 0 {
		text = text[:*cursor-1] + text[*cursor:]
		*cursor--
	} else if len(key) == 1 { // Single character
		if *cursor == len(text) {
			text += key
		} else {
			text = text[:*cursor] + key + text[*cursor:]
		}
		*cursor++
	}
	return text
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
	switch m.chosen {
	case m.choices[0]:
		return renderForm(m)
	case m.choices[1]:
		return fmt.Sprintln("Render list of old logs")
	default:
		return fmt.Sprintf("Unknown choice: %s\n\n", m.chosen)
	}
}

func renderForm(m model) string {
	if m.newLogModel == nil {
		return ""
	}

	var s string

	// Title field
	titleLabel := labelStyle.Render("Title:")
	titleInput := m.newLogModel.title
	if m.newLogModel.focusIndex == 0 {
		titleInput = focusedStyle.Render(titleInput + "█")
	} else {
		titleInput = inputStyle.Render(titleInput)
	}
	s += fmt.Sprintf("%s %s\n\n", titleLabel, titleInput)

	// Description field
	descLabel := labelStyle.Render("Description:")
	descInput := m.newLogModel.description
	if m.newLogModel.focusIndex == 1 {
		descInput = focusedStyle.Render(descInput + "█")
	} else {
		descInput = textareaStyle.Render(descInput)
	}
	s += fmt.Sprintf("%s\n%s\n\n", descLabel, descInput)

	// Help text
	s += subtleStyle.Render("tab: next field") + ", " +
		subtleStyle.Render("shift+tab: prev field") + ", " +
		subtleStyle.Render("enter: save") + ", " +
		subtleStyle.Render("b: back") + ", " +
		subtleStyle.Render("q: quit")

	return s
}

func item(label string, selected bool) string {
	if selected {
		return itemStyle.Render(" • " + label)
	}
	return fmt.Sprintf("   %s", label)
}

func initModel() model {
	return model{
		choices:     []string{"Add new log", "View logs"},
		cursor:      0,
		chosen:      "",
		quit:        false,
		newLogModel: nil,
	}
}
