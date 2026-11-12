package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Clock struct{ offset time.Duration }

func (c *Clock) Now() time.Time { return time.Now().Add(c.offset) }

func estimate(t1, t2, t3, t4 time.Time) (theta, delta time.Duration) {
	delta = t4.Sub(t1) - t3.Sub(t2)

	theta = (t2.Sub(t1) + t3.Sub(t4)) / 2

	return theta, delta
}

func measure(c_clock, s_clock *Clock) (theta, delta time.Duration) {
	t1 := c_clock.Now()
	time.Sleep(3*time.Millisecond + time.Duration(rand.Intn(10))*time.Millisecond)

	t2 := s_clock.Now()
	time.Sleep(5*time.Millisecond + time.Duration(rand.Intn(20))*time.Millisecond)

	t3 := s_clock.Now()

	time.Sleep(3*time.Millisecond + time.Duration(rand.Intn(10))*time.Millisecond)
	t4 := c_clock.Now()

	return estimate(t1, t2, t3, t4)
}

func main() {
	initialOffset := 500 * time.Millisecond
	s_clock := &Clock{offset: 0}
	c_clock := &Clock{offset: initialOffset}

	fmt.Println("=== Subpunctul (a): O singura simulare ===")
	thetaA, deltaA := measure(c_clock, s_clock)
	
	correctedOffsetA := c_clock.offset + thetaA

	fmt.Printf("Offset initial client:   %v\n", initialOffset)
	fmt.Printf("Round-trip (delta):      %v\n", deltaA)
	fmt.Printf("Decalaj estimat (theta): %v\n", thetaA)
	fmt.Printf("Offset dupa ajustare:    %v (a scazut semnificativ fata de %v)\n\n", correctedOffsetA, initialOffset)

	fmt.Println("=== Subpunctul (b): 10 simulari si alegerea delta minim ===")
	c_clock.offset = initialOffset

	var minDelta time.Duration
	var bestTheta time.Duration

	for i := 0; i < 10; i++ {
		theta, delta := measure(c_clock, s_clock)
		fmt.Printf("Iteratia %2d -> delta: %8v | theta: %10v\n", i+1, delta, theta)

		if i == 0 || delta < minDelta {
			minDelta = delta
			bestTheta = theta
		}
	}

	correctedOffsetB := c_clock.offset + bestTheta

	fmt.Println("\n--- Rezultate finale (b) si comparatie cu (a) ---")
	fmt.Printf("Delta minim gasit:       %v (fata de %v la punctul a)\n", minDelta, deltaA)
	fmt.Printf("Theta asociat (optim):   %v (fata de %v la punctul a)\n", bestTheta, thetaA)
	fmt.Printf("Offset rezidual (b):     %v (fata de %v la punctul a)\n", correctedOffsetB, correctedOffsetA)
}