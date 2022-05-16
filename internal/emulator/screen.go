package emulator

// Reference: http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#2.4=

const (
	chip8width   = 64
	chip8height  = 32
	spriteHeight = 5
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

func (s *screen) drawSprite(x int, y int, sprite []uint8) bool {
	pixelCollission := false

	for ly, byte := range sprite {
		for lx := 0; lx < 8; lx++ {
			if byte&(0b10000000>>lx) > 0 {
				// Sprites wrap around the edges of the screen
				dx := (x + lx) % chip8width
				dy := (y + ly) % chip8height

				pixelIsSet := s.isPixelSet(dx, dy)
				if pixelIsSet {
					pixelCollission = true
				}
				// XOR the previous pixel with true
				s.pixels[dy][dx] = !pixelIsSet
			}
		}
	}

	return pixelCollission
}

func (s *screen) clear() {
	s.pixels = [chip8height][chip8width]bool{}
}
