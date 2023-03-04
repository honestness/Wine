package main

import (
	"os"

	"github.com/shpota/goxygen/cli"
	"github.com/shpota/goxygen/codegen"
)

func main() {
	args := os.Args[1:]
	cli.Start(os.Stdout, args, codegen.Generate)
	// cli.Start(os.Stdout, args, codegen.Generate)
}
