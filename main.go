package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args[1:]
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
}
