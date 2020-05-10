package iterator

import (
	"testing"
)

func TestIterator(t *testing.T) {
	iter := NewIterator("Z", 29, 5)
	values := []string{
		"Z5",
		"AA5",
		"AB5",
		"AC5",
		"AD5",
		"AE5",
		"AF5",
		"AG5",
		"AH5",
		"AI5",
		"AJ5",
		"AK5",
		"AL5",
		"AM5",
		"AN5",
		"AO5",
		"AP5",
		"AQ5",
		"AR5",
		"AS5",
		"AT5",
		"AU5",
		"AV5",
		"AW5",
		"AX5",
		"AY5",
		"AZ5",
		"BA5",
		"BB5",
	}

	var i int

	for iter.Next() {
		index, address := iter.GetAddress()
		if i != index {
			t.Error("Expected ", index, ", got ", i)
		}

		if values[i] != address {
			t.Error("Expected ", values[i], ", got ", address)
		}

		i++
	}

	if i != len(values) {
		t.Error("Expected ", len(values), ", got ", i)
	}
}
