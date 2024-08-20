package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

func main() {
	var (
		box            = flag.Bool("box", false, "Box breathing: in 4s, hold 4s, out 4s, hold 4s")
		fourSevenEight = flag.Bool("478", false, "4-7-8 breathing: in 4s, hold 7s, out 8s. Triggers parasympathetic nervous system (aka, good for relaxation)")
		coherent       = flag.Bool("coh", false, "Coherent breathing: in 5.5s, out 5.5s. Try to make it as circular as possible")

		ui = flag.Bool("ui", false, "optional HTML interface")
	)

	flag.Parse()

	var sequence []string
	rate := 1 * time.Second // default

	if *box {
		sequence = []string{"IN", ".", ".", ".", "HOLD", ".", ".", ".", "OUT", ".", ".", ".", "HOLD", ".", ".", "."}
	}

	if *fourSevenEight {
		sequence = []string{"IN", ".", ".", ".", "HOLD", ".", ".", ".", ".", ".", ".", "OUT", ".", ".", ".", ".", ".", ".", "."}
	}

	if *coherent {
		sequence = []string{"IN", ".", ".", ".", ".", "OUT", ".", ".", ".", "."}
		rate = 1100 * time.Millisecond
	}

	if *box || *fourSevenEight || *coherent {
		round := 1
		tick := time.NewTicker(rate)
		for {
			for _, step := range sequence {
				<-tick.C
				fmt.Printf("%s ", step)
			}
			fmt.Printf("(%d)\n\n", round)
			round++
		}
	}

	if *ui {
		// WIP
		http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
			t, err := template.ParseFiles("breathe.html")
			if err != nil {
				log.Fatalf("err: %v", err)
			}
			t.Execute(w, nil)
		})
		http.HandleFunc("/box", func(w http.ResponseWriter, req *http.Request) {
			w.Header().Add("Content-Type", "text/html")
			t, err := template.New("breathing-technique").Parse(`<div id="box"></div>`)
			if err != nil {
				log.Fatal(err)
			}
			t.Execute(w, nil)
		})
		http.HandleFunc("/478", func(w http.ResponseWriter, req *http.Request) {
			w.Header().Add("Content-Type", "text/html")
			t, err := template.New("breathing-technique").Parse(`<div id="fourSevenEight"></div>`)
			if err != nil {
				log.Fatal(err)
			}
			t.Execute(w, nil)
		})
		http.HandleFunc("/coh", func(w http.ResponseWriter, req *http.Request) {
			w.Header().Add("Content-Type", "text/html")
			t, err := template.New("breathing-technique").Parse(`<div id="coherent"></div>`)
			if err != nil {
				log.Fatal(err)
			}
			t.Execute(w, nil)
		})
		fmt.Println("view breathing exercises at: http://localhost:3000")
		if err := http.ListenAndServe(":3000", nil); err != nil {
			log.Fatalf("error serving: %s", err)
		}
	}
}
