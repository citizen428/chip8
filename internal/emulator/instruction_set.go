package emulator

import "math/rand"

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
	n := opcode & 0x000F

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

	// 9xy0 - SNE Vx, Vy - Skip next instruction if Vx != Vy.
	case 0x9000:
		if c.registers.v[x] != c.registers.v[y] {
			c.registers.incrementPC()
		}

	// Annn - LD I, addr - Set I = nnn.
	case 0xA000:
		c.registers.i = nnn

	// Bnnn - JP V0, addr - Jump to location nnn + V0.
	case 0xB000:
		c.registers.pc = nnn + uint16(c.registers.v[0])

	// Cxkk - RND Vx, byte - Set Vx = random byte AND kk.
	case 0xC000:
		random := uint8(rand.Intn(255))
		c.registers.v[x] = random & kk

	// Dxyn - DRW Vx, Vy, nibble
	// Display n-byte sprite starting at memory location I at (Vx, Vy), set VF = collision.
	case 0xD000:
		c.registers.v[0x0F] = 0
		sprite := c.memory[c.registers.i : c.registers.i+n]
		if c.screen.drawSprite(int(c.registers.v[x]), int(c.registers.v[y]), sprite) {
			c.registers.v[0x0F] = 1
		}

	case 0xE000:
		switch opcode & 0xff {
		// Ex9E - SKP Vx - Skip next instruction if key with the value of Vx is pressed.
		case 0x9E:
			if c.keyboard.isKeyDown(int(c.registers.v[x])) {
				c.registers.incrementPC()
			}

		// ExA1 - SKNP Vx - Skip next instruction if key with the value of Vx is not pressed.
		case 0xA1:
			if !c.keyboard.isKeyDown(int(c.registers.v[x])) {
				c.registers.incrementPC()
			}
		}

	case 0xF000:
		c.execFxxx(opcode)
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

func (c *chip8) execFxxx(opcode uint16) {
	x := (opcode >> 8) & 0x000F

	switch opcode & 0xFF {

	// Fx07 - LD Vx, DT - Set Vx = delay timer value.
	case 0x07:
		c.registers.v[x] = c.registers.dt

	// Fx0A - LD Vx, K - Wait for a key press, store the value of the key in Vx.
	case 0x0A:
		key, ok := c.waitForKeypress()
		if ok {
			c.registers.v[x] = uint8(key)
		}

	// Fx15 - LD DT, Vx - Set delay timer = Vx.
	case 0x15:
		c.registers.dt = c.registers.v[x]

	// Fx18 - LD ST, Vx - Set sound timer = Vx.
	case 0x18:
		c.registers.st = c.registers.v[x]

	// Fx1E - ADD I, Vx - Set I = I + Vx.
	case 0x1E:
		c.registers.i += uint16(c.registers.v[x])

	// Fx29 - LD F, Vx - Set I = location of sprite for digit Vx.
	case 0x29:
		c.registers.i = uint16(c.registers.v[x]) * spriteHeight

	// Fx33 - LD B, Vx
	// Store BCD representation of Vx in memory locations I, I+1, and I+2.
	case 0x33:
		hundreds := c.registers.v[x] / 100
		tens := c.registers.v[x] / 10 % 10
		units := c.registers.v[x] % 10

		c.memory.set(int(c.registers.i), hundreds)
		c.memory.set(int(c.registers.i+1), tens)
		c.memory.set(int(c.registers.i+2), units)

	// Fx55 - LD [I], Vx - Store registers V0 through Vx in memory starting at location I.
	case 0x55:
		for i := uint16(0); i < x; i++ {
			c.memory.set(int(c.registers.i+i), c.registers.v[i])
		}

	// x65 - LD Vx, [I] - Read registers V0 through Vx from memory starting at location I.
	case 0x65:
		for i := uint16(0); i < x; i++ {
			c.registers.v[i] = c.memory.get(int(c.registers.i + i))
		}
	}

}
