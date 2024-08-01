package main

import (
	"os"
	"runtime/pprof"
)

var f *os.File

func startProfiling() {
	var err error
	// Create a CPU profile file
	f, err = os.Create("profile.prof")
	if err != nil {
		panic(err)
	}
	// Start CPU profiling
	if err = pprof.StartCPUProfile(f); err != nil {
		panic(err)
	}
}

func stopProfiling() {
	pprof.StopCPUProfile()
	f.Close()
}
