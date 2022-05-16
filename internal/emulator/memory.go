package emulator

const memorySize = 4096

type memory [memorySize]uint8

// An invalid memory access in the emulator is not recoverable in Go code, so we panic.
func validateMemoryIndex(index int) {
	if index < 0 || index > memorySize {
		panic("Invalid memory access")
	}
}

func (m *memory) set(index int, val uint8) {
	validateMemoryIndex(index)
	m[index] = val
}

func (m memory) get(index int) uint8 {
	validateMemoryIndex(index)
	return m[index]
}

func (m memory) ReadInstruction(index int) uint16 {
	validateMemoryIndex(index)
	byte1 := uint16(m.get(index))
	byte2 := uint16(m.get(index + 1))
	return 256*byte1 + byte2
}
