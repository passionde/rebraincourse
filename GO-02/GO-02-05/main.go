package main

import "fmt"

func contains(a []string, x string) bool {
	if len(a) == 0 {
		return false
	}

	for _, v := range a {
		if v == x {
			return true
		}
	}

	return false
}

func getMax(arr ...int) int {
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	strSlice := []string{"foo", "bar"}

	fmt.Printf("%s in %v -> %t\n", "foo", strSlice, contains(strSlice, "foo"))
	fmt.Printf("%s in %v -> %t\n", "undefined", strSlice, contains(strSlice, "undefined"))

	intSlice := []int{-1, 4, -4, 9, 10, -6}
	fmt.Printf("Find max value in %v -> %d\n", intSlice, getMax(intSlice...))
}
