package main

import (
	"fmt"

	"github.com/zetamatta/go-getch"
)

func main() {
	fmt.Print("Press any key\n")
	_ = getch.Rune()
	fmt.Print("ok\n")
}

func clear() {
	for i := 0; i < 100; i++ {
		fmt.Println()
	}
}
