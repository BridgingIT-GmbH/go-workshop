package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(Soundex(os.Args[1]))
}
