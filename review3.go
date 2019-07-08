package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type ChopS struct{ sync.Mutex }

type Philo struct {
	CS        []*ChopS
	phili_num int
	wg        *sync.WaitGroup
}

func (p Philo) eat(host chan int) {
	for i := 0; i < 3; i++ {

		<-host
		lock_1 := rand.Intn(2)

		p.CS[lock_1].Lock()
		p.CS[1-lock_1].Lock()

		fmt.Printf("starting to eat %d\n", p.phili_num)
		fmt.Printf("finishing eating %d\n", p.phili_num)

		p.CS[0].Unlock()
		p.CS[1].Unlock()

		host <- 1
	}
	p.wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	rand.Seed(time.Now().Unix())

	host := make(chan int, 2)
	host <- 1
	host <- 1

	CSticks := make([]*ChopS, 5)

	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	philos := make([]*Philo, 5)

	for i := 0; i < 5; i++ {
		philos[i] = &Philo{[]*ChopS{CSticks[i], CSticks[(i+1)%5]}, i + 1, &wg}
	}

	for i := 0; i < 5; i++ {
		go philos[i].eat(host)
	}
	wg.Wait()
}
