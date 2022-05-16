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

func TestStackPushAddsValueToStack(t *testing.T) {
	var want uint16

	chip8 := NewChip8NoAudio()

	chip8.stackPush(42)
	want = 42
	got := chip8.stack[0]

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestStackPushIncrementsStackPointer(t *testing.T) {
	var want uint8

	chip8 := NewChip8NoAudio()
	chip8.stackPush(42)
	chip8.stackPush(42)
	want = 2
	got := chip8.registers.sp

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestStackPopReturnsValue(t *testing.T) {
	var want uint16

	chip8 := NewChip8NoAudio()
	chip8.stackPush(1)
	chip8.stackPush(2)

	want = 2
	got := chip8.stackPop()

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestStackPopDecrementsStackPointer(t *testing.T) {
	var want uint8

	chip8 := NewChip8NoAudio()
	chip8.registers.sp = 5
	chip8.stackPop()
	want = 4
	got := chip8.registers.sp

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestValidateStackDepth(t *testing.T) {
	defer func() { recover() }()

	chip8 := NewChip8NoAudio()
	chip8.registers.sp = 17
	chip8.stackPush(42)

	// Unreachable if `validateStackDepth` panics as intended
	t.Errorf("did not panic")
}

func TestCharacterSetInitialization(t *testing.T) {
	c := NewChip8NoAudio()
	want := []uint8{0xf0, 0x90, 0x90, 0x90, 0xf0}
	got := c.memory[0:5]

	if !reflect.DeepEqual(got, want) {
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
