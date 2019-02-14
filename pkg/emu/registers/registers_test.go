package registers

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestRegisters(t *testing.T) {
	t.Run("Can set program counter to value", func(t *testing.T) {
		registers := NewRegisters()

		value := uint16(2345)
		registers.SetProgramCounter(value)

		registers.PushProgramCounter()
		pc := registers.PopProgramCounter()
		if pc != value {
			t.Error("Expected", value, "got", pc)
		}
	})

	t.Run("Can push 16 program counters on to stack", func(t *testing.T) {
		registers := NewRegisters()

		values := make([]uint16, 16)
		for i := 0; i <= 15; i++ {
			value := uint16(rand.Intn(4096))
			values[i] = value
			registers.SetProgramCounter(value)
			registers.PushProgramCounter()
		}

		for i := 15; i >= 0; i-- {
			value := registers.PopProgramCounter()
			if value != values[i] {
				t.Error("Expected", values[i], "got", value)
			}
		}
	})

	t.Run("Can increment program counter by increment times 2", func(t *testing.T) {
		registers := NewRegisters()
		value := uint16(2345)
		registers.SetProgramCounter(value)

		registers.IncrementProgramCounter(3)

		registers.PushProgramCounter()
		pc := registers.PopProgramCounter()
		if pc != value+6 {
			t.Error("Expected", value, "got", pc)
		}
	})

	for i := 0; i < 16; i++ {
		t.Run(fmt.Sprintf("Can get and set register values. (register %d)", i), func(t *testing.T) {
			registers := NewRegisters()
			value := uint8(rand.Intn(256))

			registers.SetRegisterValue(uint8(i), value)

			actual := registers.GetRegisterValue(uint8(i))
			if value != actual {
				t.Error("Expected", value, "got", actual)
			}
		})
	}

	t.Run("Can get and set register I", func(t *testing.T) {
		registers := NewRegisters()
		value := uint16(rand.Intn(4096))

		registers.SetRegisterI(value)

		actual := registers.GetRegisterI()
		if value != actual {
			t.Error("Expected", value, "got", actual)
		}
	})

	t.Run("Can get and set delay timer", func(t *testing.T) {
		registers := NewRegisters()
		value := uint8(rand.Intn(4096))

		registers.SetDelayTimer(value)

		actual := registers.GetDelayTimer()
		if value != actual {
			t.Error("Expected", value, "got", actual)
		}
	})

	t.Run("Can get and set sound timer", func(t *testing.T) {
		registers := NewRegisters()
		value := uint8(rand.Intn(4096))

		registers.SetSoundTimer(value)

		actual := registers.GetSoundTimer()
		if value != actual {
			t.Error("Expected", value, "got", actual)
		}
	})
}
