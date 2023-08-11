package main

import (
	"fmt"
	"math"
)

func Task1() {
	var A *int
	var B = 8

	A = &B
	fmt.Println("\nЗадание 1")
	fmt.Println(*A)

	*A = 34
	fmt.Println(B)
}

func Task2() {
	DValue := 35.0
	RValue := DValue / 2.0

	R := &RValue

	S := (*R) * (*R) * math.Pi

	fmt.Println("\nЗадание 2")
	fmt.Printf("R = %.2f, S = %.2f\n", *R, S)
}

func main() {
	Task1()
	Task2()
}
