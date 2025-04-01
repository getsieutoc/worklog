package main

import (
	"github.com/charmbracelet/bubbletea"
)

// FocusArea represents the different focusable areas in the UI
type FocusArea int

const (
	MainMenu FocusArea = iota
	ViewContent
	BottomNav
)

// FocusManager handles focus state and navigation
type FocusManager struct {
	currentArea FocusArea
	focusIndex  int
}

// NewFocusManager creates a new focus manager instance
func NewFocusManager() FocusManager {
	return FocusManager{
		currentArea: MainMenu,
		focusIndex:  0,
	}
}

// MoveFocus handles navigation between focus areas
func (fm *FocusManager) MoveFocus(direction int) {
	switch fm.currentArea {
	case MainMenu:
		// Handle vertical navigation in main menu
		fm.focusIndex += direction
	case ViewContent:
		// Handle navigation within view content
		fm.focusIndex += direction
	case BottomNav:
		// Prevent navigation in bottom nav when in main menu
		if fm.currentArea != MainMenu {
			fm.focusIndex += direction
		}
	}
}

// SetArea changes the current focus area
func (fm *FocusManager) SetArea(area FocusArea) {
	fm.currentArea = area
	fm.focusIndex = 0 // Reset focus index when changing areas
}

// GetFocusedElement returns the currently focused element
func (fm *FocusManager) GetFocusedElement() interface{} {
	switch fm.currentArea {
	case MainMenu:
		return fm.focusIndex
	case ViewContent:
		return fm.focusIndex
	case BottomNav:
		return fm.focusIndex
	}
	return nil
}

// Update handles focus-related messages
func (fm *FocusManager) Update(msg tea.Msg) tea.Cmd {
	if msg, ok := msg.(tea.KeyMsg); ok {
		switch msg.String() {
		case "tab":
			// Cycle through focus areas
			fm.SetArea(FocusArea((int(fm.currentArea) + 1) % 3))
			return nil
		case "shift+tab":
			// Cycle backwards through focus areas
			newArea := int(fm.currentArea) - 1
			if newArea < 0 {
				newArea = 2
			}
			fm.SetArea(FocusArea(newArea))
			return nil
		case "j", "down":
			fm.MoveFocus(1)
			return nil
		case "k", "up":
			fm.MoveFocus(-1)
			return nil
		}
	}
	return nil
}
