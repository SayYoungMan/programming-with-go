package main

import (
	"fmt"
	"sync"
	"time"
)

type Chops struct{ sync.Mutex }

type Philo struct {
	number          int
	leftCS, rightCS *Chops
}

func (p Philo) eat(c chan *Philo, wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		c <- &p

		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Printf("Philosopher %v started eating %v time.\n", p.number, i+1)
		fmt.Printf("Philosopher %v finished eating %v time.\n", p.number, i+1)

		p.leftCS.Unlock()
		p.rightCS.Unlock()
	}
	wg.Done()
}

type Host struct {
	philChan chan *Philo
	stopChan chan struct{}
	wg       *sync.WaitGroup
}

func newHost(philChan chan *Philo, stopChan chan struct{}, wg *sync.WaitGroup) *Host {
	return &Host{philChan, stopChan, wg}
}

func (h *Host) run() {
	fmt.Println("Host has arrived.")
	for {
		select {
		case <-h.philChan:
			if len(h.philChan) == 2 {
				<-h.philChan
				<-h.philChan
				time.Sleep(time.Second)
			}
		case <-h.stopChan:
			break
		}
	}
}

func main() {
	var wg sync.WaitGroup

	CSticks := make([]*Chops, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(Chops)
	}

	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{i + 1, CSticks[i], CSticks[(i+1)%5]}
	}

	philChan := make(chan *Philo, 2)
	stopChan := make(chan struct{})

	host := newHost(philChan, stopChan, &wg)
	go host.run()

	time.Sleep(time.Second)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philos[i].eat(philChan, &wg)
	}
	wg.Wait()

	stopChan <- struct{}{}
	fmt.Println("Done!")
}
