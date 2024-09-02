package main

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/stopwatch"
	tea "github.com/charmbracelet/bubbletea"

	cmn "github.com/mdesoeuv/rubik/common"
)

type ExploreMenu struct {
	cube      cmn.Cube
	backup    cmn.Cube
	solution  SolutionMsg
	list      list.Model
	keymap    keymap
	lastMove  string
	lastIndex int
	help      help.Model
}

func (e ExploreMenu) Update(msg tea.Msg) (Menu, tea.Cmd) {

	var myCmd, listCmd tea.Cmd
	var menu Menu

	e.list, listCmd = e.list.Update(msg)

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		e.list.SetWidth(msg.Width)
		return e, nil

	case tea.KeyMsg:

		switch {

		case key.Matches(msg, e.keymap.quit):
			myCmd = tea.Quit

		case key.Matches(msg, e.keymap.down):
			i, ok := e.list.SelectedItem().(item)
			move := string(i)
			if ok && i != "Start" && e.lastIndex != len(e.list.VisibleItems())-1 {
				choice := move
				move, err := cmn.ParseMove(choice)
				if err != nil {
					// TODO: Better
					fmt.Println(err)
				}
				e.cube.Apply(move)
			}
			e.lastMove = move
			e.lastIndex = e.list.Index()

		case key.Matches(msg, e.keymap.up):
			i, ok := e.list.SelectedItem().(item)
			move := string(i)
			if ok && e.lastMove != "Start" {
				move, err := cmn.ParseMove(e.lastMove)
				if err != nil {
					fmt.Println(err)
				}
				move = move.Reverse()
				e.cube.Apply(move)
			}
			e.lastMove = move
			e.lastIndex = e.list.Index()

		case key.Matches(msg, e.keymap.edit):
			menu = EditMenu{
				cube:      e.backup,
				list:      CreateApplyMoveList(),
				keymap:    NewEditKeyMap(),
				help:      help.New(),
				isSolving: false,
				spinner:   CreateSpinner(),
				stopwatch: stopwatch.NewWithInterval(time.Millisecond),
				solution:  e.solution,
			}
		}
	}
	batch := tea.Batch(myCmd, listCmd)

	if menu != nil {
		return menu, batch
	}
	return e, batch

}

func (e ExploreMenu) View() string {
	s := rectangleStyle.Render(e.cube.Blueprint()) + "\n"
	s += e.list.View()
	s += e.helpView()

	return s
}

func (e ExploreMenu) helpView() string {
	return "\n" + e.help.ShortHelpView([]key.Binding{
		e.keymap.edit,
		e.keymap.quit,
	})
}
