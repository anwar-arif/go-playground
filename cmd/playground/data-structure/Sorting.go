package data_structure

import (
	"fmt"
	"sort"
	"strings"
)

// Person struct for custom sorting examples
type Person struct {
	Name string
	Age  int
	City string
}

// Custom type for implementing sort.Interface
type ByAge []Person
type ByName []Person
type ByMultiField []Person

// Required methods for sort.Interface for ByAge
func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

// Required methods for sort.Interface for ByName
func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

// Required methods for sort.Interface for ByMultiField
func (a ByMultiField) Len() int      { return len(a) }
func (a ByMultiField) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByMultiField) Less(i, j int) bool {
	// First compare by age
	if a[i].Age != a[j].Age {
		return a[i].Age < a[j].Age
	}
	// If ages are equal, compare by name
	if a[i].Name != a[j].Name {
		return a[i].Name < a[j].Name
	}
	// If names are equal, compare by city
	return a[i].City < a[j].City
}

func RunSorting() {
	// 1. Sorting built-in types
	fmt.Println("1. Sorting built-in types:")

	// Sorting integers
	numbers := []int{5, 2, 6, 3, 1, 4}
	sort.Ints(numbers)
	fmt.Printf("Sorted integers: %v\n", numbers)

	// Sorting strings
	fruits := []string{"banana", "apple", "orange", "grape"}
	sort.Strings(fruits)
	fmt.Printf("Sorted strings: %v\n", fruits)

	// Sorting float64s
	floats := []float64{3.14, 1.41, 2.71, 1.73}
	sort.Float64s(floats)
	fmt.Printf("Sorted floats: %v\n", floats)

	// 2. Custom sorting with sort.Interface
	fmt.Println("\n2. Custom sorting with sort.Interface:")

	people := []Person{
		{"Alice", 25, "New York"},
		{"Bob", 30, "Chicago"},
		{"Charlie", 20, "Boston"},
		{"David", 25, "Miami"},
	}

	// Sort by age
	sort.Sort(ByAge(people))
	fmt.Printf("Sorted by age: %v\n", people)

	// Sort by name
	sort.Sort(ByName(people))
	fmt.Printf("Sorted by name: %v\n", people)

	// Sort by multiple fields
	sort.Sort(ByMultiField(people))
	fmt.Printf("Sorted by multiple fields: %v\n", people)

	// 3. Sorting with sort.Slice
	fmt.Println("\n3. Sorting with sort.Slice:")

	// Sort by age using sort.Slice
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Printf("Sorted by age (sort.Slice): %v\n", people)

	// Sort by multiple fields using sort.Slice
	sort.Slice(people, func(i, j int) bool {
		if people[i].Age != people[j].Age {
			return people[i].Age < people[j].Age
		}
		return people[i].Name < people[j].Name
	})
	fmt.Printf("Sorted by age and name (sort.Slice): %v\n", people)

	// 4. Reverse sorting
	fmt.Println("\n4. Reverse sorting:")

	// Reverse sort integers
	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
	fmt.Printf("Reverse sorted integers: %v\n", numbers)

	// Reverse sort custom type
	sort.Sort(sort.Reverse(ByAge(people)))
	fmt.Printf("Reverse sorted by age: %v\n", people)

	// 5. Case-insensitive string sorting
	fmt.Println("\n5. Case-insensitive string sorting:")

	mixedCaseStrings := []string{"banana", "Apple", "orange", "Grape"}
	sort.Slice(mixedCaseStrings, func(i, j int) bool {
		return strings.ToLower(mixedCaseStrings[i]) < strings.ToLower(mixedCaseStrings[j])
	})
	fmt.Printf("Case-insensitive sorted strings: %v\n", mixedCaseStrings)

	// 6. Stable sorting
	fmt.Println("\n6. Stable sorting:")

	data := []Person{
		{"Alice", 25, "New York"},
		{"Bob", 25, "Chicago"},
		{"Charlie", 25, "Boston"},
	}
	sort.SliceStable(data, func(i, j int) bool {
		return data[i].City < data[j].City
	})
	fmt.Printf("Stable sorted by city: %v\n", data)
}
