package main

import (
	"unicode"

	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/generators"
)

type gen struct {
	sine beep.Streamer
	tick int
}

func (g gen) convert(r rune) beep.Streamer {
	var st []beep.Streamer
	if s, ok := code[unicode.ToUpper(r)]; ok {
		for _, x := range s {
			switch x {
			case '-':
				st = append(st, beep.Take(g.tick*3, g.sine))
			case '.':
				st = append(st, beep.Take(g.tick, g.sine))
			}
			st = append(st, generators.Silence(g.tick))
		}
	} else if unicode.IsSpace(r) {
		st = append(st, generators.Silence(g.tick*2))
	}
	st = append(st, generators.Silence(g.tick*2))
	return beep.Seq(st...)
}
