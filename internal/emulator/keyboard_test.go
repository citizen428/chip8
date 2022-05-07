package emulator

import (
	"testing"

	"github.com/veandco/go-sdl2/sdl"
)

func TestIsKeyDown(t *testing.T) {
	k := keyboard{}
	key := 0xE
	k.keyDown(key)

	want := true
	got := k.isKeyDown(key)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestKeyUp(t *testing.T) {
	k := keyboard{}
	key := 0xA
	k.keyDown(key)
	k.keyUp(key)

	want := false
	got := k.isKeyDown(key)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestMap(t *testing.T) {
	var tests = map[sdl.Keycode]int{
		sdl.K_1: 1,
		sdl.K_w: 5,
		sdl.K_d: 9,
		sdl.K_v: 15,
	}

	for key, want := range tests {
		got, _ := mapKey(key)
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}
}
