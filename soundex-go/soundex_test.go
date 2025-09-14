package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGolang(t *testing.T) {
	soundex, _ := Soundex("Golang")
	assert.Equal(t, "G452", soundex)
}

func TestHallo(t *testing.T) {
	soundex, _ := Soundex("Hallo")
	assert.Equal(t, "H400", soundex)
}

func TestEntwicklertag(t *testing.T) {
	soundex, _ := Soundex("Entwicklertag")
	assert.Equal(t, "E532", soundex)
}

func TestWow(t *testing.T) {
	soundex, _ := Soundex("Wow")
	assert.Equal(t, "W000", soundex)
}

func TestError(t *testing.T) {
	_, err := Soundex("")
	assert.NotNil(t, err)
	assert.Equal(t, "empty input string", err.Error())
}
