package soundex

import (
	"testing"
)

func TestGolang(t *testing.T) {
	soundex, _ := Soundex("Golang")
	if soundex != "G452" {
		t.FailNow()
	}
}

func TestHallo(t *testing.T) {
	soundex, _ := Soundex("Hallo")
	if soundex != "H400" {
		t.FailNow()
	}
}

func TestEntwicklertag(t *testing.T) {
	soundex, _ := Soundex("Entwicklertag")
	if soundex != "E532" {
		t.FailNow()
	}
}

func TestWow(t *testing.T) {
	soundex, _ := Soundex("Wow")
	if soundex != "W000" {
		t.FailNow()
	}
}

func TestErrorEntwicklertag(t *testing.T) {
	_, err := Soundex("")
	if err == nil {
		t.FailNow()
	}
}
