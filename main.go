package main

import (
	"flag"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	tuiFlag = flag.Bool("tui", false, "Enable Terminal User Interface")
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("Usage: go run main.go <move list>")
		return
	}

	moveListStr := args[0]

	moveList, err := ParseMoveList(moveListStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	cube := NewCubeSolved()

	for _, move := range moveList {
		cube.apply(move)
	}
	if *tuiFlag {
		p := tea.NewProgram(initialModel(cube))
		if _, err := p.Run(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
		}
	} else {
		fmt.Println(cube.blueprint())
	}
}
