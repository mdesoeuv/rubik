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

type ExploreMenu struct {
	lastMove    string
	lastIndex   int
	cube        Cube
	solution    SolutionMsg
	isExploring bool
}

type EditMenu struct {
	spinner   spinner.Model
	stopwatch stopwatch.Model
	isSolving bool
}

type model struct {
	displayedCube Cube
	menu          list.Model
	edit          EditMenu
	explore       ExploreMenu
	keymap        keymap
	help          help.Model
}

func initialModel(c *Cube) model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	cubeCopy := *c

	editMenu := EditMenu{
		spinner:   s,
		stopwatch: stopwatch.NewWithInterval(time.Millisecond),
		isSolving: false,
	}

	exploreMenu := ExploreMenu{
		lastMove:    "Start",
		lastIndex:   0,
		cube:        cubeCopy,
		solution:    SolutionMsg{},
		isExploring: false,
	}

	return model{
		displayedCube: cubeCopy,
		menu:          CreateApplyMoveList(),
		edit:          editMenu,
		explore:       exploreMenu,
		keymap:        NewKeyMap(),
		help:          help.New(),
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
	m.edit.stopwatch, stopWatchCmd = m.edit.stopwatch.Update(msg)
	m.edit.spinner, spinnerCmd = m.edit.spinner.Update(msg)
	m.menu, listCmd = m.menu.Update(msg)

	var myCmd tea.Cmd
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.menu.SetWidth(msg.Width)
		return m, nil

	case SolutionMsg:
		m.edit.isSolving = false
		m.keymap.solve.SetEnabled(true)
		m.keymap.reset.SetEnabled(true)
		m.keymap.explore.SetEnabled(true)

		stopWatchCmd = m.edit.stopwatch.Stop()
		m.explore.solution = msg

	case tea.KeyMsg:

		switch {

		case key.Matches(msg, m.keymap.quit):
			myCmd = tea.Quit

		case key.Matches(msg, m.keymap.right):
			i, ok := m.menu.SelectedItem().(item)
			if ok {
				if len(i) == 1 {
					m.menu.SetItem(m.menu.Index(), item(string(i)+"2"))
				} else if i[1] == '\'' {
					m.menu.SetItem(m.menu.Index(), item(string(i[0])))
				}

			}

		case key.Matches(msg, m.keymap.left):
			i, ok := m.menu.SelectedItem().(item)
			if ok {
				if len(i) == 1 {
					m.menu.SetItem(m.menu.Index(), item(string(i)+"'"))
				} else if i[1] == '2' {
					m.menu.SetItem(m.menu.Index(), item(string(i[0])))
				}
			}

		case key.Matches(msg, m.keymap.solve):
			cube := m.displayedCube
			m.edit.isSolving = true
			m.keymap.solve.SetEnabled(false)
			m.keymap.reset.SetEnabled(false)
			myCmd = tea.Batch(
				m.edit.stopwatch.Reset(),
				m.edit.stopwatch.Start(),
				m.edit.spinner.Tick,
				func() tea.Msg {
					return SolutionMsg{
						moves: cube.solve(),
					}
				},
			)

		case key.Matches(msg, m.keymap.enter):
			i, ok := m.menu.SelectedItem().(item)
			if ok {
				choice := string(i)
				move, err := ParseMove(choice)
				if err != nil {
					// TODO: Better
					fmt.Println(err)
				}
				m.displayedCube.apply(move)
				m.keymap.explore.SetEnabled(false)
			}

		case key.Matches(msg, m.keymap.reset):
			m.keymap.explore.SetEnabled(false)
			m.edit.stopwatch.Reset()
			m.displayedCube = *NewCubeSolved()
			m.explore.solution = SolutionMsg{}

		case key.Matches(msg, m.keymap.explore):
			m.SwitchKeyBindings(m.explore.isExploring)
			if !m.explore.isExploring {
				m.menu = CreateExploreMoveList(m.explore.solution.moves)
				m.explore.cube = m.displayedCube
				m.explore.lastIndex = 0
			} else {
				m.menu = CreateApplyMoveList()
				m.displayedCube = m.explore.cube
			}
			m.explore.isExploring = !m.explore.isExploring

		case key.Matches(msg, m.keymap.down):
			if m.explore.isExploring {
				i, ok := m.menu.SelectedItem().(item)
				move := string(i)
				if ok && i != "Start" && m.explore.lastIndex != len(m.menu.VisibleItems())-1 {
					choice := move
					move, err := ParseMove(choice)
					if err != nil {
						// TODO: Better
						fmt.Println(err)
					}
					m.displayedCube.apply(move)
				}
				m.explore.lastMove = move
				m.explore.lastIndex = m.menu.Index()
			}

		case key.Matches(msg, m.keymap.up):
			if m.explore.isExploring {
				i, ok := m.menu.SelectedItem().(item)
				move := string(i)
				if ok && m.explore.lastMove != "Start" {
					move, err := ParseMove(m.explore.lastMove)
					if err != nil {
						fmt.Println(err)
					}
					move = move.Reverse()
					m.displayedCube.apply(move)
				}
				m.explore.lastMove = move
				m.explore.lastIndex = m.menu.Index()
			}

		}

	}
	return m, tea.Batch(myCmd, stopWatchCmd, spinnerCmd, listCmd)
}

func (m model) View() string {

	s := rectangleStyle.Render(m.displayedCube.blueprint()) + "\n"

	s += m.menu.View()

	if m.edit.isSolving {
		s += resultStyle.Render("\n" + m.edit.spinner.View() + "Solving..." + fmt.Sprintf(" (%s)", m.edit.stopwatch.View()))
	} else if m.explore.solution.moves != nil {
		solutionString := "\nSolution found: "

		for _, move := range m.explore.solution.moves {
			solutionString += move.CompactString() + " "
		}
		solutionString += fmt.Sprintf("(%s)", m.edit.stopwatch.View())
		s += resultStyle.Render(solutionString)
	}
	s += m.helpView()

	return s
}
