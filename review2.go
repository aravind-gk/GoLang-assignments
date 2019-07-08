//package main
package review2

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	slice := make([]int, 0)
	fmt.Println("Enter integers one at a time")
	fmt.Println("Enter s to sort")

	var input string

	for {
		fmt.Printf(">")
		scanner := bufio.NewScanner(os.Stdin)

		if scanner.Scan() {
			input = scanner.Text()

			if input == "s" {
				fmt.Printf("\nInput: %v\n\n", slice)
				splitAndSort(&slice, 4)
				return
			}

			num, err := strconv.ParseInt(input, 10, 64)

			if err != nil {
				fmt.Println(err)
				return
			}

			slice = append(slice, int(num))
		}
	}
}

func splitAndSort(slice *[]int, partitions int) {
	var chunk int
	n := len(*slice)

	if n%partitions == 0 {
		chunk = n / partitions
	} else {
		chunk = n/partitions + 1
	}

	c := make(chan []int, 4)
	finalArray := make([]int, 0)

	for i := 0; i < n; i += chunk {
		end := i + chunk

		if end > n {
			end = n
		}
		go sortSlice((*slice)[i:end], c)
		finalArray = append(finalArray, <-c...)
	}

	sort.Ints(finalArray)
	fmt.Printf("\nSorted: %v\n", finalArray)
}

func sortSlice(slice []int, c chan []int) {
	fmt.Printf("Sorting %v\n", slice)
	sort.Ints(slice)
	c <- slice
}
