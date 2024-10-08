package tui

import (
	"fmt"
	"math/rand/v2"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/stopwatch"
	tea "github.com/charmbracelet/bubbletea"

	cmn "github.com/mdesoeuv/rubik/common"
)

type EditMenu struct {
	cube      cmn.Cube
	solved    cmn.Cube
	solver    cmn.Solver
	list      list.Model
	solution  SolutionMsg
	spinner   spinner.Model
	stopwatch stopwatch.Model
	isSolving bool
	keymap    keymap
	help      help.Model
}

func (e EditMenu) Update(msg tea.Msg) (Menu, tea.Cmd) {

	var stopWatchCmd, spinnerCmd, listCmd, myCmd tea.Cmd

	var menu Menu

	e.stopwatch, stopWatchCmd = e.stopwatch.Update(msg)
	e.spinner, spinnerCmd = e.spinner.Update(msg)
	e.list, listCmd = e.list.Update(msg)

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		e.list.SetWidth(msg.Width)
		return e, nil

	case SolutionMsg:
		e.isSolving = false
		e.keymap.solve.SetEnabled(true)
		e.keymap.reset.SetEnabled(true)
		e.keymap.shuffle.SetEnabled(true)
		e.keymap.explore.SetEnabled(true)
		stopWatchCmd = e.stopwatch.Stop()
		msg.time = e.stopwatch.View()
		e.solution = msg

	case tea.KeyMsg:

		switch {

		case key.Matches(msg, e.keymap.quit):
			myCmd = tea.Quit

		case key.Matches(msg, e.keymap.right):
			i, ok := e.list.SelectedItem().(item)
			if ok {
				if len(i) == 1 {
					e.list.SetItem(e.list.Index(), item(string(i)+"2"))
				} else if i[1] == '\'' {
					e.list.SetItem(e.list.Index(), item(string(i[0])))
				}
			}

		case key.Matches(msg, e.keymap.left):
			i, ok := e.list.SelectedItem().(item)
			if ok {
				if len(i) == 1 {
					e.list.SetItem(e.list.Index(), item(string(i)+"'"))
				} else if i[1] == '2' {
					e.list.SetItem(e.list.Index(), item(string(i[0])))
				}
			}

		case key.Matches(msg, e.keymap.solve):
			e.isSolving = true
			e.keymap.solve.SetEnabled(false)
			e.keymap.reset.SetEnabled(false)
			e.keymap.shuffle.SetEnabled(false)
			myCmd = tea.Batch(
				e.stopwatch.Reset(),
				e.stopwatch.Start(),
				e.spinner.Tick,
				func() tea.Msg {
					return SolutionMsg{
						moves: e.solver.Solve(e.cube),
					}
				},
			)

		case key.Matches(msg, e.keymap.enter):
			i, ok := e.list.SelectedItem().(item)
			if ok {
				choice := string(i)
				move, err := cmn.ParseMove(choice)
				if err != nil {
					panic("invalid move")
				}
				e.cube.Apply(move)
				e.keymap.explore.SetEnabled(false)
			}

		case key.Matches(msg, e.keymap.reset):
			e.keymap.explore.SetEnabled(false)
			e.stopwatch.Reset()
			e.cube = e.solved.Clone()
			e.solution = SolutionMsg{}

		case key.Matches(msg, e.keymap.shuffle):
			e.keymap.explore.SetEnabled(false)
			e.stopwatch.Reset()
			r := rand.New(rand.NewPCG(rand.Uint64(), rand.Uint64()))
			n := rand.IntN(100) + 20
			cmn.Shuffle(e.cube, r, n)
			e.solution = SolutionMsg{}

		case key.Matches(msg, e.keymap.explore):
			menu = ExploreMenu{
				lastMove:  "Start",
				lastIndex: 0,
				cube:      e.cube,
				backup:    e.cube.Clone(),
				solved:    e.solved,
				solver:    e.solver,
				solution:  e.solution,
				list:      CreateExploreMoveList(e.solution.moves),
				keymap:    NewExploreKeyMap(),
				help:      help.New(),
			}

		}

	}

	batch := tea.Batch(myCmd, stopWatchCmd, spinnerCmd, listCmd)

	if menu != nil {
		return menu, batch
	}
	return e, batch
}

func (e EditMenu) View() string {

	s := rectangleStyle.Render(e.cube.Blueprint()) + "\n"
	s += e.list.View()

	if e.isSolving {
		s += resultStyle.Render("\n" + e.spinner.View() + "Solving..." + fmt.Sprintf(" (%s)", e.stopwatch.View()))
	} else if e.solution.moves != nil {
		solutionString := "\nSolution found: "
		solutionString += fmt.Sprintf("(%v moves in %s) ", len(e.solution.moves), e.solution.time)
		for _, move := range e.solution.moves {
			solutionString += move.String() + " "
		}
		s += resultStyle.Render(solutionString)
	}

	s += e.helpView()

	return s
}

func (e EditMenu) helpView() string {
	return "\n" + e.help.ShortHelpView([]key.Binding{
		e.keymap.solve,
		e.keymap.explore,
		e.keymap.reset,
		e.keymap.shuffle,
		e.keymap.quit,
	})
}
