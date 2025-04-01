package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// Available actions for the bottom menu
const (
	ActionBack = "back"
	ActionExit = "exit"
)

// Common actions shared across all screens
var CommonActions = map[string]string{
	"b, left": "back",
	"q, esc":  "quit",
}

// RenderBottomMenu combines common actions with screen-specific actions
func RenderBottomMenu(screenActions map[string]string) string {
	// Combine common actions with screen-specific actions
	actions := make(map[string]string)
	
	// Copy common actions first
	for k, v := range CommonActions {
		actions[k] = v
	}
	
	// Add screen-specific actions, which will override commons if there's overlap
	for k, v := range screenActions {
		actions[k] = v
	}
	
	return BottomMenu(actions)
}

// BottomMenu generates a consistently styled bottom menu with navigation commands
func BottomMenu(actions map[string]string) string {
	var menuItems []string
	for key, description := range actions {
		menuItems = append(menuItems, fmt.Sprintf("%s: %s", key, description))
	}

	menu := strings.Join(menuItems, ", ")

	return subtleStyle.Render(menu)
}

// HandleNavigation processes key presses for navigation actions
func HandleNavigation(msg tea.Msg) (bool, string) {
	if msg, ok := msg.(tea.KeyMsg); ok {
		switch msg.String() {
		case "b", "left":
			return true, ActionBack
		case "q", "esc", "ctrl+c":
			return true, ActionExit
		}
	}
	return false, ""
}
