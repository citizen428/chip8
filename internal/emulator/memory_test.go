package emulator

import "testing"

func TestSet(t *testing.T) {
	m := memory{}
	index := 23

	m.set(index, 42)

	var want uint8 = 42
	got := m[index]
	assertEqual(t, want, got)
}

func TestGet(t *testing.T) {
	m := memory{}
	index := 23

	var want uint8 = m.get(index)
	got := m[index]
	assertEqual(t, want, got)
}

func TestMemoryLowerBound(t *testing.T) {
	defer func() { recover() }()

	NewChip8NoAudio().memory.get(-1)

	// Unreachable if `Get` panics as intended
	t.Errorf("did not panic")
}

func TestMemoryUpperBound(t *testing.T) {
	defer func() { recover() }()

	chip8 := NewChip8NoAudio()
	chip8.memory.get(memorySize + 1)

	// Unreachable if `Get` panics as intended
	t.Errorf("did not panic")
}

func TestReadInstruction(t *testing.T) {
	c := NewChip8NoAudio()
	c.memory[0x200] = 192
	c.memory[0x201] = 168

	var want uint16 = 0xc0a8
	got := c.memory.readInstruction(0x200)
	assertEqual(t, got, want)
}
