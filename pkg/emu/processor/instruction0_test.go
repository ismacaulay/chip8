package processor

import (
	"testing"

	"github.com/golang/mock/gomock"

	mock_display "github.com/ismacaulay/chip8/pkg/emu/display/mock"
	mock_registers "github.com/ismacaulay/chip8/pkg/emu/registers/mock"
)

func TestInstruction0(t *testing.T) {
	ctrl := gomock.NewController(t)
	display := mock_display.NewMockDisplay(ctrl)
	registers := mock_registers.NewMockRegisters(ctrl)
	instruction := newInstruction0(display, registers)

	t.Run("[00E0] Clear the display", func(t *testing.T) {
		display.EXPECT().Clear()
		registers.EXPECT().IncrementProgramCounter(1)

		opcode := uint16(0x0E0)

		instruction.execute(opcode)
	})

	t.Run("[00EE] Return from a subroutine", func(t *testing.T) {
		address := uint16(0x1234)
		registers.EXPECT().PopProgramCounter().Return(address)
		registers.EXPECT().SetProgramCounter(address)

		opcode := uint16(0x0EE)

		instruction.execute(opcode)
	})

	t.Run("[0---] Igore all other opcodes", func(t *testing.T) {
		registers.EXPECT().IncrementProgramCounter(1)

		opcode := uint16(0x0123)

		instruction.execute(opcode)
	})

	ctrl.Finish()
}
