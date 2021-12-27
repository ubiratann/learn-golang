package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("1º - 3º:", slice[:3])
	fmt.Println("5º - last:", slice[4:])
	fmt.Println("2º - 7º:", slice[1:7])
	fmt.Println("3º - 9º:", slice[2:9])
	fmt.Println("last but one item:", slice[len(slice)-2:len(slice)-1])

}
