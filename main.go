package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"time"

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

func open(fname string) (*os.File, error) {
	if fname == "-" {
		return os.Stdin, nil
	}
	return os.Open(fname)
}

// tuneFreq to full half wave length
func tuneFreq(tick time.Duration, freq float64) float64 {
	t := time.Duration(0.5 * float64(time.Second) / freq)
	n := time.Duration(math.Round(float64(tick) / float64(t)))
	return 0.5 / (tick / n).Seconds()
}

func main() {
	fname := flag.String("f", "-", "file to read")
	freq := flag.Float64("fq", 425.0, "Hz")
	wpm := flag.Int("wpm", 20, "words per minute")
	flag.Parse()

	sr := beep.SampleRate(48000)
	speaker.Init(sr, 4800)

	tick := wpmDuration(*wpm)
	*freq = tuneFreq(tick, *freq)

	wave, err := generators.SineTone(sr, *freq)
	if err != nil {
		panic(err)
	}

	g := gen{
		wave: wave,
		tick: sr.N(tick),
	}
	fd, err := open(*fname)
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	for v := range read(fd) {
		fmt.Print(string(v))
		snds := g.convert(v)
		speaker.PlayAndWait(snds)
	}
}
