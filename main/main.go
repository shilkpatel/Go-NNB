package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Aaaaaaaaaaaaa: %v", err)
		os.Exit(1)
	}
}

type state int

// state of TUI
const (
	Start state = iota
	NN_settings
	Create_NN

	Database_settings
	Load_database
)

type model struct {
	menu            state
	choices         map[int][]string
	cursor          int
	selected        int
	current_network network
	current_dataset data

	network_selected bool
	dataset_selected bool

	NN_setting c_NN_model
}

func initialModel() model {
	return model{
		menu: 0,
		choices: map[int][]string{0: {"Neural Network Settings", "Database Settings"},
			1: {"Create Neural Network", "Run Neural Network", "Load Neural Network", "Save Neural Network"}},
		cursor:           0,
		selected:         -1,
		network_selected: false,
		dataset_selected: false,
	}
}

// not run by developer
func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k", "w":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j", "s":
			if m.cursor < len(m.choices[int(m.menu)])-1 {
				m.cursor++
			}

		case "enter", " ":

			if m.selected == -1 {
				m.selected = m.cursor
				switch m.menu {
				case Start:
					if m.selected == 0 {
						m.menu = NN_settings
						m.cursor = 0
						m.selected = -1

					} else if m.selected == 1 {
						m.menu = Database_settings
						m.cursor = 0
						m.selected = -1
					}

				case NN_settings:
					if m.selected == 0 {
						m.menu = Create_NN
						m.NN_setting = c_NN_initiate()
						m.cursor = 0
						m.selected = -1

					}

				}
			} else if m.selected == m.cursor {
				m.selected = -1
			}
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() string {
	// The header
	s := "\nSelected Network:\n"
	if m.network_selected {
		s += m.current_network.name
	}
	s += "Selected Database:\n"
	if m.dataset_selected {
		s += m.current_dataset.name
	}

	switch m.menu {
	case Create_NN:
		return s + m.NN_setting.View()
	}

	// Iterate over our choices
	for i, choice := range m.choices[int(m.menu)] {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">"
		}
		/*
			checked := " "
			if m.selected == i {
				checked = "x"
			}
		*/

		//s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}
