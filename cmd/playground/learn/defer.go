package learn

import (
	"fmt"
	"log"
	"os"
)

func writeToTempFile(text string) error {
	file, err := os.Open("temp.txt")
	if err != nil {
		return err
	}
	defer file.Close()
	n, err := file.WriteString("some text")
	if err != nil {
		return err
	}

	fmt.Printf("number of bytes written: %d\n", n)
	return nil
}

func RunDefer() {
	if err := writeToTempFile("some text"); err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("write to file successfull")
}
