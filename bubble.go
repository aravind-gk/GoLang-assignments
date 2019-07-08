//package main
package bubble

import (
	"fmt"
)

func Swap(arr []int, index int) {
	arr[index], arr[index+1] = arr[index+1], arr[index]
}

func BubbleSort(arr []int) {

	for i := 0; i < len(arr)-1; i = i + 1 {
		for j := 0; j < len(arr)-1-i; j = j + 1 {
			if arr[j] > arr[j+1] {
				Swap(arr, j)
			}
		}
	}
}

func main() {
	var size, num int
	var arr []int
	fmt.Println("Enter the size of the array (between 1 and 10):")
	fmt.Scan(&size)

	if size < 1 || size > 10 {
		panic("Size should be 1 to 10.")
	}

	fmt.Println("Now enter the elements of the array:")

	for i := 0; i < size; i++ {
		fmt.Scan(&num)
		arr = append(arr, num)
	}

	BubbleSort(arr)

	fmt.Println("Sorted array:", arr)
}
