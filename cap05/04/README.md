Exercise description:

1. Create a program that:
- Assign an int value to a variable
- Show this value in decimal, binary and hexadecimal
- Shift the bits of this variable 1 to the left, and assign the result to another variable
- Show this other variable in decimal, binary and hexadecimal

> Code :
```go
import "fmt"

func main() {
	x := 2
	fmt.Printf("x : %d - %b %#x\n", x, x, x)

	y := x << 1
	fmt.Printf("y : %d - %b %#x\n", y, y, y)

}

```

> Output:
```console
x : 2 - 10 0x2
y : 4 - 100 0x4
```