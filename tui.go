package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choices  []string         // items on the to-do list
	cursor   int              // which to-do list item our cursor is pointing at
	selected map[int]struct{} // which to-do items are selected
	cube     *Cube
}

func resetChoices() []string {
	return []string{"F", "R", "L", "U", "D", "B"}
}

func initialModel() model {
	return model{
		choices:  resetChoices(),
		selected: make(map[int]struct{}),
		cube:     NewCubeSolved(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		case "right":
			if len(m.choices[m.cursor]) == 1 {
				m.choices[m.cursor] += "2"
			} else if m.choices[m.cursor][1] == '\'' {
				m.choices[m.cursor] = string(m.choices[m.cursor][0])
			}
			return m, nil

		case "left":
			if len(m.choices[m.cursor]) == 1 {
				m.choices[m.cursor] += "'"
			} else if m.choices[m.cursor][1] == '2' {
				m.choices[m.cursor] = string(m.choices[m.cursor][0])
			}
			return m, nil

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter":
			move, err := ParseMove(m.choices[m.cursor])
			if err != nil {
				fmt.Println(err)
				return m, nil
			}
			m.cube.apply(move)
			m.choices = resetChoices()
			// _, ok := m.selected[m.cursor]
			// if ok {
			//     delete(m.selected, m.cursor)
			// } else {
			//     m.selected[m.cursor] = struct{}{}
			// }
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() string {
	// The header

	s := m.cube.blueprint()

	s += "What type of move do you want to execute ?\n\n"

	// Iterate over our choices
	for i, choice := range m.choices {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if _, ok := m.selected[i]; ok {
			checked = "x" // selected!
		}

		// Render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}
