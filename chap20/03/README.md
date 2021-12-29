Exercise description:

1. Using goroutines, create an incrementer program:

1. Have a variable with the count value
1. Create a bunch of goroutines, where each one should:
    - Read counter value
    - Save this value
    - Yield the thread with runtime.Gosched()
    - Increment the saved value
    - Copy the new value to the initial variable
1. Use WaitGroups so that your program doesn't finish before your goroutines.
1. Demonstrate that there is a race condition using the -race flag

> Code:
```go
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var gs = runtime.Gosched

func main() {

	maxGoRoutine := 1000
	iterator := 0

	wg.Add(maxGoRoutine)
	for count := 0; count < maxGoRoutine; count++ {
		go func() {
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

```

> Output:
```console
==================
WARNING: DATA RACE
Read at 0x00c000136008 by goroutine 8:
  main.main.func1()
     $GOPATH/src/github.com/ubiratann/learning-golang/chap20/03/main.go:20 +0x34

Previous write at 0x00c000136008 by goroutine 7:
  main.main.func1()
     $GOPATH/src/github.com/ubiratann/learning-golang/chap20/03/main.go:23 +0x64

Goroutine 8 (running) created at:
  main.main()
     $GOPATH/src/github.com/ubiratann/learning-golang/chap20/03/main.go:19 +0x70

Goroutine 7 (finished) created at:
  main.main()
     $GOPATH/src/github.com/ubiratann/learning-golang/chap20/03/main.go:19 +0x70
==================


Final iterator: 810
GoRoutines: 1000


Found 1 data race(s)
exit status 66
```