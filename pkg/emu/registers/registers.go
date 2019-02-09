package registers

type Registers interface {
	SetProgramCounter(address uint16)
	IncrementProgramCounter(increment int)
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

type Chip8Registers struct {
}

func (r *Chip8Registers) SetProgramCounter(address uint16) {
}

func (r *Chip8Registers) IncrementProgramCounter(increment int) {
}

func (r *Chip8Registers) PushProgramCounter() {
}

func (r *Chip8Registers) PopProgramCounter() uint16 {
	return 0
}

func (r *Chip8Registers) GetRegisterValue(register uint8) uint8 {
	return 0
}

func (r *Chip8Registers) SetRegisterValue(register, value uint8) {
}

func (r *Chip8Registers) SetRegisterI(value uint16) {
}

func (r *Chip8Registers) GetRegisterI() uint16 {
	return 0
}
