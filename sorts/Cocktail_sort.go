package main

import "fmt"

func CocktailSort(items []int) []int {
	swapped := true
	start := 0
	end := len(items) - 1

	for swapped == true {
		swapped = false

		for i := start; i < end-1; i++ {
			if items[i] > items[i+1] {
				items[i], items[i+1] = items[i+1], items[i]
				swapped = true
			}
		}

		if swapped == false {
			break
		}

		swapped = false

		end -= 1

		for i := end; i > start-1; i-- {
			if items[i] > items[i+1] {
				items[i], items[i+1] = items[i+1], items[i]
				swapped = true
			}

		}

		start = start + 1
	}

	return items
}

func main() {

	a := []int{5, 1, 4, 2, 8, 10, 2, 0}

	fmt.Println(CocktailSort(a))

}
