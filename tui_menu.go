package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type item string

func (i item) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := string(i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

func CreateApplyMoveList() list.Model {
	items := []list.Item{
		item("F"),
		item("R"),
		item("L"),
		item("U"),
		item("D"),
		item("B"),
	}
	l := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	l.Title = "What type of move do you want to execute ?\n"
	l.Title += "(Use <- / -> to select alternative moves)"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.SetShowHelp(false)
	return l
}

func CreateExploreMoveList(solution []Move) list.Model {

	items := []list.Item{item("Start")}

	for _, move := range solution {
		items = append(items, item(move.CompactString()))
	}
	l := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	l.Title = "Solution steps"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.SetShowHelp(false)
	return l
}
