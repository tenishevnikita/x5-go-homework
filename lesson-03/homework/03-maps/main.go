package main

import "fmt"

func StudentScores() map[string]int {
	scores := make(map[string]int)
	scores["Анна"] = 3
	scores["Борис"] = 0
	scores["Вера"] = 5

	scores["Анна"] += 2

	delete(scores, "Вера")

	return scores
}

func main() {
	scores := StudentScores()
	fmt.Println(scores["Анна"]) // 5

	borisScore, exists := scores["Борис"]
	fmt.Println("Борис:", borisScore, "существует:", exists)

	glebScore, exists := scores["Глеб"]
	fmt.Println("Глеб:", glebScore, "существует:", exists)

	fmt.Println(len(scores)) // 2
}
