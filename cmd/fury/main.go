package main

import (
	"fmt"
	"os"

	"github.com/patdhlk/fury/repl"
)

func main() {
	fmt.Printf("FURY shell - the Fury programming language!\n")
	repl.Start(os.Stdin, os.Stdout)
}
