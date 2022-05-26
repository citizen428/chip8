package emulator

import (
	"testing"

	"github.com/veandco/go-sdl2/sdl"
)

func TestIsKeyDown(t *testing.T) {
	k := keyboard{}
	key := 0xE

	t.Run("key is pressed", func(t *testing.T) {
		k.keyDown(key)

		want := true
		got := k.isKeyDown(key)
		assertEqual(t, want, got)
	})

	t.Run("key is not pressed", func(t *testing.T) {
		k.keyDown(key)
		k.keyUp(key)

		want := false
		got := k.isKeyDown(key)
		assertEqual(t, want, got)
	})
}

func TestMapKey(t *testing.T) {
	t.Run("looking up a mapped keys", func(t *testing.T) {
		var tests = map[sdl.Keycode]int{
			sdl.K_1: 1,
			sdl.K_w: 5,
			sdl.K_d: 9,
			sdl.K_v: 15,
		}

		for key, want := range tests {
			got, _ := mapKey(key)
			assertEqual(t, got, want)
		}
	})

	t.Run("looking up a unmapped keys", func(t *testing.T) {
		want := -1
		got, _ := mapKey(sdl.K_RETURN)
		assertEqual(t, got, want)
	})
}
