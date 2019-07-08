//package main
package makejson

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	m := make(map[string]string)
	var name1, add1 string

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter a name")
	fmt.Scan(&name1)
	fmt.Println("Enter the address")
	scanner.Scan()
	add1 = scanner.Text()

	m["name"] = name1
	m["address"] = add1

	barr, err := json.Marshal(m)

	if err != nil {
		os.Exit(1)
	}

	fmt.Println(string(barr))
}
