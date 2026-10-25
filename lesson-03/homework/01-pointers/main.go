package main

import "fmt"

func Increment(value *int) {
	if value != nil {
		*value++
	}
}

func main() {
	number := 10
	Increment(&number)
	fmt.Println(number) // 11

	Increment(&number)
	fmt.Println(number) // 12

	Increment(nil) // программа продолжает работать
	fmt.Println("после nil:", number)
}
