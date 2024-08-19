package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	var (
		box            = flag.Bool("box", false, "Box breathing: in 4s, hold 4s, out 4s, hold 4s")
		fourSevenEight = flag.Bool("478", false, "4-7-8 breathing: in 4s, hold 7s, out 8s. Triggers parasympathetic nervous system (aka, good for relaxation)")
		coherent       = flag.Bool("coh", false, "Coherent breathing: in 5.5s, out 5.5s. Try to make it as circular as possible")
	)

	flag.Parse()

	if *box {
		sequence := []string{"IN", ".", ".", ".", "HOLD", ".", ".", ".", "OUT", ".", ".", ".", "HOLD", ".", ".", "."}

		round := 1
		tick := time.NewTicker(1000 * time.Millisecond)
		for {
			for _, step := range sequence {
				<-tick.C
				fmt.Printf("%s ", step)
			}
			fmt.Printf("(end of round %d)\n\n", round)
			round++
		}
	}
	if *fourSevenEight {
		sequence := []string{"IN", ".", ".", ".", "HOLD", ".", ".", ".", ".", ".", ".", "OUT", ".", ".", ".", ".", ".", ".", "."}

		round := 1
		tick := time.NewTicker(1000 * time.Millisecond)
		for {
			for _, step := range sequence {
				<-tick.C
				fmt.Printf("%s ", step)
			}
			fmt.Printf("(end of round %d)\n\n", round)
			round++
		}
	}
	if *coherent {
		sequence := []string{"IN", ".", ".", ".", ".", "OUT", ".", ".", ".", "."}

		round := 1
		tick := time.NewTicker(1100 * time.Millisecond)
		for {
			for _, step := range sequence {
				<-tick.C
				fmt.Printf("%s ", step)
			}
			fmt.Printf("(end of round %d)\n\n", round)
			round++
		}
	}
}
