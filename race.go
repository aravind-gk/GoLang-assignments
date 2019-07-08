/*
* The following program is an example of a race-condition implemented in GO
*
* A race condition is said to be present, when the output of two or more processes
* running concurrently possilbly produce different outputs under execution, which
* depends on the order of interleaving of those concurrent routines.
*
* In golang, the keyword go is used to create a new goroutine
*
* In this program, apart from the main goroutine, two more goroutines are created.
* Both goroutines run the same function(print the parameter string) but with different
* parameters. And the main goroutine(parent goroutine) waits for the other two
* goroutines to complete execution using sync.Waitgroup
*
* Both of the new created goroutines run concurrently, thus are allowed to interleave.
* This can be seen when the program is executed. The string 'hello' is given to 1st
* goroutine and 'bye' is given to other goroutine. While executing, we can see that the
* output of the main program is a random sequence of hello's and bye's, which may not
* be the same sequence if the program is run again. This is the race condition.
 */

//package main
package race

import (
	"fmt"
	"sync"
	"time"
)

func say(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 50; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go say("hello", &wg) //first goroutine
	go say("bye", &wg)   //second goroutine

	wg.Wait()
	fmt.Println("Main goroutine exiting.")
}
