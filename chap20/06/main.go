package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("OS: %s\tArch: %s\n", runtime.GOOS, runtime.GOARCH)
}
