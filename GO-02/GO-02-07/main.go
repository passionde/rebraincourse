package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func ParseRowFunc() func(string2 string) string {
	var counter int
	return func(row string) string {
		counter++

		fields := strings.Split(row, "|")
		if len(fields) != 3 {
			panic(fmt.Sprintf("parse error: empty field on string %d", counter))
		}

		return fmt.Sprintf(
			"Row: %d\nName: %s\nAddress: %s\nCity: %s\n\n\n",
			counter, fields[0], fields[1], fields[2],
		)

	}
}

func OutputFileContent(filename string) {
	fileOut, _ := os.Open(filename)
	s := bufio.NewScanner(fileOut)
	for s.Scan() {
		fmt.Println(s.Text())
	}
}

func HandlePanic() {
	err := recover()
	if err == nil {
		return
	}

	msg, ok := err.(string)
	if !ok {
		log.Fatal("Unknown panic data type")
	}
	if strings.HasPrefix(msg, "parse error") {
		OutputFileContent("GO-02-07/data/data_out.txt")
	}
}

func main() {
	fileIn, err := os.Open("GO-02-07/data/in.txt")
	if err != nil {
		log.Fatal("error opening data/in.txt file:", err)
	}

	fileOut, err := os.OpenFile("GO-02-07/data/data_out.txt", os.O_CREATE, 0666)
	if err != nil {
		log.Fatal("error opening data/data_out.txt file:", err)
	}

	defer HandlePanic()
	defer fileIn.Close()
	defer fileOut.Close()

	parseRow := ParseRowFunc()
	s := bufio.NewScanner(fileIn)

	for s.Scan() {
		_, err = fileOut.WriteString(parseRow(s.Text()))
		if err != nil {
			log.Fatal("error write in file:", err)
		}
	}
}
