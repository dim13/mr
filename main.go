package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/generators"
	"github.com/gopxl/beep/v2/speaker"
)

/*
Dot (dit): One time unit.
Dash (dah): Three time units (three dots).
Space between dots and dashes within a letter: One time unit.
Space between letters: Three time units.
Space between words: Seven time units.
*/

func main() {
	fq := flag.Float64("fq", 425.0, "Hz")
	wpm := flag.Int("wpm", 20, "words per minute")
	flag.Parse()

	sr := beep.SampleRate(48000)
	speaker.Init(sr, 4800)

	sine, err := generators.SineTone(sr, *fq)
	if err != nil {
		panic(err)
	}

	tick := wpmDuration(*wpm)
	g := gen{
		sine: sine,
		tick: sr.N(tick),
	}

	for v := range read(os.Stdin) {
		fmt.Print(string(v))
		snds := g.convert(v)
		speaker.PlayAndWait(snds)
	}
}
