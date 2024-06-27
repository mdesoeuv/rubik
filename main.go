package main

import (
	"fmt"
	"os"
	"strings"
)

type Face = int 

const (
	up 		Face = 0
	down 	Face = iota
	left 	Face = iota
	right 	Face = iota
	front 	Face = iota
	back 	Face = iota
)

type Move struct {
	face 			Face
	clockwise 		bool
	numRotations 	int
}

func parseMove(str string) (move Move) {
	if len(str) == 0 || len(str) > 2 {
		fmt.Println("ERROR: Invalid move string.")
		os.Exit(1)
	}
	move.numRotations = 1
	move.clockwise = true
	for i, c := range str {
		switch i {
		case 0:
			switch c {
			case 'U':
				move.face = up
			case 'D':
				move.face = down
			case 'L':
				move.face = left
			case 'R':
				move.face = right
			case 'F':
				move.face = front
			case 'B':
				move.face =back
			default:
				fmt.Println("ERROR: Invalid face.")
				os.Exit(1)
			}
		case 1:
			switch c {
			case '\'':
				move.clockwise = false
			case '2':
				move.numRotations = 2
			default:
				fmt.Println("ERROR: Invalid rotation.")
				os.Exit(1)
			}
		}
	}
	return
}


func parseMoveList(str string) (moveList []Move) {
	strArray := strings.Split(str, " ")
	for _, moveStr := range strArray {
		move := parseMove(moveStr)
		moveList = append(moveList, move)
	}
	return
}


func main() {

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Usage: go run main.go <move list>")
		return
	}
	moveListStr := args[0]

	for _, arg := range args {
		fmt.Println(arg)
	}

	moveList := parseMoveList(moveListStr)
	for _, move := range moveList {
		fmt.Println(move)
	}
}
