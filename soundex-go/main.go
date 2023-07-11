package main

import (
	"fmt"
	"github.com/go-workshop/soundex-go/soundex"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("input string missing")
		return
	}
	result, err := soundex.Soundex(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
