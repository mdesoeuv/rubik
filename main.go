package main

import (
	"flag"
	"fmt"
)

var (
	tuiFlag = flag.Bool("tui", false, "Enable Terminal User Interface")
)

func main() {
	flag.Parse()
	fmt.Println("tuiFlag:", *tuiFlag)
	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("Usage: go run main.go <move list>")
		return
	}
	for _, arg := range args {
		fmt.Println(arg)
	}

	moveListStr := args[0]

	moveList, err := ParseMoveList(moveListStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, move := range moveList {
		fmt.Println(move)
	}

	cube := NewCubeSolved()
	cube.print()
}
