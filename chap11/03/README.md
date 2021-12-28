Exercise description:

1. Create a new type named "vehicle"
    - The underlying type must be `struct`
    - Must have the fields: doors and color
1. Create 2 new types: truck and sedan
    - The underlying types must be `struct`
    - Both must have vehicle types
    - The truck type must have a `bool` field called "fourWheelDrive"
    - The sedan type must have a `bool` field called "luxuryModel"
1. Create a sedan and a truck example using composite literal
1. Show the field values from all
1. Show only a field value from each

> Code:
```go
package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWhellDrive bool
}

type sedan struct {
	vehicle
	luxuryModel bool
}

func main() {
	tk := truck{vehicle{2, "pink"}, true}
	sd := sedan{vehicle{4, "blue"}, false}

	fmt.Printf("Truck:\t%v\n", tk)
	fmt.Printf("Sedan:\t%v\n", sd)

	fmt.Printf("Truck color:\t%v\n", tk.color)
	fmt.Printf("Sedan color:\t%v\n", sd.color)

}

```

> Output:
```console
Truck:  {{2 pink} true}
Sedan:  {{4 blue} false}
Truck color:    pink
Sedan color:    blue
```