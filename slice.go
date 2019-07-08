//package main
package slice

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	arr := make([]int, 0, 3)
	var x string
	for {
		fmt.Println("Enter a number:")
		fmt.Scan(&x)
		num, err := strconv.Atoi(x)
		if err == nil {
			arr = append(arr, num)
			sort.Ints(arr)
			fmt.Printf("Sorted array: %v\n\n", arr)
		} else {
			if x == "X" || x == "x" {
				fmt.Println("Exiting the loop")
				break
			} else {
				fmt.Println("Error, enter a valid number or enter 'X' to quit")
			}
		}
	}
}
