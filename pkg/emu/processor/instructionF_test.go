package processor

import (
	"math/rand"
	"testing"

	"github.com/golang/mock/gomock"

	mock_keyboard "github.com/ismacaulay/chip8/pkg/emu/keyboard/mock"
	mock_memory "github.com/ismacaulay/chip8/pkg/emu/memory/mock"
	mock_registers "github.com/ismacaulay/chip8/pkg/emu/registers/mock"
)

func TestInstructionF(t *testing.T) {
	ctrl := gomock.NewController(t)
	keyboard := mock_keyboard.NewMockKeyboard(ctrl)
	memory := mock_memory.NewMockReaderWriter(ctrl)
	registers := mock_registers.NewMockRegisterReaderWriter(ctrl)
	instruction := newInstructionF(keyboard, memory, registers)

	t.Run("[Fx07] Load delay timer into Vx", func(t *testing.T) {
		dt := uint8(234)
		registers.EXPECT().GetDelayTimer().Return(dt)
		registers.EXPECT().SetRegisterValue(uint8(0x0C), dt)
		registers.EXPECT().IncrementProgramCounter(uint16(1))

		opcode := uint16(0xC07)

		instruction.execute(opcode)
	})

	t.Run("[Fx0A] Wait for keypress and store key value in Vx", func(t *testing.T) {
		key := uint8(0x5)
		keyboard.EXPECT().GetKeyPress().Return(key)
		registers.EXPECT().SetRegisterValue(uint8(0x0E), key)
		registers.EXPECT().IncrementProgramCounter(uint16(1))

		opcode := uint16(0xE0A)

		instruction.execute(opcode)
	})

	t.Run("[Fx15] Set delay timer to Vx", func(t *testing.T) {
		value := uint8(123)
		registers.EXPECT().GetRegisterValue(uint8(0x0A)).Return(value)
		registers.EXPECT().SetDelayTimer(value)
		registers.EXPECT().IncrementProgramCounter(uint16(1))

		opcode := uint16(0xA15)

		instruction.execute(opcode)
	})

	t.Run("[Fx18] Set sound timer to Vx", func(t *testing.T) {
		value := uint8(123)
		registers.EXPECT().GetRegisterValue(uint8(0x0A)).Return(value)
		registers.EXPECT().SetSoundTimer(value)
		registers.EXPECT().IncrementProgramCounter(uint16(1))

		opcode := uint16(0xA18)

		instruction.execute(opcode)
	})

	t.Run("[Fx1E] Set I = I + Vx", func(t *testing.T) {
		vxValue := uint8(42)
		iValue := uint16(1024)
		registers.EXPECT().GetRegisterValue(uint8(0x07)).Return(vxValue)
		registers.EXPECT().GetRegisterI().Return(iValue)
		registers.EXPECT().SetRegisterI(iValue + uint16(vxValue))
		registers.EXPECT().IncrementProgramCounter(uint16(1))

		opcode := uint16(0x71E)

		instruction.execute(opcode)
	})

	t.Run("[Fx29] Set I location of sprite for digit Vx", func(t *testing.T) {
		vxValue := uint8(0x09)
		spriteAddr := uint16(0x0050)
		registers.EXPECT().GetRegisterValue(uint8(0x07)).Return(vxValue)
		memory.EXPECT().GetHexDigitAddress(vxValue).Return(spriteAddr)
		registers.EXPECT().SetRegisterI(spriteAddr)
		registers.EXPECT().IncrementProgramCounter(uint16(1))

		opcode := uint16(0x729)

		instruction.execute(opcode)
	})

	t.Run("[Fx33] Store BCD representation of Vx in memory I, I+1, and I+2", func(t *testing.T) {
		vxValue := uint8(39)
		iValue := uint16(2345)
		registers.EXPECT().GetRegisterValue(uint8(0x07)).Return(vxValue)
		registers.EXPECT().GetRegisterI().Return(iValue)
		memory.EXPECT().WriteValue(uint16(iValue), uint8(0))
		memory.EXPECT().WriteValue(uint16(iValue+1), uint8(3))
		memory.EXPECT().WriteValue(uint16(iValue+2), uint8(9))
		registers.EXPECT().IncrementProgramCounter(uint16(1))

		opcode := uint16(0x733)

		instruction.execute(opcode)
	})

	t.Run("[Fx55] Store registers V0 through Vx in memory starting at location I", func(t *testing.T) {
		rand.Seed(42)
		iValue := uint16(2048)
		registers.EXPECT().GetRegisterI().Return(iValue)
		for i := uint8(0); i <= 0xE; i++ {
			value := uint8(rand.Intn(256))
			registers.EXPECT().GetRegisterValue(i).Return(value)
			memory.EXPECT().WriteValue(iValue+uint16(i), value)
		}
		rand.Seed(42)
		registers.EXPECT().IncrementProgramCounter(uint16(1))

		opcode := uint16(0xE55)

		instruction.execute(opcode)
	})

	t.Run("[Fx65] Read registers V0 through Vx from memory starting at location I.", func(t *testing.T) {
		rand.Seed(42)
		iValue := uint16(2048)
		registers.EXPECT().GetRegisterI().Return(iValue)
		for i := uint8(0); i <= 0xE; i++ {
			value := uint8(rand.Intn(256))
			memory.EXPECT().ReadValue(iValue + uint16(i)).Return(value)
			registers.EXPECT().SetRegisterValue(i, value)
		}
		rand.Seed(42)
		registers.EXPECT().IncrementProgramCounter(uint16(1))

		opcode := uint16(0xE65)

		instruction.execute(opcode)
	})
	ctrl.Finish()
}
