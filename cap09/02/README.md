Exercise description:
1. Using literal composite:
    - Create an slice of int type
    - Assign 10 values 
2. Use `range` to show the slice values
3. Use format printing to show the slice type

> Code:
```go
package main

import "fmt"

func main() {
	slice := []int{}

	slice = append(slice, 1)
	slice = append(slice, 2)
	slice = append(slice, 3)

	for _, value := range slice {
		fmt.Println(value)
	}

	fmt.Printf("%T\n", slice)
}
```

> Output:
```console
1
2
3
[]int
```