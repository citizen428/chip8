package emulator

const (
	chip8width  = 64
	chip8height = 32
)

type screen struct {
	pixels [chip8height][chip8width]bool
}

func validateScreenCoordinates(x, y int) {
	if x < 0 || x > chip8width-1 || y < 0 || y > chip8height-1 {
		panic("Invalid coordinate")
	}
}

func (s *screen) setPixel(x, y int) {
	validateScreenCoordinates(x, y)
	s.pixels[y][x] = true
}

func (s *screen) isPixelSet(x, y int) bool {
	validateScreenCoordinates(x, y)
	return s.pixels[y][x]
}
