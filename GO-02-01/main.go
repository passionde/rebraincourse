package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func Task1() {
	foo := "104"
	bar := 35

	fooInt, _ := strconv.Atoi(foo)
	barString := fmt.Sprintf("%d", bar)

	fmt.Println("\nЗадание 1")
	fmt.Printf(
		"foo(type): %s, fooInt(type): %s\n",
		reflect.TypeOf(foo).String(), reflect.TypeOf(fooInt).String(),
	)

	fmt.Printf(
		"bar(type): %s, barString(type): %s\n",
		reflect.TypeOf(bar).String(), reflect.TypeOf(barString).String(),
	)
}

func Task2() {
	const meterPerSecToKilometerPerHour float64 = 3.588
	const meterPerSecToMilePerHour float64 = 2.237

	type AmericanVelocity float64
	type EuropeanVelocity float64

	val1 := 130.0
	val2 := 120.4

	var convertVal1 AmericanVelocity = AmericanVelocity(val1 * meterPerSecToMilePerHour)
	var convertVal2 EuropeanVelocity = EuropeanVelocity(val2 * meterPerSecToKilometerPerHour)

	fmt.Println("\nЗадание 2")
	fmt.Printf("%.2f м/с = %.2f миль/ч\n", val1, convertVal1)
	fmt.Printf("%.2f м/с = %.2f км/ч\n", val2, convertVal2)
}

func main() {
	Task1()
	Task2()
}
