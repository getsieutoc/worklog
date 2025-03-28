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

const (
	hotPink  = lipgloss.Color("#FF06B7")
	darkGray = lipgloss.Color("#767676")
)

// Styles for the UI
var (
	subtleStyle = lipgloss.NewStyle().Foreground(darkGray)
	itemStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))
	mainStyle   = lipgloss.NewStyle().MarginLeft(2)

	// Form styles
	labelStyle = lipgloss.NewStyle().Foreground(hotPink)
	// inputStyle    = lipgloss.NewStyle().Foreground(hotPink)
	// focusedStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("212")).Bold(true)
	// textareaStyle = lipgloss.NewStyle().Foreground(darkGray)
)

type model struct {
	choices     []string
	cursor      int
	chosen      string
	quit        bool
	newLogModel *NewLogModel
	focus       FocusManager
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Handle application-specific commands first
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

	// Then let the focus manager handle navigation events
	if cmd := m.focus.Update(msg); cmd != nil {
		return m, cmd
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
			if m.chosen == m.choices[0] { // "Add new log"
				m.newLogModel = InitNewLogModel()
			}
			return m, nil
		}
	}
	return m, nil
}

// Update loop for the chosen view
func updateChosen(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	if m.newLogModel != nil {
		var cmd tea.Cmd
		m.newLogModel, cmd = UpdateNewLog(msg, m.newLogModel)
		if m.newLogModel == nil {
			m.chosen = ""
			return m, cmd
		}
		return m, cmd
	}
	return m, nil
}

// The first view, where you're choosing a task
func renderChoices(m model) string {
	s := "What to do today?\n\n"

	// Render main menu with focus indication
	menuStyle := lipgloss.NewStyle()
	if m.focus.currentArea == MainMenu {
		menuStyle = menuStyle.Border(lipgloss.RoundedBorder()).BorderForeground(hotPink)
	}

	menuContent := ""
	for i, choice := range m.choices {
		menuContent += fmt.Sprintf(
			"%s\n",
			item(choice, m.cursor == i),
		)
	}
	s += menuStyle.Render(menuContent)

	// Render static bottom navigation
	screenActions := map[string]string{
		"j/k, up/down": "select",
		"enter":        "choose",
	}
	navContent := RenderBottomMenu(screenActions)
	s += "\n" + subtleStyle.Render(navContent)

	s += "\n"

	return s
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
		return "Error: newLogModel was not initialized.\n\n"
	}
	return RenderForm(m.newLogModel)
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
		focus:       NewFocusManager(),
	}
}
