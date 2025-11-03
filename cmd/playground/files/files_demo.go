package files

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadFromFile() {
	file, err := os.Open("./files/input.txt")
	if err != nil {
		log.Printf("cannot open file: %v\n", err.Error())
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Printf("%v\n", err.Error())
	}
}

func WriteToFile() {
	file, err := os.OpenFile("./files/output.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("cannot open file %v\n", err.Error())
		return
	}
	defer file.Close()
	if _, err := file.WriteString("Welcome to Go Programming\n"); err != nil {
		log.Printf("error writing in the file: %v\n", err.Error())
		return
	}
}

func RunFilesDemo() {
	ReadFromFile()
	WriteToFile()
}
