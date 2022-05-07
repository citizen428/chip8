package emulator

import (
	"reflect"
	"testing"
)

func TestCharacterSetInitialization(t *testing.T) {
	c := NewChip8()
	want := []uint8{0xf0, 0x90, 0x90, 0x90, 0xf0}
	got := c.memory[0:5]

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
