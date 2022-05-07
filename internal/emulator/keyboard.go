package emulator

import (
	"github.com/veandco/go-sdl2/sdl"
)

// Reference: http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#2.3

const keyCount = 16

type keyboard struct {
	keys [keyCount]bool
}

// Mapping from user's keyboard to emulated keyboard
var keyboardMap = map[sdl.Keycode]int{
	sdl.K_1: 1,
	sdl.K_2: 2,
	sdl.K_3: 3,
	sdl.K_4: 12,
	sdl.K_q: 4,
	sdl.K_w: 5,
	sdl.K_e: 6,
	sdl.K_r: 13,
	sdl.K_a: 7,
	sdl.K_s: 8,
	sdl.K_d: 9,
	sdl.K_f: 14,
	sdl.K_z: 10,
	sdl.K_x: 0,
	sdl.K_c: 11,
	sdl.K_v: 15,
}

func mapKey(key sdl.Keycode) (int, bool) {
	if v, ok := keyboardMap[key]; ok {
		return v, true
	}
	return -1, false
}

func (k *keyboard) keyDown(key int) {
	k.keys[key] = true
}

func (k *keyboard) keyUp(key int) {
	k.keys[key] = false
}

func (k keyboard) isKeyDown(key int) bool {
	return k.keys[key]
}
