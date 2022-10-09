package main

import (
	"log"
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

func RemoveAccent(s string) string {
	b := make([]byte, len(s))

	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	_, _, err := t.Transform(b, []byte(s), true)
	if err != nil {
		log.Fatal(err)
	}
	r := strings.TrimFunc(string(b), func(r rune) bool {
		return !unicode.IsGraphic(r)
	})
	return r
}

func RemoveBlanks(s string) string {
	return strings.ReplaceAll(s, " ", "")
}
