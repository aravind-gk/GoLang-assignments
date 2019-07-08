//package main
package trunc

import "fmt"

func main() {
	fmt.Printf("Enter a floating point number\n")
	var x float64
	fmt.Scan(&x)
	y := int64(x)
	fmt.Printf("Truncated number: %d\n", y)
}
