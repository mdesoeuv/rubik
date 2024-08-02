package main

import (
	"github.com/charmbracelet/bubbles/key"
)

type keymap struct {
	solve   key.Binding
	reset   key.Binding
	quit    key.Binding
	up      key.Binding
	down    key.Binding
	right   key.Binding
	left    key.Binding
	enter   key.Binding
	explore key.Binding
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
		explore: key.NewBinding(
			key.WithKeys("e"),
			key.WithHelp("e", "explore"),
		),
	}
}

func (m model) helpView() string {
	return "\n" + m.help.ShortHelpView([]key.Binding{
		m.keymap.solve,
		m.keymap.explore,
		m.keymap.reset,
		m.keymap.quit,
	})
}
