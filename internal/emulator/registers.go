package emulator

const dataRegistersCount = 16

// Reference: http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#2.2
type registers struct {
	v  [dataRegistersCount]uint8 // data registers
	i  uint16                    // generally used to store addresses
	dt uint8                     // delay timer
	st uint8                     // sound timer
	pc uint16                    // program counter
	sp uint8                     // stack pointer
}

func (r *registers) incrementPC() {
	r.pc += 2
}
