// package main
package semaphore

type semaphore chan struct{}

func (s semaphore) Acquire(buffers int) {
	var e struct{}
	for buffer := 0; buffer < buffers; buffer++ {
		s <- e
	}
}

func (s semaphore) Release(buffers int) {
	for buffer := 0; buffer < buffers; buffer++ {
		<-s
	}
}
