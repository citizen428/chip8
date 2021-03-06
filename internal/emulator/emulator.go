package emulator

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

func Run(romPath string, scaleFactor int) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	emulatorWidth := int32(chip8width * scaleFactor)
	emulatorHeight := int32(chip8height * scaleFactor)

	window, err := sdl.CreateWindow("CHIP-8 — "+romPath, sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED, emulatorWidth, emulatorHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.TEXTUREACCESS_STATIC)
	if err != nil {
		panic(err)
	}

	chip8, closer := NewChip8()
	defer closer()
	rom, err := os.ReadFile(romPath)
	if err != nil {
		fmt.Printf("Error reading ROM: %+v", err)
		os.Exit(-1)
	}
	chip8.load(rom)

EventLoop:
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) {
			case *sdl.QuitEvent:
				break EventLoop
			case *sdl.KeyboardEvent:
				mappedKey, ok := mapKey(e.Keysym.Sym)
				if ok {
					if e.GetType() == sdl.KEYDOWN {
						chip8.keyboard.keyDown(mappedKey)
					} else {
						chip8.keyboard.keyUp(mappedKey)
					}
				}
			}
		}

		renderer.SetDrawColor(0, 0, 0, 0)
		renderer.Clear()
		renderer.SetDrawColor(255, 255, 255, 0)

		for y := 0; y < chip8height; y++ {
			for x := 0; x < chip8width; x++ {
				if chip8.screen.isPixelSet(x, y) {
					r := sdl.Rect{
						X: int32(x * scaleFactor),
						Y: int32(y * scaleFactor),
						W: int32(scaleFactor),
						H: int32(scaleFactor),
					}
					renderer.FillRect(&r)
				}
			}
		}

		renderer.Present()
		chip8.cycle()
	}
}
