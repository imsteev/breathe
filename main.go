package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	box := flag.Bool("box", false, "box breathe")
	fourSevenEight := flag.Bool("478", false, "four-seven-eight breathing")
	coherent := flag.Bool("coherent", false, "coherent breathing")
	flag.Parse()

	if *box {
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
	}
	if *fourSevenEight {
		i := 0
		tick := time.NewTicker(1 * time.Second)
		for range tick.C {
			rem := i % 19
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
	}
	if *coherent {
		i := 0
		tick := time.NewTicker(1 * time.Second)
		for range tick.C {
			rem := i % 10
			if rem < 5 {
				fmt.Print("IN ")
			} else {
				fmt.Print("OUT ")
			}
			i++
			if i%10 == 0 {
				fmt.Printf("(end of round %d)\n\n", i/10)
			}
		}
	}
}
