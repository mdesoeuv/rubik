package main

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/stopwatch"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	cube        Cube
	initialCube Cube
	solution    SolutionMsg
	list        list.Model
	spinner     spinner.Model
	stopwatch   stopwatch.Model
	keymap      keymap
	help        help.Model
	isSolving   bool
	isExploring bool
	lastMove    string
}

func initialModel(c *Cube) model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	cubeCopy := *c

	return model{
		cube:        cubeCopy,
		initialCube: *c,
		list:        CreateApplyMoveList(),
		spinner:     s,
		stopwatch:   stopwatch.NewWithInterval(time.Millisecond),
		keymap:      NewKeyMap(),
		help:        help.New(),
		isSolving:   false,
		isExploring: false,
		lastMove:    "Start",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

type SolutionMsg struct {
	moves []Move
}

func (m *model) SwitchKeyBindings(isExploring bool) {
	m.keymap.enter.SetEnabled(isExploring)
	m.keymap.left.SetEnabled(isExploring)
	m.keymap.right.SetEnabled(isExploring)
	m.keymap.reset.SetEnabled(isExploring)
	m.keymap.solve.SetEnabled(isExploring)
	if isExploring {
		m.keymap.explore.SetHelp("e", "explore")
	} else {
		m.keymap.explore.SetHelp("e", "edit")
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var stopWatchCmd tea.Cmd
	var spinnerCmd tea.Cmd
	var listCmd tea.Cmd
	m.stopwatch, stopWatchCmd = m.stopwatch.Update(msg)
	m.spinner, spinnerCmd = m.spinner.Update(msg)
	m.list, listCmd = m.list.Update(msg)

	var myCmd tea.Cmd
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case SolutionMsg:
		m.isSolving = false
		m.keymap.solve.SetEnabled(true)
		m.keymap.reset.SetEnabled(true)
		m.keymap.explore.SetEnabled(true)

		stopWatchCmd = m.stopwatch.Stop()
		m.solution = msg

	case tea.KeyMsg:

		switch {

		case key.Matches(msg, m.keymap.quit):
			myCmd = tea.Quit

		case key.Matches(msg, m.keymap.right):
			i, ok := m.list.SelectedItem().(item)
			if ok {
				if len(i) == 1 {
					m.list.SetItem(m.list.Index(), item(string(i)+"2"))
				} else if i[1] == '\'' {
					m.list.SetItem(m.list.Index(), item(string(i[0])))
				}

			}

		case key.Matches(msg, m.keymap.left):
			i, ok := m.list.SelectedItem().(item)
			if ok {
				if len(i) == 1 {
					m.list.SetItem(m.list.Index(), item(string(i)+"'"))
				} else if i[1] == '2' {
					m.list.SetItem(m.list.Index(), item(string(i[0])))
				}
			}

		case key.Matches(msg, m.keymap.solve):
			cube := m.cube
			m.isSolving = true
			m.keymap.solve.SetEnabled(false)
			m.keymap.reset.SetEnabled(false)
			myCmd = tea.Batch(
				m.stopwatch.Start(),
				m.spinner.Tick,
				func() tea.Msg {
					return SolutionMsg{
						moves: cube.solve(),
					}
				},
			)

		case key.Matches(msg, m.keymap.enter):
			i, ok := m.list.SelectedItem().(item)
			if ok {
				choice := string(i)
				move, err := ParseMove(choice)
				if err != nil {
					// TODO: Better
					fmt.Println(err)
				}
				m.cube.apply(move)
				m.keymap.explore.SetEnabled(false)
			}

		case key.Matches(msg, m.keymap.reset):
			m.stopwatch.Reset()
			m.cube = *NewCubeSolved()
			m.solution = SolutionMsg{}

		case key.Matches(msg, m.keymap.explore):
			m.SwitchKeyBindings(m.isExploring)
			if !m.isExploring {
				m.list = CreateExploreMoveList(m.solution.moves)
				m.initialCube = m.cube
			} else {
				m.list = CreateApplyMoveList()
				m.cube = m.initialCube
			}
			m.isExploring = !m.isExploring

		case key.Matches(msg, m.keymap.down):
			if m.isExploring {
				i, ok := m.list.SelectedItem().(item)
				move := string(i)
				if ok && move != "Solved" && i != "Start" {
					choice := move
					move, err := ParseMove(choice)
					if err != nil {
						// TODO: Better
						fmt.Println(err)
					}
					m.cube.apply(move)
				}
				m.lastMove = move
			}

		case key.Matches(msg, m.keymap.up):
			if m.isExploring {
				i, ok := m.list.SelectedItem().(item)
				move := string(i)
				if ok && m.lastMove != "Start" {
					move, err := ParseMove(m.lastMove)
					if err != nil {
						fmt.Println(err)
					}
					move = move.Reverse()
					m.cube.apply(move)
				}
				m.lastMove = move
			}

		}

	}
	return m, tea.Batch(myCmd, stopWatchCmd, spinnerCmd, listCmd)
}

func (m model) View() string {

	s := rectangleStyle.Render(m.cube.blueprint()) + "\n"

	s += m.list.View()

	if m.isSolving {
		s += resultStyle.Render("\n" + m.spinner.View() + "Solving..." + fmt.Sprintf(" (%s)", m.stopwatch.View()))
	} else if m.solution.moves != nil {
		solutionString := "\nSolution found: "

		for _, move := range m.solution.moves {
			solutionString += move.CompactString() + " "
		}
		solutionString += fmt.Sprintf("(%s)", m.stopwatch.View())
		s += resultStyle.Render(solutionString)
	}
	s += m.helpView()

	return s
}
