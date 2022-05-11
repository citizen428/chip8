package emulator

// Reference: http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#3.0

func (c *chip8) exec(opcode uint16) {
	switch opcode {
	// CLS - Clear the display.
	case 0x00E0:
		c.screen.clear()

	// RET - Return from a subroutine.
	case 0x00EE:
		c.registers.pc = c.stackPop()

	default:
		c.execExtended(opcode)
	}
}

// nnn or addr - A 12-bit value, the lowest 12 bits of the instruction
// n or nibble - A 4-bit value, the lowest 4 bits of the instruction
// x - A 4-bit value, the lower 4 bits of the high byte of the instruction
// y - A 4-bit value, the upper 4 bits of the low byte of the instruction
// kk or byte - An 8-bit value, the lowest 8 bits of the instruction
func (c *chip8) execExtended(opcode uint16) {
	nnn := opcode & 0x0FFF
	x := (opcode >> 8) & 0x000F
	y := (opcode >> 4) & 0x000F
	kk := uint8(opcode & 0x00FF)

	switch opcode & 0xF000 {

	// 0nnn - SYS addr - Jump to a machine code routine at nnn.
	case 0x000:
		// Ignored by modern interpreters.

	// 1nnn - JP addr - Jump to location nnn
	case 0x1000:
		c.registers.pc = nnn

	// 2nnn - CALL addr - Call subroutine at nnn.
	case 0x2000:
		c.stackPush(c.registers.pc)
		c.registers.pc = nnn

	// 3xkk - SE Vx, byte - Skip next instruction if Vx = kk.
	case 0x3000:
		if c.registers.v[x] == kk {
			c.registers.incrementPC()
		}

	// 4xkk - SE Vx, byte - Skip next instruction if Vx = kk.
	case 0x4000:
		if c.registers.v[x] == kk {
			c.registers.incrementPC()
		}

	// 5xy0 - SE Vx, Vy - Skip next instruction if Vx = Vy.
	case 0x5000:
		if c.registers.v[x] == c.registers.v[y] {
			c.registers.incrementPC()
		}

	// 6xkk - LD Vx, byte - Set Vx = kk.
	case 0x6000:
		c.registers.v[x] = kk

	// 7xkk - ADD Vx, byte - Set Vx = Vx + kk.
	case 0x7000:
		c.registers.v[x] += kk
	}
}
