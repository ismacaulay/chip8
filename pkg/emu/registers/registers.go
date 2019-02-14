package registers

// ReaderWriter interface provide functions to interact with the registers
type ReaderWriter interface {
	SetProgramCounter(address uint16)
	GetProgramCounter() uint16
	IncrementProgramCounter(increment uint16)
	PushProgramCounter()
	PopProgramCounter() uint16

	GetRegisterValue(register uint8) uint8
	SetRegisterValue(register, value uint8)
	SetRegisterI(value uint16)
	GetRegisterI() uint16

	SetDelayTimer(value uint8)
	GetDelayTimer() uint8
	SetSoundTimer(value uint8)
	GetSoundTimer() uint8
}

// Registers is the implementation of the Registers interface
type Registers struct {
	v          []uint8
	i          uint16
	pc         uint16
	sp         uint8
	stack      []uint16
	delayTimer uint8
	soundTimer uint8
}

// NewRegisters creates a new instance of Registers
func NewRegisters() *Registers {
	return &Registers{
		v:          make([]uint8, 16),
		i:          0,
		pc:         0x200,
		sp:         0,
		stack:      make([]uint16, 16),
		delayTimer: 0,
		soundTimer: 0,
	}
}

// SetProgramCounter updates the program counter to the new address
func (r *Registers) SetProgramCounter(address uint16) {
	r.pc = address
}

// GetProgramCounter returns the current value of the program counter
func (r *Registers) GetProgramCounter() uint16 {
	return r.pc
}

// IncrementProgramCounter increments the program counter by the specified amount
func (r *Registers) IncrementProgramCounter(increment uint16) {
	r.pc = r.pc + (2 * increment)
}

// PushProgramCounter saves the current program counter on the stack
func (r *Registers) PushProgramCounter() {
	r.stack[r.sp] = r.pc
	r.sp++
}

// PopProgramCounter removes the current program counter off the top of the stack
// and returns it
func (r *Registers) PopProgramCounter() uint16 {
	r.sp--
	return r.stack[r.sp]
}

// GetRegisterValue returns the value of the register
func (r *Registers) GetRegisterValue(register uint8) uint8 {
	return r.v[register]
}

// SetRegisterValue sets the value of the register
func (r *Registers) SetRegisterValue(register, value uint8) {
	r.v[register] = value
}

// SetRegisterI sets the value of register I
func (r *Registers) SetRegisterI(value uint16) {
	r.i = value
}

// GetRegisterI return the value of register I
func (r *Registers) GetRegisterI() uint16 {
	return r.i
}

// SetDelayTimer sets the delay timer value
func (r *Registers) SetDelayTimer(value uint8) {
	r.delayTimer = value
}

// GetDelayTimer returns the delay timer value
func (r *Registers) GetDelayTimer() uint8 {
	return r.delayTimer
}

// SetSoundTimer sets the sound timer value
func (r *Registers) SetSoundTimer(value uint8) {
	r.soundTimer = value
}

// GetSoundTimer returns the sound timer value
func (r *Registers) GetSoundTimer() uint8 {
	return r.soundTimer
}
