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
	t := float64(time.Second/2) / freq
	n := time.Duration(math.Round(float64(tick) / t))
	return 0.5 / (tick / n).Seconds()
}

func main() {
	freq := flag.Float64("fq", 425.0, "Hz")
	fname := flag.String("f", "-", "file to read")
	srate := flag.Int("sr", 48000, "sample rate")
	bufsz := flag.Int("bs", 4800, "buffer size")
	wpm := flag.Int("wpm", 20, "words per minute")
	flag.Parse()

	sr := beep.SampleRate(*srate)
	speaker.Init(sr, *bufsz)

	tick := wpmDuration(*wpm)

	wave, err := generators.SineTone(sr, tuneFreq(tick, *freq))
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
