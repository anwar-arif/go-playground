package one_billion_rows

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const PropertyNameMaxLen = 7

func generateRandomName() string {
	result := make([]byte, PropertyNameMaxLen)

	for i := 0; i < PropertyNameMaxLen; i++ {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

const MaxRandomNum float64 = 999.00
const MinRandomNum float64 = 100.00
const Decimals = 2

func randomFloatWithPrecision() float64 {
	random := MinRandomNum + rand.Float64()*(MaxRandomNum-MinRandomNum)
	factor := math.Pow(10, float64(Decimals))
	return math.Round(random*factor) / factor
}

type Property struct {
	Name  string
	Price float64
}

func getProperty() Property {
	return Property{
		Name:  generateRandomName(),
		Price: randomFloatWithPrecision(),
	}
}

const TotalEntries = 1000000000

var filename string

func init() {
	rootDir, _ := os.Getwd()
	absPath := "cmd/playground/one_billion_rows/data.txt"
	filename = rootDir + "/" + absPath
}

func PopulateData() {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for i := 0; i < TotalEntries; i++ {
		property := getProperty()
		content := fmt.Sprintf("%s;%f\n", property.Name, property.Price)
		if _, err := writer.WriteString(content); err != nil {
			panic(err)
		}
		writer.Flush()
	}

}
