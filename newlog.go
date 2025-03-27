package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// NewLogModel handles the form for creating new logs
type NewLogModel struct {
	inputs     []textinput.Model
	focusIndex int // 0 for title, 1 for description
}

// initNewLogModel creates a new instance of NewLogModel
func InitNewLogModel() *NewLogModel {
	inputs := make([]textinput.Model, 2)
	
	// Title input
	inputs[0] = textinput.New()
	inputs[0].Placeholder = "Enter title"
	inputs[0].Focus()
	inputs[0].CharLimit = 100
	inputs[0].Width = 30
	inputs[0].Prompt = ""

	// Description input
	inputs[1] = textinput.New()
	inputs[1].Placeholder = "Enter description"
	inputs[1].CharLimit = 500
	inputs[1].Width = 50
	inputs[1].Prompt = ""

	return &NewLogModel{
		inputs:     inputs,
		focusIndex: 0,
	}
}

// UpdateNewLog handles updates for the new log form
func UpdateNewLog(msg tea.Msg, m *NewLogModel) (*NewLogModel, tea.Cmd) {
	var cmds []tea.Cmd = make([]tea.Cmd, len(m.inputs))

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab":
			m.focusIndex = (m.focusIndex + 1) % len(m.inputs)
		case "shift+tab":
			m.focusIndex = (m.focusIndex - 1)
			if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs) - 1
			}
		case "enter":
			if m.focusIndex == len(m.inputs)-1 {
				return m, tea.Quit
			}
		}
		
		// Update focus
		for i := range m.inputs {
			m.inputs[i].Blur()
		}
		m.inputs[m.focusIndex].Focus()
	}

	// Update all inputs
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return m, tea.Batch(cmds...)
}

// RenderForm displays the new log form
func RenderForm(m *NewLogModel) string {
	var s string

	// Title field
	titleLabel := labelStyle.Render("Title:")
	titleInput := m.inputs[0].View()
	s += fmt.Sprintf("%s %s\n\n", titleLabel, titleInput)

	// Description field
	descLabel := labelStyle.Render("Description:")
	descInput := m.inputs[1].View()
	s += fmt.Sprintf("%s\n%s\n\n", descLabel, descInput)

	// Help text
	s += subtleStyle.Render("tab: next field") + ", " +
		subtleStyle.Render("shift+tab: prev field") + ", " +
		subtleStyle.Render("enter: save") + ", " +
		subtleStyle.Render("b: back") + ", " +
		subtleStyle.Render("q: quit")

	return s
}
