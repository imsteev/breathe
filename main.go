package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/imsteev/breathe/countdown"
)

func main() {
	box := flag.Bool("box", false, "box breathe")
	fourSevenEight := flag.Bool("478", false, "four-seven-eight breathing")
	coherent := flag.Bool("coherent", false, "coherent breathing")
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
					fmt.Print("IN (4) ")
				} else if firstHold.Next() {
					fmt.Print("HOLD (4) ")
				} else if secIn.Next() {
					fmt.Print("OUT (4) ")
				} else if secHold.Next() {
					fmt.Print("HOLD (4) ")
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
		round := 0
		tick := time.NewTicker(1 * time.Second)
		for range tick.C {
			if in.Next() {
				fmt.Print("IN ")
			} else if hold.Next() {
				fmt.Print("HOLD ")
			} else if out.Next() {
				fmt.Print("OUT ")
			} else {
				fmt.Printf("(end of round %d)\n\n", round)
				in.Reset()
				hold.Reset()
				out.Reset()
				round++
			}
		}
	}
	if *coherent {
		var (
			in  = countdown.New(5)
			out = countdown.New(5)
		)
		round := 0
		tick := time.NewTicker(1100 * time.Millisecond)
		for range tick.C {
			if in.Next() {
				fmt.Print("IN ")
			} else if out.Next() {
				fmt.Print("OUT ")
			} else {
				fmt.Printf("(end of round %d)\n\n", round)
				in.Reset()
				out.Reset()
				round++
			}
		}
	}
	if !*box && !*fourSevenEight && !*coherent {
		fmt.Println(`
Box breathing: breathe in for 4 seconds, hold for 4 seconds,
exhale for 4 seconds, hold for 4 seconds.

4-7-8: breathe in for 4 seconds, hold for 7 seconds, exhale for 8 seconds.
Good for activating parasympathetic nervous system (aka, relaxation)

Coherent breathing: inhale and exhale, circularly, for 5.5 seconds each way.`)}
}
