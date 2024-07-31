package main

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/stopwatch"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	choices   []string
	cursor    int
	selected  map[int]struct{}
	cube      *Cube
	solution  DoneSolving
	loader    spinner.Model
	isSolving bool
	stopwatch stopwatch.Model
	keymap    keymap
	help      help.Model
}

type keymap struct {
	solve key.Binding
	reset key.Binding
	quit  key.Binding
	up    key.Binding
	down  key.Binding
	right key.Binding
	left  key.Binding
	enter key.Binding
}

func NewKeyMap() keymap {
	return keymap{
		solve: key.NewBinding(
			key.WithKeys("s"),
			key.WithHelp("s", "start solving"),
		),
		reset: key.NewBinding(
			key.WithKeys("r"),
			key.WithHelp("r", "reset the cube"),
		),
		quit: key.NewBinding(
			key.WithKeys("q", "ctrl+c"),
			key.WithHelp("q", "quit"),
		),
		up: key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("up", "move up"),
		),
		down: key.NewBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("down", "move down"),
		),
		right: key.NewBinding(
			key.WithKeys("right", "l"),
			key.WithHelp("right", "move right"),
		),
		left: key.NewBinding(
			key.WithKeys("left", "h"),
			key.WithHelp("left", "move left"),
		),
		enter: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "validate"),
		),
	}
}

func (m model) helpView() string {
	return "\n" + m.help.ShortHelpView([]key.Binding{
		m.keymap.solve,
		m.keymap.reset,
		m.keymap.quit,
	})
}

func resetChoices() []string {
	return []string{"F", "R", "L", "U", "D", "B"}
}

func initialModel(c *Cube) model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return model{
		choices:   resetChoices(),
		selected:  make(map[int]struct{}),
		cube:      c,
		loader:    s,
		isSolving: false,
		stopwatch: stopwatch.NewWithInterval(time.Millisecond),
		help:      help.New(),
		keymap:    NewKeyMap(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

type DoneSolving struct {
	states *[]Cube
	moves  []Move
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var stopWatchCmd tea.Cmd
	var loaderCmd tea.Cmd
	m.stopwatch, stopWatchCmd = m.stopwatch.Update(msg)
	m.loader, loaderCmd = m.loader.Update(msg)

	var myCmd tea.Cmd
	switch msg := msg.(type) {

	case DoneSolving:
		m.isSolving = false
		m.stopwatch.Stop()
		m.solution = msg

	case tea.KeyMsg:

		switch {

		case key.Matches(msg, m.keymap.quit):
			myCmd = tea.Quit

		case key.Matches(msg, m.keymap.up):
			if m.cursor > 0 {
				m.cursor--
			}

		case key.Matches(msg, m.keymap.down):
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		case key.Matches(msg, m.keymap.right):
			if len(m.choices[m.cursor]) == 1 {
				m.choices[m.cursor] += "2"
			} else if m.choices[m.cursor][1] == '\'' {
				m.choices[m.cursor] = string(m.choices[m.cursor][0])
			}

		case key.Matches(msg, m.keymap.left):
			if len(m.choices[m.cursor]) == 1 {
				m.choices[m.cursor] += "'"
			} else if m.choices[m.cursor][1] == '2' {
				m.choices[m.cursor] = string(m.choices[m.cursor][0])
			}

		case key.Matches(msg, m.keymap.solve):
			cube := *m.cube
			m.isSolving = true
			myCmd = tea.Batch(
				m.stopwatch.Start(),
				m.loader.Tick,
				func() tea.Msg {
					return DoneSolving{
						states: cube.solve(),
						moves:  AllMoves,
					}
				},
			)

		case key.Matches(msg, m.keymap.enter):
			move, err := ParseMove(m.choices[m.cursor])
			if err != nil {
				// TODO: Better
				fmt.Println(err)
			}
			m.cube.apply(move)
			m.choices = resetChoices()

		case key.Matches(msg, m.keymap.reset):
			m.stopwatch.Reset()
			m.cube = NewCubeSolved()
			m.solution = DoneSolving{}
		}
	}
	return m, tea.Batch(myCmd, stopWatchCmd, loaderCmd)
}

func (m model) View() string {

	// The header

	s := m.cube.blueprint()

	s += "What type of move do you want to execute ?\n"
	s += "(Use <- / -> to select alternative moves)\n\n"

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

	if m.isSolving {
		s += m.loader.View() + "Solving..." + fmt.Sprintf(" (%s)", m.stopwatch.View()) + "\n"
	} else if m.solution.moves != nil {
		s += "Solution found: "
		for _, move := range m.solution.moves {
			s += move.CompactString() + " "
		}
		s += "\n"
	}

	// The footer
	s += m.helpView()

	// Send the UI for rendering
	return s
}
