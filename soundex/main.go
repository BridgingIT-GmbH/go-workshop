package main

import (
	"fmt"
	"github.com/go-workshop/soundex/soundex"
	"os"
)

func main() {
	if os.Args[1] == "" {
		fmt.Println("Input String missing")
		return
	}
	fmt.Println(soundex.Soundex(os.Args[1]))
}
