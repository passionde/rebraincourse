package main

import (
	"fmt"
	"sort"
)

func Task() {
	users := map[string]map[string][]string{
		"Иванов Иван Иванович": {
			"Книги":   {"Книга 1", "Книга 12"},
			"Издания": {"Журнал 1", "Журнал 2"},
		},
		"Михайлов Никита Иванович": {
			"Книги":   {"Книга 45", "Книга 22"},
			"Издания": {"Журнал 5", "Журнал 1", "Журнал 3"},
		},
		"Степанов Михаил Степанович": {
			"Книги":   {"Книга 8"},
			"Издания": {"Журнал 3"},
		},
		"Борисов Сергей Павлович": {
			"Книги":   {"Книга 9"},
			"Издания": {"Журнал 89"},
		},
	}
	fmt.Printf("Количество читателей: %d\n", len(users))

	// Сортировка ключей users
	usersSorted := make([]string, 0, len(users))
	for user := range users {
		usersSorted = append(usersSorted, user)
	}
	sort.Strings(usersSorted)

	// Вывод информации о читателе
	var amount int
	var total int
	for _, user := range usersSorted {
		amount = len(users[user]["Книги"]) + len(users[user]["Издания"])
		total += amount
		fmt.Printf(" - %-30s: %d\n", user, amount)
	}
	fmt.Printf("\n * %-30s: %d\n", "На руках у читателей", total)
}

func main() {
	Task()
}
