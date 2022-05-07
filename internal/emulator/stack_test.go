package emulator

import "testing"

func TestStackPushAddsValueToStack(t *testing.T) {
	var want uint16

	chip8 := NewChip8()
	chip8.stackPush(42)
	want = 42
	got := chip8.stack[0]

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestStackPushIncrementsStackPointer(t *testing.T) {
	var want uint8

	chip8 := NewChip8()
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

	chip8 := NewChip8()
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

	chip8 := NewChip8()
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

	chip8 := NewChip8()
	chip8.registers.sp = 17
	chip8.stackPush(42)

	// Unreachable if `validateStackDepth` panics as intended
	t.Errorf("did not panic")
}
