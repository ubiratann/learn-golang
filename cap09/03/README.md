Exercise description:

1. Using the previous exercise, use slicing to show the follow values:
	- From first to third item of the slice (including the third)
	- From fifth to last item of the slice (including the last one)
	- From second to seventh of the slice (including the seventh)
	- From third to penultimate item of the slice(including the penultimate)
	- The penultimate item using using the `len` to show it

> Code:
```go
package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("1º - 3º:", slice[:3])
	fmt.Println("5º - last:", slice[4:])
	fmt.Println("2º - 7º:", slice[1:7])
	fmt.Println("3º - 9º:", slice[2:9])
	fmt.Println("penultimate", slice[len(slice)-2:len(slice)-1])

}

```

> Output:
```console
1º - 3º: [1 2 3]
5º - last: [5 6 7 8 9 10]
2º - 7º: [2 3 4 5 6 7]
3º - 9º: [3 4 5 6 7 8 9]
lpenultimate: [9]
```