/*
 * The below program is an implementation of the Dining Philosophers problem, with the following features:
 *
 * 1. There are 5 philosophers, sharing 5 chopsticks (two on each side of philosopher).
 * 2. Each philosopher eats only 3 times.
 * 3. The philosphers pick either the leftChopstick first or the rightChopstick first, depending on the value
 *	  of a random value generator(so that both orders of picking up chopsticks are possible). This decision of
 *    the order of picking csticks is taken after acquiring the semaphore lock first.
 * 4. The main goroutine itself acts as the host, which uses a shared 'counting-semaphore' variable initialized
 *    to 2, so that a maximum of only two philosphers can eat at a time.
 * 5. The counting-semaphore is implemented using channels by using struct{} with P() function for blocking
*     goroutine and V() funtion for releasing goroutine.
 * 6. Each philosopher is numbered through 1 to 5.
 * 7. When philosopher starts eating after obtaining all necessary locks, it prints "starting to eat <philID>"
 * 8. When philosopher finishes eating before releasing any locks, it prints "finishing eating <philID>"
*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/* counting semaphore implementation using channels and struct{} */
type empty struct{}
type semaphore chan empty

/* decrease the counting semaphore (blocking) */
func (s semaphore) P(n int) {
	e := empty{}
	for i := 0; i < n; i++ {
		s <- e
	}
}

/* increasing the counting semaphore (releasing) */
func (s semaphore) V(n int) {
	for i := 0; i < n; i++ {
		<-s
	}
}

/* struct for chopsticks */
type ChopS struct {
	sync.Mutex
}

/* struct for philosophers */
type Philo struct {
	philId  int
	leftCS  *ChopS
	rightCS *ChopS
}

func (p Philo) TryToEat(wg *sync.WaitGroup, sem *semaphore) {
	defer wg.Done() /* after philospher eats 3 times, signal the main goroutine */
	for i := 0; i < 3; i++ {
		sem.P(1) /* take permission for eating (only 2 can eat at a time, so wait till then) */

		rand.Seed(time.Now().UnixNano()) /* generate random number based on system time */

		/* either leftCS is chosen first or rightCS is chosen first depending on the random value, then eat */
		if r := rand.Intn(2); r == 0 {
			p.PickLeftCSthenRightCSthenEat()
		} else {
			p.PickRightCSthenLeftCSthenEat()
		}

		sem.V(1) /* increment the counting semaphore when finished eating once */
	}
}

/* philosopher eating by acquiring lock for right chopstick first */
func (p Philo) PickRightCSthenLeftCSthenEat() {
	p.rightCS.Lock()
	p.leftCS.Lock()

	fmt.Println("starting to eat ", p.philId)
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("finishing eating ", p.philId)

	p.leftCS.Unlock()
	p.rightCS.Unlock()
}

/* philosopher eating by acquiring lock for left chopstick first */
func (p Philo) PickLeftCSthenRightCSthenEat() {
	p.leftCS.Lock()
	p.rightCS.Lock()

	fmt.Println("starting to eat ", p.philId)
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("finishing eating ", p.philId)

	p.rightCS.Unlock()
	p.leftCS.Unlock()
}

func main() {
	/* create 5 new chopsticks */
	Csticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		Csticks[i] = new(ChopS)
	}

	/* create 5 new philosophers and assign each with 2 chopsticks */
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{i + 1, Csticks[i], Csticks[(i+1)%5]}
	}

	var wg sync.WaitGroup /* to make the main goroutine wait for other goroutines to complete */
	wg.Add(5)

	/* initialize the counting semaphore to 2 so that maximum of only 2 phils can eat at a time */
	sem := make(semaphore, 2)

	/* all philosophers now try to eat */
	for i := 0; i < 5; i++ {
		go philos[i].TryToEat(&wg, &sem)
	}

	/* wait for all phils to finish eating 3 times each */
	wg.Wait()
}
