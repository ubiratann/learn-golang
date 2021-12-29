package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var gs = runtime.Gosched
var mt sync.Mutex

func main() {

	maxGoRoutine := 1000
	iterator := 0

	wg.Add(maxGoRoutine)
	for count := 0; count < maxGoRoutine; count++ {
		go func() {

			mt.Lock()
			defer mt.Unlock()

			x := iterator
			gs()
			x++
			iterator = x

			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("\n\nFinal iterator: %v\nGoRoutines: %v\n\n\n", iterator, maxGoRoutine)
}
