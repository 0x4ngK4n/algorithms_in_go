package main

import "fmt"

func selection_sort(sliceToSort []int) []int {
	for i := range sliceToSort {
		sliceToCompare := sliceToSort[i+1:]
		for j := range sliceToCompare {
			if sliceToCompare[j] < sliceToSort[i] {
				sliceToCompare[j], sliceToSort[i] = sliceToSort[i], sliceToCompare[j]
			}
		}
	}
	return sliceToSort
}

func main() {
	sliceToSort := []int{18, 6, 3, 4, 9, 2, 17, 15, 23}
	fmt.Println("slice to sort: ", sliceToSort)
	fmt.Println("selection sorted array: ", selection_sort(sliceToSort))
}
