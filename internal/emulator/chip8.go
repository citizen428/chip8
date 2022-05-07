package emulator

const (
	dataRegistersCount = 16
	memorySize         = 4096
	stackDepth         = 16
)

type memory [memorySize]uint8
type stack [stackDepth]uint16

// Reference: http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#2.2
type registers struct {
	v  [dataRegistersCount]uint8 // data registers
	i  uint16                    // generally used to store addresses
	dt uint8                     // delay timer
	st uint8                     // sound timer
	pc uint16                    // program counter
	sp uint8                     // stack pointer
}

type chip8 struct {
	memory    memory
	registers registers
	stack     stack
	keyboard  keyboard
}

func NewChip8() chip8 {
	c := chip8{}
	copy(c.memory[0:80], defaultCharacterSet)
	return c
}

// An invalid memory access in the emulator is not recoverable in Go code, so we panic.
func validateMemoryIndex(index int) {
	if index < 0 || index > memorySize {
		panic("Invalid memory access")
	}
}

func (m *memory) memSet(index int, val uint8) {
	validateMemoryIndex(index)
	m[index] = val
}

func (m memory) memGet(index int) uint8 {
	validateMemoryIndex(index)
	return m[index]
}

func (c chip8) validateStackDepth() {
	if c.registers.sp > stackDepth {
		panic("Stack overflow")
	}
}

func (c *chip8) stackPush(val uint16) {
	c.validateStackDepth()
	c.stack[c.registers.sp] = val
	c.registers.sp++
}

func (c *chip8) stackPop() uint16 {
	c.registers.sp--
	c.validateStackDepth()
	val := c.stack[c.registers.sp]
	return val
}
