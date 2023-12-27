package main


import (
	"os"
	"runtime/pprof"
)

func cpu() {

	file, error := os.Create("cpuProfile.prof")
	if error != nil {
		panic(error)
	}

	defer file.Close()

	if err := pprof.StartCPUProfile(file); err != nil {
		panic(err)
	}

	defer pprof.StopCPUProfile()

	randGen()
}


