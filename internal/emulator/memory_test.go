package emulator

import "testing"

func TestSetGet(t *testing.T) {
	var want uint8

	m := memory{}
	index := 23
	m.set(index, 42)
	want = 42
	got := m.get(index)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
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

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
