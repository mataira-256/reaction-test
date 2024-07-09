package main

import (
	"fmt"
	"time"

	"github.com/zetamatta/go-getch"
)

func main() {
	var times int
	var delaySum float64

	fmt.Print("Number of times:")
	fmt.Scan(&times)

	for i := 0; i < times; i++ {
		clear()
		delaySum += delay_measure(i + 1)
	}

	fmt.Printf("ave delay: %.3fsec", delaySum/float64(times))
}

func delay_measure(n int) float64 {
	/* start */
	fmt.Printf("%d time\n", n)
	fmt.Print("Press any key\n")
	_ = getch.Rune()

	/* count down */
	const (
		NANO_SEC = 1000000000
		CD_SEC   = 3
	)
	cdStart := time.Now()
	elapsed := time.Since(cdStart)
	for elapsed < CD_SEC*NANO_SEC {
		fmt.Printf("%.1fsec\n", CD_SEC-float64(elapsed)/float64(NANO_SEC))
		elapsed = time.Since(cdStart)
	}

	/* measure */
	msStart := time.Now()
	_ = getch.Rune()
	return time.Since(msStart).Seconds()
}

func clear() {
	for i := 0; i < 100; i++ {
		fmt.Println()
	}
}
