package tui

import (
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/stopwatch"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	cmn "github.com/mdesoeuv/rubik/common"
)

type Menu interface {
	Update(msg tea.Msg) (Menu, tea.Cmd)
	View() string
}

type model struct {
	menu Menu
}

func CreateSpinner() spinner.Model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	return s
}

func InitialModel(c cmn.Cube, solved cmn.Cube) model {
	editMenu := EditMenu{
		cube:      c,
		solved:    solved,
		list:      CreateApplyMoveList(),
		spinner:   CreateSpinner(),
		stopwatch: stopwatch.NewWithInterval(time.Millisecond),
		isSolving: false,
		keymap:    NewEditKeyMap(),
		help:      help.New(),
	}
	editMenu.keymap.explore.SetEnabled(false)

	return model{
		menu: editMenu,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

type SolutionMsg struct {
	moves []cmn.Move
	time  string
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd
	m.menu, cmd = m.menu.Update(msg)

	return m, cmd
}

func (m model) View() string {
	return m.menu.View()
}
