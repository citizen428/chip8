package emulator

import (
	"reflect"
	"testing"
)

func NewChip8NoAudio() chip8 {
	c := chip8{}
	copy(c.memory[0:80], defaultCharacterSet)
	return c
}

func TestCharacterSetInitialization(t *testing.T) {
	c := NewChip8NoAudio()
	want := []uint8{0xf0, 0x90, 0x90, 0x90, 0xf0}
	got := c.memory[0:5]

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
