package tui

import (
	"github.com/charmbracelet/bubbles/key"
)

type keymap struct {
	solve   key.Binding
	reset   key.Binding
	shuffle key.Binding
	quit    key.Binding
	up      key.Binding
	down    key.Binding
	right   key.Binding
	left    key.Binding
	enter   key.Binding
	explore key.Binding
	edit    key.Binding
}

func NewEditKeyMap() keymap {
	return keymap{
		solve: key.NewBinding(
			key.WithKeys("s"),
			key.WithHelp("s", "solve"),
		),
		reset: key.NewBinding(
			key.WithKeys("r"),
			key.WithHelp("r", "reset"),
		),
		shuffle: key.NewBinding(
			key.WithKeys("f"),
			key.WithHelp("f", "shuffle"),
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
		explore: key.NewBinding(
			key.WithKeys("e"),
			key.WithHelp("e", "explore"),
		),
	}
}

func NewExploreKeyMap() keymap {
	return keymap{
		quit: key.NewBinding(
			key.WithKeys("q", "ctrl+c"),
			key.WithHelp("q", "quit"),
		),
		down: key.NewBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("down", "move down"),
		),
		up: key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("up", "move up"),
		),
		edit: key.NewBinding(
			key.WithKeys("e"),
			key.WithHelp("e", "edit"),
		),
	}
}
