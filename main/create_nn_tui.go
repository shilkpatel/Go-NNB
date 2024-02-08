package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	create_NN_start state = iota
	Add_layer
	Run_NN
	Load_NN
	Save_NN
)

type c_NN_model struct {
	menu        state
	choices     map[int][]string
	cursor      int
	new_network network
	layers      string
	textbox_1   textinput.Model
	textbox_2   textinput.Model
	nodes       int
	transfer    activation
}

func c_NN_initiate() c_NN_model {
	a := c_NN_model{
		menu:      create_NN_start,
		choices:   map[int][]string{0: {"Add layer", "Run network", "Load network", "Save network"}},
		textbox_1: textinput.New(),
		textbox_2: textinput.New(),
	}
	return a
}

func (m c_NN_model) Init() tea.Cmd {
	return nil
}

func (m c_NN_model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			switch m.menu {
			case Add_layer:
				if m.cursor == 1 {
					m.textbox_1.Focus()
					m.cursor = 0
				}

			default:
				if m.cursor > 0 {
					m.cursor--
				}
			}

		case "down":
			switch m.menu {
			case Add_layer:
				if m.cursor == 0 {
					m.textbox_2.Focus()
					m.cursor = 1
				}
			default:
				if m.cursor < len(m.choices[int(m.menu)])-1 {
					m.cursor++
				}
			}

		case "enter":
			switch m.menu {
			case create_NN_start:
				switch m.cursor {
				case 0:
					m.menu = Add_layer
					m.textbox_1.Focus()
					var cmd tea.Cmd
					m.textbox_1, cmd = m.textbox_1.Update(msg)
					cmds = append(cmds, cmd)
				}
			}
		}

	}
	var cmd tea.Cmd
	if m.menu == Add_layer {
		if m.cursor == 0 {
			m.textbox_1, cmd = m.textbox_1.Update(msg)
		} else {
			m.textbox_2, cmd = m.textbox_2.Update(msg)
		}
	}
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m c_NN_model) View() string {
	s := ""
	s += m.layers

	for i, choice := range m.choices[int(m.menu)] {
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	switch m.menu {
	case Add_layer:
		s += "Enter number of nodes\n"
		s += m.textbox_1.View()
		s += "\nEnter activation function\n"
		s += m.textbox_2.View()
	}

	// if done then return "done" string
	return s
}
