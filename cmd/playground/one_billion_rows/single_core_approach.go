package one_billion_rows

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stat struct {
	Min   float64
	Max   float64
	Sum   float64
	Count int64
}

func ProcessFileSingleCore(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	statMap := make(map[string]*Stat)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ";")
		if len(parts) < 2 {
			continue
		}

		propertyID := parts[0]
		price, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			continue
		}

		current, exists := statMap[propertyID]
		if !exists {
			statMap[propertyID] = &Stat{
				Min:   price,
				Max:   price,
				Sum:   price,
				Count: 1,
			}
		} else {
			if price < current.Min {
				current.Min = price
			}
			if price > current.Max {
				current.Max = price
			}
			current.Sum += price
			current.Count++
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Print("{")
	for prop, stat := range statMap {
		avg := stat.Sum / float64(stat.Count)
		fmt.Printf("%s=%.1f/%.1f/%.1f, ", prop, stat.Min, avg, stat.Max)
	}
	fmt.Println("\b\b}")
}

func RunSingleCoreApproach() {
	ProcessFileSingleCore(filename)
}
