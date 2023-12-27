package main

import (
	"os"
	"runtime/pprof"
)

func headMem(){

	heapFile, hErr := os.Create("heap.prof")
	if hErr != nil{
		panic(hErr)
	}

	defer heapFile.Close()

	randGen()
	
	if err := pprof.WriteHeapProfile(heapFile); err!=nil{
		panic(err)
	}

}