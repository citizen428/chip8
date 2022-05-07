package main

import (
	"flag"

	"github.com/citizen428/chip8/internal/emulator"
)

func main() {
	scaleFactor := flag.Int("scaleFactor", 10, "Display scale factor")
	flag.Parse()

	emulator.Run(*scaleFactor)
}
