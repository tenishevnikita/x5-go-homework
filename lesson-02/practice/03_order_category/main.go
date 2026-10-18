package main

import "fmt"

func printOrderCategory(amount int) {
	var category string
	switch {
	case amount < 0:
		fmt.Println("Ошибка: сумма заказа не может быть отрицательной")
		return
	case amount < 10000:
		category = "small"
	case amount < 50000:
		category = "regular"
	default:
		category = "large"
	}
	fmt.Printf("Сумма заказа: %d коп.\n", amount)
	fmt.Printf("Категория: %s\n", category)
}

func main() {
	printOrderCategory(0)
	fmt.Println()
	printOrderCategory(9999)
	fmt.Println()
	printOrderCategory(10000)
	fmt.Println()
	printOrderCategory(49999)
	fmt.Println()
	printOrderCategory(50000)
	fmt.Println()
	printOrderCategory(-1)
}
