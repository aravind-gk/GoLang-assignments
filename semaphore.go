// package main
package semaphore

type empty struct{}
type semaphore chan empty

func (s semaphore) P(n int) {
	e := empty{}
	for i := 0; i < n; i++ {
		s <- e
	}
}

func (s semaphore) V(n int) {
	for i := 0; i < n; i++ {
		<-s
	}
}

// func main() {
// 	sem := make(semaphore, N);
// }
