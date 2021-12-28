Exercise description:

1. Create a loop using the  `for condition {}` syntax and use it to show the years since you were born.

> Code:
```go
package main

import "fmt"

func main() {
	year := 1999

	for year <= 2021 {
		fmt.Println(year)
		year++
	}
}

```

> Output:
```console
1999
2000
2001
...
2019
2020
2021
```