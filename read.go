//package main
package read

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	fname, lname string
}

func check(e error) {
	if e != nil {
		fmt.Println("file not present")
		panic(e)
	}
}

func main() {
	arr := make([]Name, 0)
	var filename string

	fmt.Println("Enter the filename (Please put file in the same directory):")
	fmt.Scan(&filename)

	file, err := os.Open(filename)
	check(err)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		name := strings.Split(scanner.Text(), " ")
		arr = append(arr, Name{fname: name[0], lname: name[1]})
	}

	for i, name := range arr {
		fmt.Printf("(Line %d) First Name: %s, Last Name: %s\n", i, name.fname, name.lname)
	}

	file.Close()
}
