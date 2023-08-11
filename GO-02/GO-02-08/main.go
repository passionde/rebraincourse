package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Работа с файловым дескриптором
	fileIn, err := os.Open("GO-02-08/data/in.txt")
	if err != nil {
		log.Fatal("error opening file:", err)
	}

	defer func(fileIn *os.File) {
		err := fileIn.Close()
		if err != nil {
			log.Fatal("error closing file:", err)
		}
	}(fileIn)

	// Подсчет строк. Лучше реализовать через bufio.NewScanner, но в задании нужно обработать EOF
	counter := 1
	r := bufio.NewReader(fileIn)
	_, err = r.ReadString('\n')

	for !errors.As(err, &io.EOF) {
		counter++
		_, err = r.ReadString('\n')
	}

	fmt.Printf("Total strings: %d", counter)
}
