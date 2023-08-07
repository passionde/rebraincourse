package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	// Открытие файлов
	fileIn, _ := os.Open("GO-02-06/data/in.txt")
	fileOut, _ := os.OpenFile("GO-02-06/data/out.txt", os.O_CREATE, 0666)

	defer fileIn.Close()
	defer fileOut.Close()

	// Вызов logTime
	startTime := time.Now()
	logTime := func() {
		fmt.Printf("Time elapsed: %s", time.Since(startTime).String())
	}
	defer logTime()

	// Обработка данных файлов
	amountLine := 0
	amountByte := 0

	s := bufio.NewScanner(fileIn)
	for i := 1; s.Scan(); i++ {
		n, _ := fileOut.WriteString(fmt.Sprintf("%4d. %s\n", i, s.Text()))
		amountLine++
		amountByte += n
	}

	// Вывод результатов
	fmt.Printf("Program finish:\n * Total line: %d\n * Total byte: %d\n", amountLine, amountByte)
}
