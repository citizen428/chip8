package emulator

const stackDepth = 16

type stack [stackDepth]uint16

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
