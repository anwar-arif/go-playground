package hackerrank

import "sort"

/*
 * Complete the 'customSorting' function below.
 *
 * The function is expected to return a STRING_ARRAY.
 * The function accepts STRING_ARRAY strArr as parameter.
 */

func customSorting(strArr []string) []string {
	even := make([]string, 0)
	odd := make([]string, 0)

	for i := 0; i < len(strArr); i++ {
		val := strArr[i]
		if len(val)%2 == 0 {
			even = append(even, val)
		} else {
			odd = append(odd, val)
		}
	}

	sort.SliceStable(even, func(i, j int) bool {
		if len(even[i]) != len(even[j]) {
			return len(even[i]) > len(even[j])
		}
		return even[i] < even[j]
	})
	sort.SliceStable(odd, func(i, j int) bool {
		if len(odd[i]) != len(odd[j]) {
			return len(odd[i]) < len(odd[j])
		}
		return odd[i] < odd[j]
	})

	for i, val := range odd {
		strArr[i] = val
	}
	for i, val := range even {
		strArr[i+len(odd)] = val
	}

	return strArr
}

func isOdd(x string) bool {
	return len(x)%2 == 1
}
func isEven(x string) bool {
	return len(x)%2 == 0
}

func customSortingV2(strArr []string) []string {
	arr := make([]string, 0)
	arr = append(arr, strArr...)

	sort.SliceStable(arr, func(i, j int) bool {
		first := arr[i]
		second := arr[j]
		if isOdd(first) == isOdd(second) {
			if len(first) == len(second) {
				return first < second
			}
			return len(first) < len(second)
		} else if isEven(first) == isEven(second) {
			if len(first) == len(second) {
				return first < second
			}
			return len(first) < len(second)
		}
		return isOdd(first)
	})

	for i, val := range arr {
		strArr[i] = val
	}

	return strArr
}
