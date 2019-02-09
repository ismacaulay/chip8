package processor

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"

	mock_keyboard "github.com/ismacaulay/chip8/pkg/emu/keyboard/mock"
	mock_registers "github.com/ismacaulay/chip8/pkg/emu/registers/mock"
)

func TestInstructionE(t *testing.T) {
	ctrl := gomock.NewController(t)
	keyboard := mock_keyboard.NewMockKeyboard(ctrl)
	registers := mock_registers.NewMockRegisters(ctrl)
	instruction := newInstructionE(keyboard, registers)

	cases := []struct {
		isPressed         bool
		expectedIncrement int
	}{
		{true, 2},
		{false, 1},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("[Ex9E] Skip next instruction if key with the value of Vx is pressed. (pressed %t)", c.isPressed), func(t *testing.T) {
			key := uint8(0x04)

			registers.EXPECT().GetRegisterValue(uint8(0x3)).Return(key)
			keyboard.EXPECT().IsPressed(key).Return(c.isPressed)
			registers.EXPECT().IncrementProgramCounter(c.expectedIncrement)

			opcode := uint16(0x39E)

			instruction.execute(opcode)
		})
	}

	cases = []struct {
		isPressed         bool
		expectedIncrement int
	}{
		{false, 2},
		{true, 1},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("[ExA1] Skip next instruction if key with the value of Vx is not pressed. (pressed %t)", c.isPressed), func(t *testing.T) {
			key := uint8(0x04)

			registers.EXPECT().GetRegisterValue(uint8(0x3)).Return(key)
			keyboard.EXPECT().IsPressed(key).Return(c.isPressed)
			registers.EXPECT().IncrementProgramCounter(c.expectedIncrement)

			opcode := uint16(0x3A1)

			instruction.execute(opcode)
		})
	}

	ctrl.Finish()
}
