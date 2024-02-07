package main

import (
	"fmt"

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
}

func c_NN_initiate() c_NN_model {
	a := c_NN_model{
		menu:    create_NN_start,
		choices: map[int][]string{int(create_NN_start): {"Add layer", "Run network", "Load network", "Save network"}},
	}
	return a
}

func (m c_NN_model) Init() tea.Cmd {
	return nil
}

func (m c_NN_model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	return m, nil
}

func (m c_NN_model) View() string {
	s := ""

	for i, choice := range m.choices[int(m.menu)] {
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	// if done then return "done" string
	return s
}
