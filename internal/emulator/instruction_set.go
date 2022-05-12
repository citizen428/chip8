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

	case 0x8000:
		c.exec8xxx(opcode)
	}
}

func (c *chip8) exec8xxx(opcode uint16) {
	x := (opcode >> 8) & 0x000F
	y := (opcode >> 4) & 0x000F
	finalFourBits := opcode & 0x000F

	switch finalFourBits {

	// 8xy0 - LD Vx, Vy - Set Vx = Vy.
	case 0x00:
		c.registers.v[x] = c.registers.v[y]

	// 8xy1 - OR Vx, Vy - Set Vx = Vx OR Vy.
	case 0x01:
		c.registers.v[x] |= c.registers.v[y]

	// 8xy2 - AND Vx, Vy - Set Vx = Vx AND Vy.
	case 0x02:
		c.registers.v[x] &= c.registers.v[y]

	// 8xy3 - XOR Vx, Vy - Set Vx = Vx XOR Vy.
	case 0x03:
		c.registers.v[x] ^= c.registers.v[y]

	// 8xy4 - ADD Vx, Vy - Set Vx = Vx + Vy, set VF = carry.
	case 0x04:
		result := int(c.registers.v[x]) + int(c.registers.v[y])
		c.registers.v[0xF] = 0
		if result > 0xFF {
			c.registers.v[0xF] = 1
		}
		c.registers.v[x] = uint8(result)

	// 8xy5 - SUB Vx, Vy - Set Vx = Vx - Vy, set VF = NOT borrow.
	case 0x05:
		c.registers.v[0xF] = 0
		if c.registers.v[x] > c.registers.v[y] {
			c.registers.v[0xF] = 1
		}
		c.registers.v[x] -= c.registers.v[y]

	// 8xy6 - SHR Vx {, Vy} - Set Vx = Vx SHR 1.
	case 0x06:
		c.registers.v[0xF] = c.registers.v[x] & 1
		c.registers.v[x] /= 2

	// 8xy7 - SUBN Vx, Vy - Set Vx = Vy - Vx, set VF = NOT borrow.
	case 0x07:
		c.registers.v[0xF] = 0
		if c.registers.v[y] > c.registers.v[x] {
			c.registers.v[0xF] = 1
		}
		c.registers.v[x] = c.registers.v[y] - c.registers.v[x]

	// 8xyE - SHL Vx {, Vy} - Set Vx = Vx SHL 1.
	case 0x0E:
		c.registers.v[0xF] = c.registers.v[x] & 0b1000_0000
		c.registers.v[x] *= 2
	}
}
