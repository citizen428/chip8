package emulator

import "testing"

func TestSetPixels(t *testing.T) {
	s := screen{}
	s.setPixel(4, 2)
	want := true
	got := s.pixels[2][4]

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestPixelFuncs(t *testing.T) {
	s := screen{}
	s.pixels[2][4] = true
	want := true
	got := s.isPixelSet(4, 2)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestSetPixelValidatesCoordinates(t *testing.T) {
	defer func() { recover() }()

	s := screen{}
	s.setPixel(-1, 5)

	// Unreachable if `validateScreenCoordinates` panics as intended
	t.Errorf("did not panic")
}

func TestIsPixelSetValidatesCoordinates(t *testing.T) {
	defer func() { recover() }()

	s := screen{}
	s.isPixelSet(1, 100)

	// Unreachable if `validateScreenCoordinates` panics as intended
	t.Errorf("did not panic")
}
