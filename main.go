package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/imsteev/breathe/countdown"
)

func main() {
	var (
		box            = flag.Bool("box", false, "Box breathing: in 4s, hold 4s, out 4s, hold 4s")
		fourSevenEight = flag.Bool("478", false, "4-7-8 breathing: in 4s, hold 7s, out 8s. Triggers parasympathetic nervous system (aka, good for relaxation)")
		coherent       = flag.Bool("coh", false, "Coherent breathing: in 5.5s, out 5.5s. Try to make it as circular as possible")
	)

	flag.Parse()

	if *box {
		var (
			firstIn   = countdown.New(4)
			firstHold = countdown.New(4)
			secIn     = countdown.New(4)
			secHold   = countdown.New(4)
		)

		round := 1
		tick := time.NewTicker(1000 * time.Millisecond)
		for {
			for range tick.C {
				if firstIn.Next() {
					fmt.Print("IN ")
				} else if firstHold.Next() {
					fmt.Print("HOLD ")
				} else if secIn.Next() {
					fmt.Print("OUT ")
				} else if secHold.Next() {
					fmt.Print("HOLD ")
				} else {
					break
				}
			}
			fmt.Printf("(end of round %d)\n\n", round)
			firstIn.Reset()
			firstHold.Reset()
			secIn.Reset()
			secHold.Reset()
			round++
		}
	}
	if *fourSevenEight {
		var (
			in   = countdown.New(4)
			hold = countdown.New(7)
			out  = countdown.New(8)
		)
		round := 1
		tick := time.NewTicker(1 * time.Second)
		for {
			for range tick.C {
				if in.Next() {
					fmt.Print("IN ")
				} else if hold.Next() {
					fmt.Print("HOLD ")
				} else if out.Next() {
					fmt.Print("OUT ")
				} else {
					break
				}
			}
			fmt.Printf("(end of round %d)\n\n", round)
			in.Reset()
			hold.Reset()
			out.Reset()
			round++
		}
	}
	if *coherent {
		var (
			in  = countdown.New(5)
			out = countdown.New(5)
		)
		round := 1
		tick := time.NewTicker(1100 * time.Millisecond)
		for {
			for range tick.C {
				if in.Next() {
					fmt.Print("IN ")
				} else if out.Next() {
					fmt.Print("OUT ")
				} else {
					break
				}
			}
			fmt.Printf("(end of round %d)\n\n", round)
			in.Reset()
			out.Reset()
			round++
		}
	}
}
