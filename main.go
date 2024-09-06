package main

import (
	"flag"
	"fmt"
	"math/rand/v2"
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
	shuffleFlag = flag.Int("shuffle", 0, "Shuffle the cube with N moves")
)

func main() {
	var err error
	flag.Parse()

	if *profileFlag {
		startProfiling()
		defer stopProfiling()
	}

	var (
		moveListStr = ""
		moveList    = []cmn.Move{}
	)

	args := flag.Args()
	if len(args) > 0 {
		moveListStr = args[0]
		moveList, err = cmn.ParseMoveList(moveListStr)
		if err != nil {
			fmt.Print("invalid move sequence: ")
			fmt.Println(err)
			return
		}
	}

	if len(args) > 1 {
		fmt.Println("invalid move sequence: expected only one argument.")
		return
	}

	cube := vc.NewCubeSolved()
	solvedCube := cube.Clone()

	for _, move := range moveList {
		cube.Apply(move)
	}
	solver := cube.NewSolver()

	if *shuffleFlag > 0 {
		if *verboseFlag {
			fmt.Println("Shuffling cube: ", *shuffleFlag)
		}
		r := rand.New(rand.NewPCG(rand.Uint64(), rand.Uint64()))
		cmn.Shuffle(cube, r, *shuffleFlag)
	}

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
