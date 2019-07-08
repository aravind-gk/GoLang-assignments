///package main
package review

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	rstdin := bufio.NewReader(os.Stdin)
	var sli []int
	fmt.Println("Enter a sequence up to 10 numbers (empty to stop): ")
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d. Enter a number: ", i)
		elem, _ := rstdin.ReadString('\n')
		elem = strings.TrimSpace(elem)
		elemInt, err := strconv.Atoi(elem)
		if err != nil {
			break
		}
		sli = append(sli, elemInt)
	}
	fmt.Println("The final array is: ", sli)
	BubbleSort(sli)
	fmt.Println("The sorted array is: ", sli)
}

func BubbleSort(sli []int) {
	for i := 1; i < len(sli); i++ {
		for j := 0; j < len(sli)-i; j++ {
			if sli[j] > sli[j+1] {
				Swap(sli, j)
			}
		}
	}
}

func Swap(sli []int, index int) {
	aux := sli[index]
	sli[index] = sli[index+1]
	sli[index+1] = aux
}
