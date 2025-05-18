package main

import (
	"bufio"
	"io"
	"unicode/utf8"
)

// read runes from io.Reader
func read(r io.Reader) <-chan rune {
	ch := make(chan rune)
	go func() {
		defer close(ch)
		s := bufio.NewScanner(r)
		s.Split(bufio.ScanRunes)
		for s.Scan() {
			r, _ := utf8.DecodeRuneInString(s.Text())
			ch <- r
		}
	}()
	return ch
}
