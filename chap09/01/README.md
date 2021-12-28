Exercise description:

1. Using literal composite slices:
    - Create an array wich suports 5 elements of int type
    - Assign values for the indexes
2. Use `range` to show the array values
3. Use format printing to show the array type

> Code:
```go
package main

import "fmt"

func main() {
	array := [5]int{0, 0, 0, 0, 0}

	array[0] = 1
	array[1] = 2
	array[2] = 3

	for _, value := range array {
		fmt.Println(value)
	}
	fmt.Printf("%T\n", array)
}
```

> Output:
```console
1
2
3
[3]int
```