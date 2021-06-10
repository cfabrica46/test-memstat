package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	m := [][]int{}

	for {
		printMemStatus()
		time.Sleep(time.Second)
		s := make([]int, 999999)
		m = append(m, s)
	}

}

func printMemStatus() {
	var m runtime.MemStats

	runtime.ReadMemStats(&m)

	fmt.Printf("\rAlloc = %v MB", convertByteToMegaByte(m.Alloc))
	fmt.Printf("\tHeapAlloc = %v MB", convertByteToMegaByte(m.HeapAlloc))
	fmt.Printf("\tTotalAlloc = %v MB", convertByteToMegaByte(m.TotalAlloc))
}

func convertByteToMegaByte(b uint64) (r uint64) {
	r = b / 1024 / 1024
	return
}
