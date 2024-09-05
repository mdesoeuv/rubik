package main

import (
	"flag"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	cmn "github.com/mdesoeuv/rubik/common"
	tui "github.com/mdesoeuv/rubik/tui"
	vc "github.com/mdesoeuv/rubik/visual_cepo"
)

var (
	tuiFlag     = flag.Bool("tui", false, "Enable Terminal User Interface")
	profileFlag = flag.Bool("profile", false, "Enable CPU profiling")
	verboseFlag = flag.Bool("verbose", false, "Enable verbose output")
)

func main() {
	var err error
	flag.Parse()

	if *profileFlag {
		startProfiling()
		defer stopProfiling()
	}

	args := flag.Args()
	if len(args) != 1 && !*tuiFlag {
		fmt.Println("Usage: go run main.go <move list>")
		return
	}

	var (
		moveListStr = ""
		moveList    = []cmn.Move{}
	)
	if len(args) > 0 {
		moveListStr = args[0]
		moveList, err = cmn.ParseMoveList(moveListStr)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	cube := vc.NewCubeSolved()
	solvedCube := cube.Clone()

	for _, move := range moveList {
		cube.Apply(move)
	}
	solver := cube.NewSolver()
	if *tuiFlag {
		p := tea.NewProgram(tui.InitialModel(cube, solvedCube, solver))
		if _, err := p.Run(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
		}
	} else {
		solution := solver.Solve(cube)
		if *verboseFlag {
			fmt.Println(cube.Blueprint())
			fmt.Printf("Solution found in %v steps: ", len(solution))
		}
		output := ""
		for _, move := range solution {
			output += move.String() + " "
		}
		fmt.Println(output)
	}
}
