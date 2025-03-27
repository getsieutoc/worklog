package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

// NewLogModel handles the form for creating new logs
type NewLogModel struct {
	title       string
	description string
	focusIndex  int // 0 for title, 1 for description
	titleCursor int
	descCursor  int
}

// NewNewLogModel creates a new instance of NewLogModel
func NewNewLogModel() *NewLogModel {
	return &NewLogModel{
		title:       "",
		description: "",
		focusIndex:  0,
		titleCursor: 0,
		descCursor:  0,
	}
}

// UpdateNewLog handles updates for the new log form
func UpdateNewLog(msg tea.Msg, m *NewLogModel) (*NewLogModel, tea.Cmd) {
	if msg, ok := msg.(tea.KeyMsg); ok {
		switch msg.String() {
		case "tab":
			m.focusIndex = (m.focusIndex + 1) % 2
		case "shift+tab":
			m.focusIndex = (m.focusIndex - 1)
			if m.focusIndex < 0 {
				m.focusIndex = 1
			}
		case "enter":
			if m.focusIndex == 1 {
				return m, tea.Quit
			}
		default:
			// Handle typing in the focused field
			if m.focusIndex == 0 {
				m.title = handleInput(msg.String(), m.title, &m.titleCursor)
			} else {
				m.description = handleInput(msg.String(), m.description, &m.descCursor)
			}
		}
	}
	return m, nil
}

// RenderForm displays the new log form
func RenderForm(m *NewLogModel) string {
	var s string

	// Title field
	titleLabel := labelStyle.Render("Title:")
	titleInput := m.title
	if m.focusIndex == 0 {
		titleInput = focusedStyle.Render(titleInput + "█")
	} else {
		titleInput = inputStyle.Render(titleInput)
	}
	s += fmt.Sprintf("%s %s\n\n", titleLabel, titleInput)

	// Description field
	descLabel := labelStyle.Render("Description:")
	descInput := m.description
	if m.focusIndex == 1 {
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
