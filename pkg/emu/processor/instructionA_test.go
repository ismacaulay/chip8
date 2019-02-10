package processor

import (
	"testing"

	"github.com/golang/mock/gomock"

	mock_registers "github.com/ismacaulay/chip8/pkg/emu/registers/mock"
)

func TestInstructionA(t *testing.T) {
	ctrl := gomock.NewController(t)
	registers := mock_registers.NewMockReaderWriter(ctrl)
	instruction := newInstructionA(registers)

	t.Run("[Annn] Set I = nnn", func(t *testing.T) {
		value := uint16(42)
		registers.EXPECT().SetRegisterI(value)
		registers.EXPECT().IncrementProgramCounter(uint16(1))

		opcode := uint16(0x000) | value

		instruction.execute(opcode)
	})

	ctrl.Finish()
}
