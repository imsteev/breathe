package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	fmt.Fprint(os.Stdout, "Press any key to start breathing\n")

	// bug: this only works if newline is entered?
    _, err := fmt.Scanf("%c", new(rune)) // do not care about the input
    if err != nil {
        log.Fatal(err)
    }

	var t *time.Ticker
	if false {
		t = boxBreathing()
	} else {
		t = fourSevenEightBreathing()
	}
	defer t.Stop()

	for {}
}

// caller responsible for stopping ticker
func boxBreathing() *time.Ticker {
	i := 0
	tick := time.NewTicker(1000 * time.Millisecond)
	for range tick.C {
		if i%16 < 4 {
			fmt.Print("IN (4) ")
		} else if i%16 < 8 {
			fmt.Print("HOLD (4) ")
		} else if i%16 < 12 {
			fmt.Print("OUT (4) ")
		} else if i%16 < 16 {
			fmt.Print("HOLD (4) ")
		}
		i++
		if i%16 == 0 {
			fmt.Printf("(end of round %d)\n\n", i/16)
		}
	}
	return tick
}

// caller responsible for stopping ticker
func fourSevenEightBreathing() *time.Ticker {
	i := 0
	tick := time.NewTicker(1 * time.Second)
	for range tick.C {
		rem := i%19
		if rem < 4 {
			fmt.Print("IN ")
		} else if rem < 11 {
			fmt.Print("HOLD ")
		} else if rem < 19 {
			fmt.Print("OUT ")
		}
		i++
		if i%19 == 0 {
			fmt.Printf("(end of round %d)\n\n", i/19)
		}
	}
	return tick	
}
