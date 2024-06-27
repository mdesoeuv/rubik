package main

import (
	"fmt"
	"os"
)

type Face = int 

const (
	up 		Face = 0
	down 	Face = iota
	left 	Face = iota
	right 	Face = iota
	front 	Face = iota
	rear 	Face = iota
)

type Move struct {
	face Face
	clockwise bool
}

func main() {

	args := os.Args[1:]
	for _, arg := range args {
		fmt.Println(arg)
	}
}
