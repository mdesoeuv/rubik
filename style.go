package main

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

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
