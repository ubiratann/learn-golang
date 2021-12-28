Exercise description:

1. Show the remainder by dividing by 4 of all numbers between 10 and 100

> Code:
```go
package main

import "fmt"

func main() {
	for x := 10; x < 101; x++ {
		fmt.Println(x % 4)
	}
}

```

> Output:
```console
2
3
0
1
2
3
0
...
0
1
2
3
0
```