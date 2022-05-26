package emulator

import "testing"

func TestStackPush(t *testing.T) {
	chip8 := NewChip8NoAudio()

	t.Run("pushes a value onto the stack", func(t *testing.T) {
		chip8.stackPush(42)

		want := uint16(42)
		got := chip8.stack[0]
		assertEqual(t, got, want)
	})

	t.Run("increments the stack pointer", func(t *testing.T) {
		chip8.registers.sp = 3
		chip8.stackPush(42)

		want := uint8(4)
		got := chip8.registers.sp
		assertEqual(t, got, want)
	})
}

func TestStackPop(t *testing.T) {
	chip8 := NewChip8NoAudio()

	t.Run("pops a value off the stack", func(t *testing.T) {
		chip8.stackPush(1)
		chip8.stackPush(2)

		want := uint16(2)
		got := chip8.stackPop()
		assertEqual(t, got, want)
	})

	t.Run("decrements the stack pointer", func(t *testing.T) {
		chip8.registers.sp = 5
		chip8.stackPop()

		want := uint8(4)
		got := chip8.registers.sp
		assertEqual(t, got, want)
	})
}

func TestValidateStackDepth(t *testing.T) {
	defer func() { recover() }()

	chip8 := NewChip8NoAudio()
	chip8.registers.sp = 17
	chip8.stackPush(42)

	// Unreachable if `validateStackDepth` panics as intended
	t.Errorf("did not panic")
}
