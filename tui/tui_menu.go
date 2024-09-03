package tui

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	cmn "github.com/mdesoeuv/rubik/common"
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
			return selectedItemStyle.Render("➤ " + strings.Join(s, " "))
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

func CreateExploreMoveList(solution []cmn.Move) list.Model {

	items := []list.Item{item("Start")}

	for _, move := range solution {
		items = append(items, item(move.String()))
	}
	l := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	l.Title = "Solution steps"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.SetShowHelp(false)
	l.KeyMap.NextPage.SetEnabled(false)
	l.KeyMap.PrevPage.SetEnabled(false)
	l.KeyMap.GoToStart.SetEnabled(false)
	l.KeyMap.GoToEnd.SetEnabled(false)
	return l
}

const listHeight = 11
const defaultWidth = 20

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2).Foreground(lipgloss.Color("170")).Background(lipgloss.Color("241"))
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	rectangleStyle    = lipgloss.NewStyle().MarginLeft(2).MarginTop(1).MarginBottom(1).Padding(1).PaddingLeft(2).PaddingRight(2).BorderStyle(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("170"))
	resultStyle       = lipgloss.NewStyle().MarginLeft(2).Foreground(lipgloss.Color("170"))
)
