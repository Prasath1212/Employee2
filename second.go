package main

import "fmt"

func findSecondLargest(arr []int) int {
	largest := arr[0]
	secondLargest := -1

	for i := 1; i < len(arr); i++ {
		if arr[i] > largest {
			secondLargest = largest
			largest = arr[i]
		} else if arr[i] > secondLargest && arr[i] != largest {
			secondLargest = arr[i]
		}
	}

	return secondLargest
}

func main() {
	arr := []int{45, 24, 10, 10, 42, 12}
	secondLargest := findSecondLargest(arr)
	fmt.Println("The second largest element is:", secondLargest)
}
