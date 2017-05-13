package main

import (
	"fmt"
	"time"
)

func main() {
	//var data = "Aum"
	//number := 56
	fmt.Println("Hello World !!!", validate("hello", 5))
	fmt.Println("Time", time.Now())
}                  

func validate(input string, number int) int{
	if input == "hello" {
		return 3 * number
	}
	return 0 * number
}