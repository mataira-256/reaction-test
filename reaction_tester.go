package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/zetamatta/go-getch"
)

func main() {
	var reactionTime []float64
	var reactionSum float64
	var delayTime float64
	var times int

	fmt.Print("delay time(sec):")
	fmt.Scan(&delayTime)
	fmt.Print("Number of times:")
	fmt.Scan(&times)

	for i := 0; i < times; i++ {
		clear()
		reactionTime = append(reactionTime, reaction_tester(i+1)-delayTime)
		reactionSum += reactionTime[i]
	}

	clear()
	for i := 0; i < times; i++ {
		fmt.Printf("score %d: %.3fsec\n", i+1, reactionTime[i])
	}
	fmt.Printf("ave score: %.3fsec\n", reactionSum/float64(times))
	fmt.Printf("delay time: %.3fsec", delayTime)
}

func reaction_tester(n int) float64 {
	/* start */
	fmt.Printf("%d time\n", n)
	fmt.Print("Press any key\n")
	_ = getch.Rune()

	/* wait */
	const (
		MAX_SEC = 8
		MIN_SEC = 2
	)
	clear()
	waitTime := (rand.Float64()*(MAX_SEC-MIN_SEC) + MIN_SEC) * 1000
	fmt.Printf("wait...")
	time.Sleep(time.Duration(waitTime) * time.Millisecond)

	/* push */
	clear()
	fmt.Print("push!!")
	msStart := time.Now()
	_ = getch.Rune()
	return time.Since(msStart).Seconds()
}

func clear() {
	for i := 0; i < 100; i++ {
		fmt.Println()
	}
}
