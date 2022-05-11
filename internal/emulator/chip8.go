package emulator

import (
	"time"
)

const (
	dataRegistersCount = 16
	memorySize         = 4096
	stackDepth         = 16
	delayMs            = 17 // 60 Hz ~ every 17ms
	programLoadAddress = 0x200
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
	screen    screen
	speaker   speaker
}

// Reference: http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#2.4
var defaultCharacterSet = []uint8{
	0xf0, 0x90, 0x90, 0x90, 0xf0, // 0
	0x20, 0x60, 0x20, 0x20, 0x70, // 1
	0xf0, 0x10, 0xf0, 0x80, 0xf0, // 2
	0xf0, 0x10, 0xf0, 0x10, 0xf0, // 3
	0x90, 0x90, 0xf0, 0x10, 0x10, // 4
	0xf0, 0x80, 0xf0, 0x10, 0xf0, // 5
	0xf0, 0x80, 0xf0, 0x90, 0xf0, // 6
	0xf0, 0x10, 0x20, 0x40, 0x40, // 7
	0xf0, 0x90, 0xf0, 0x90, 0xf0, // 8
	0xf0, 0x90, 0xf0, 0x10, 0xf0, // 9
	0xf0, 0x90, 0xf0, 0x90, 0x90, // A
	0xe0, 0x90, 0xe0, 0x90, 0xe0, // B
	0xf0, 0x80, 0x80, 0x80, 0xf0, // C
	0xe0, 0x90, 0x90, 0x90, 0xe0, // D
	0xf0, 0x80, 0xf0, 0x80, 0xf0, // E
	0xf0, 0x80, 0xf0, 0x80, 0x80, // F
}

func NewChip8() chip8 {
	c := chip8{}
	copy(c.memory[0:80], defaultCharacterSet)
	c.speaker = NewSpeaker()
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

func (m memory) ReadInstruction(index int) uint16 {
	validateMemoryIndex(index)
	byte1 := uint16(m.memGet(index))
	byte2 := uint16(m.memGet(index + 1))
	return 256*byte1 + byte2
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

func (r *registers) incrementPC() {
	r.pc += 2
}

// Reference: http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#2.5
func (c *chip8) handleDelayTimer() {
	if c.registers.dt > 0 {
		time.Sleep(delayMs * time.Millisecond)
		c.registers.dt--
	}
}

// Reference: http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#2.5
func (c *chip8) handleSoundTimer() {
	play := c.registers.st > 0
	c.speaker.beep(play)
	if play {
		time.Sleep(delayMs * time.Millisecond)
		c.registers.st--
	}
}

func (c *chip8) load(rom []byte) {
	size := len(rom)
	if size+programLoadAddress >= memorySize {
		panic("ROM too big")
	}

	copy(c.memory[programLoadAddress:], rom)
	c.registers.pc = programLoadAddress
}
