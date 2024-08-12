package main

import (
	"flag"
	"fmt"
	"time"
)

type CountDown struct {
	to      int
	current int
}

func (c *CountDown) Next() bool {
	if c.current < c.to {
		c.current++
		return true
	}
	return false
}

func (c *CountDown) Reset() {
	c.current = 0
}

func main() {
	box := flag.Bool("box", false, "box breathe")
	fourSevenEight := flag.Bool("478", false, "four-seven-eight breathing")
	coherent := flag.Bool("coherent", false, "coherent breathing")
	flag.Parse()

	if *box {
		var (
			firstIn   = &CountDown{to: 4}
			firstHold = &CountDown{to: 4}
			secIn     = &CountDown{to: 4}
			secHold   = &CountDown{to: 4}
		)

		round := 1
		tick := time.NewTicker(1000 * time.Millisecond)
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
				fmt.Printf("(end of round %d)\n\n", round)
				firstIn.Reset()
				firstHold.Reset()
				secIn.Reset()
				secHold.Reset()
				round++
			}
		}
	}
	if *fourSevenEight {
		var (
			in   = &CountDown{to: 4}
			hold = &CountDown{to: 7}
			out  = &CountDown{to: 8}
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
			in  = &CountDown{to: 5}
			out = &CountDown{to: 5}
		)
		round := 0
		tick := time.NewTicker(1 * time.Second)
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
}
