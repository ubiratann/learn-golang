Exercise description:

1. Write expressions using the follow operators, and assign their values to variables

- !=
- ==
- \<=
- \>=
- \<
- \>

> Code :
```go
import "fmt"

func main() {
	x := 2
	y := 2

	a := x != y
	b := x == y
	fmt.Printf("a:%v - b:%v\n", a, b)

	a = x <= y
	b = x >= y
	fmt.Printf("a:%v - b:%v\n", a, b)

	a = x < y
	b = x > y
	fmt.Printf("a:%v - b:%v\n", a, b)

}

```

> Output:
```console
a:false - b:true
a:true - b:true
a:false - b:false
```