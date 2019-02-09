package processor

import (
	"testing"

	"github.com/golang/mock/gomock"
	mock_registers "github.com/ismacaulay/chip8/pkg/emu/registers/mock"
)

func TestInstruction1(t *testing.T) {
	ctrl := gomock.NewController(t)
	registers := mock_registers.NewMockRegisters(ctrl)
	instruction := newInstruction1(registers)

	t.Run("[1nnn] Jump to location nnn", func(t *testing.T) {
		registers.EXPECT().SetProgramCounter(uint16(0x0123))

		opcode := uint16(0x123)

		instruction.execute(opcode)
	})

	ctrl.Finish()
}
