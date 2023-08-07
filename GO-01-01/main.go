package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("Hello! Текущее время: %s", time.Now().Format("02.01.2006 15:04"))
}
