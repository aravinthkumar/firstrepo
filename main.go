package main

import (
	"fmt"
)

func sumOfNumbers(a *int) {
	fmt.Println("During Function Call:", *a, a)
	*a = 2
	fmt.Println("During Function Call Value is changed", *a, a)

}

func main() {
	var a int = 1
	fmt.Println("Before Function Call: Value and Address", a, &a)
	sumOfNumbers(&a)
	fmt.Println("After Function Call: Value and Address", a, &a)
}
