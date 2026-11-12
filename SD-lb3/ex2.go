package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

type LamportClock struct{ counter int }

func (l *LamportClock) Tick() {
	l.counter++
}

func (l *LamportClock) Send() {
	l.counter++
}

func (l *LamportClock) Receive(ts int) {
	if ts > l.counter {
		l.counter = ts
	}
	l.counter++
}

type Message struct{ sender, ts int }

type Event struct {
	process          int
	internal_counter int
	kind             string
}

func less(a, b Event) bool {
	if a.internal_counter != b.internal_counter {
		return a.internal_counter < b.internal_counter
	}
	return a.process < b.process
}

func worker(id int, in <-chan Message, out chan<- Message, rounds int, eventsCh chan<- Event, wg *sync.WaitGroup) {
	defer wg.Done()
	clock := &LamportClock{counter: 0}

	record := func(kind string) {
		eventsCh <- Event{
			process:          id,
			internal_counter: clock.counter,
			kind:             kind,
		}
	}

	if id == 1 {
		if rand.Intn(2) == 0 {
			clock.Tick()
			record("internal")
		}
		clock.Send()
		record("send")
		out <- Message{sender: id, ts: clock.counter}
	}

	for r := 0; r < rounds; r++ {
		time.Sleep(time.Duration(rand.Intn(15)) * time.Millisecond)
		if rand.Intn(2) == 0 {
			clock.Tick()
			record("internal")
		}

		msg := <-in
		clock.Receive(msg.ts)
		record("recv")

		if rand.Intn(2) == 0 {
			clock.Tick()
			record("internal")
		}

		if id == 1 && r == rounds-1 {
			break
		}

		clock.Send()
		record("send")
		out <- Message{sender: id, ts: clock.counter}
	}
}

func main() {
	ch1to2 := make(chan Message, 1)
	ch2to3 := make(chan Message, 1)
	ch3to1 := make(chan Message, 1)

	eventsCh := make(chan Event, 100)
	var wg sync.WaitGroup

	rounds := 2 
	wg.Add(3)

	go worker(1, ch3to1, ch1to2, rounds, eventsCh, &wg)
	go worker(2, ch1to2, ch2to3, rounds, eventsCh, &wg)
	go worker(3, ch2to3, ch3to1, rounds, eventsCh, &wg)

	wg.Wait()
	close(eventsCh)

	var allEvents []Event
	for ev := range eventsCh {
		allEvents = append(allEvents, ev)
	}

	sort.Slice(allEvents, func(i, j int) bool {
		return less(allEvents[i], allEvents[j])
	})

	fmt.Println("Evenimentele ordonate crescator conform ceasurilor logice Lamport:")
	fmt.Println("-----------------------------------------------------------------")
	for _, ev := range allEvents {
		fmt.Printf("Ceas Lamport: %2d | Proces: P%d | Tip eveniment: %-8s\n",
			ev.internal_counter, ev.process, ev.kind)
	}
}