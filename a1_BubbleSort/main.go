// Bubble Sort Implementation in Go
package main

import "fmt"

func bubble_sort(dataset []int) {
	for i := 1; i <= len(dataset)-1; i++ {

		if dataset[i] >= dataset[i-1] {
			continue
		}

		var y int = dataset[i]
		var j int = i - 1

		for {
			if j < 0 || dataset[j] <= y {
				break
			}
			dataset[j+1] = dataset[j]
			j--
		}

		dataset[j+1] = y
	}
}

func main() {

	dataset := []int{5, 2, 4, 6, 1, 3, 15, 6, 23645, 23, 462, 2, 2, -12}

	fmt.Println("Lenghth of the Array: \t", len(dataset))

	fmt.Println(dataset)

	bubble_sort(dataset)

	fmt.Println(dataset)
}
