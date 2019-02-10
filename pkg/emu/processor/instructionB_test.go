package processor

import (
	"testing"

	"github.com/golang/mock/gomock"

	mock_registers "github.com/ismacaulay/chip8/pkg/emu/registers/mock"
)

func TestInstructionB(t *testing.T) {
	ctrl := gomock.NewController(t)
	registers := mock_registers.NewMockRegisterReaderWriter(ctrl)
	instruction := newInstructionB(registers)

	t.Run("[Bnnn] Jump to location nnn + V0", func(t *testing.T) {
		v0Value := uint8(100)
		value := uint16(42)
		registers.EXPECT().GetRegisterValue(uint8(0x0)).Return(v0Value)
		registers.EXPECT().SetProgramCounter(value + uint16(v0Value))

		opcode := uint16(0x000) | value

		instruction.execute(opcode)
	})

	ctrl.Finish()
}
