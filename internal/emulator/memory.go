package emulator

// Reference: http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#2.1

const memorySize = 4096

type memory struct {
	storage [memorySize]uint8
}

// An invalid memory access in the emulator is not recoverable in Go code, so we panic.
func validateMemoryIndex(index int) {
	if index < 0 || index > memorySize {
		panic("Invalid memory access")
	}
}

func (m *memory) Set(index int, val uint8) {
	validateMemoryIndex(index)
	m.storage[index] = val
}

func (m memory) Get(index int) uint8 {
	validateMemoryIndex(index)
	return m.storage[index]
}
