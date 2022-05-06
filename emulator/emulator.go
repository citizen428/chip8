package emulator

import "github.com/veandco/go-sdl2/sdl"

const windowTitle = "CHIP-8"
const emulatorWidth = 640
const emulatorHeight = 320

func Run() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow(windowTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		emulatorWidth, emulatorHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.TEXTUREACCESS_STATIC)
	if err != nil {
		panic(err)
	}

EventLoop:
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				break EventLoop
			}
		}
		renderer.SetDrawColor(0, 0, 0, 0)
		renderer.Clear()
		renderer.SetDrawColor(255, 255, 255, 0)
		r := sdl.Rect{X: 0, Y: 0, W: 40, H: 40}
		renderer.FillRect(&r)
		renderer.Present()

	}
}
