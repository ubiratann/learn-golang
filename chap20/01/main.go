package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)

	go func1()
	go func2()

	wg.Wait()
}

func func1() {
	fmt.Println("Func 1")
	wg.Done()
}

func func2() {
	fmt.Println("Func 2")
	wg.Done()
}
