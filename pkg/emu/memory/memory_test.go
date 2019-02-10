package memory

import (
	"math/rand"
	"testing"
)

func TestMemory(t *testing.T) {
	t.Run("Can read and write memory locations", func(t *testing.T) {
		memory := NewMemory()
		addr := uint16(rand.Intn(4096))
		value := uint8(rand.Intn(256))

		memory.WriteValue(addr, value)

		actual := memory.ReadValue(addr)
		if actual != value {
			t.Error("Expected", value, "got", actual)
		}
	})

	t.Run("Can read n bytes of memory", func(t *testing.T) {
		memory := NewMemory()
		start := uint16(50)
		values := make([]uint8, 16)
		for i := 0; i < len(values); i++ {
			value := uint8(rand.Intn(256))
			values[i] = value
			memory.WriteValue(start+uint16(i), value)
		}

		data := memory.ReadNBytes(start, 16)
		if len(data) != len(values) {
			t.Error("Lengths not equal. Expected", len(values), "got", len(data))
		}
		for i := 0; i < len(values); i++ {
			if data[i] != values[i] {
				t.Error("Expected", values[i], "got", data[i])
			}
		}
	})

	t.Run("Can get hex digit addresses", func(t *testing.T) {
		memory := NewMemory()
		for i := 0; i < 16; i++ {
			addr := memory.GetHexDigitAddress(uint8(i))

			expected := uint16(i * 5)
			if addr != expected {
				t.Error("Expected", expected, "got", addr)
			}
		}
	})
}
