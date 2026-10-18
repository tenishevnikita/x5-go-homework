package main

import "fmt"

func printOrderStats(orders []int) {
	var countOrders, minOrder, maxOrder, totalAmount int
	if len(orders) > 0 {
		minOrder = orders[0]
		maxOrder = orders[0]
	} else {
		fmt.Println("Нет заказов для анализа")
		return
	}

	for _, value := range orders {
		if value < minOrder {
			minOrder = value
		}
		if value > maxOrder {
			maxOrder = value
		}
		totalAmount += value
		countOrders++
	}

	avgOrder := float64(totalAmount) / float64(countOrders)

	fmt.Printf("Количество заказов: %d\n", countOrders)
	fmt.Printf("Минимальный заказ: %d коп.\n", minOrder)
	fmt.Printf("Максимальный заказ: %d коп.\n", maxOrder)
	fmt.Printf("Общая сумма: %d коп.\n", totalAmount)
	fmt.Printf("Средний заказ: %v коп.\n", avgOrder)
}

func main() {
	fmt.Println([]int{12000, 7500, 31000, 4500})
	printOrderStats([]int{12000, 7500, 31000, 4500})
	fmt.Println()
	fmt.Println([]int{5000})
	printOrderStats([]int{5000})
	fmt.Println()
	fmt.Println([]int{100, 100, 100})
	printOrderStats([]int{100, 100, 100})
	fmt.Println()
	fmt.Println([]int{100, 101})
	printOrderStats([]int{100, 101})
	fmt.Println()
	fmt.Println([]int{})
	printOrderStats([]int{})
}
