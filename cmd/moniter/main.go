package main

import (
	"fmt"
	"runtime"
)

func main() {
	var memStatus runtime.MemStats
	runtime.ReadMemStats(&memStatus)
	fmt.Printf("Mem Status: %v \n", memStatus.Alloc)
}
