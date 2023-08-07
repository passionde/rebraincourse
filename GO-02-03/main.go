package main

import "fmt"

func Task() {
	daysWeek := []string{
		"Понедельник",
		"Вторник",
		"Среда",
		"Четверг",
		"Пятница",
		"Суббота",
		"Воскресенье",
	}

	// Рабочие дни
	workDays := make([]string, 5)
	copy(workDays, daysWeek)

	// Удаление рабочих дней
	daysWeek = daysWeek[5:]

	fmt.Println("\nЗадание 1")
	fmt.Printf("Рабочие дни: %v\n", workDays)
	fmt.Printf("Выходные дни: %v\n", daysWeek)

	// Объединение
	workDays = append(workDays, daysWeek...)

	fmt.Println("\nЗадание 2")
	fmt.Printf("Дни недели: %v\n", workDays)
}

func main() {
	Task()
}
