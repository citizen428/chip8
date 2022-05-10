package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/citizen428/chip8/internal/emulator"
)

func main() {
	scaleFactor := flag.Int("scaleFactor", 10, "Display scale factor")
	flag.Parse()
	if flag.NArg() < 1 {
		flag.Usage()
		fmt.Println("  rom\n\tPath to ROM (mandatory)")
		os.Exit(1)
	}

	emulator.Run(flag.Arg(0), *scaleFactor)
}
