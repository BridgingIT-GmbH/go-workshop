package main

import (
	"testing"
)

func TestGolang(t *testing.T) {
	soundex := Soundex("Golang")
	if soundex != "G452" {
		t.Errorf("Soundex of Golang should be G452 but was %s", soundex)
	}
}

func TestHallo(t *testing.T) {
	soundex := Soundex("Hallo")
	if soundex != "H400" {
		t.Errorf("Soundex of Hallo should be H400 but was %s", soundex)
	}
}

func TestEntwicklertag(t *testing.T) {
	soundex := Soundex("Entwicklertag")
	if soundex != "E532" {
		t.Errorf("Soundex of Entwicklertag should be E532 but was %s", soundex)
	}
}

func TestWow(t *testing.T) {
	soundex := Soundex("Wow")
	if soundex != "W000" {
		t.Errorf("Soundex of Wow should be W000 but was %s", soundex)
	}
}
