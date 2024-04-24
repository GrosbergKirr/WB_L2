package num1

import (
	"fmt"
	"testing"
)

func TestUnpackString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		err      bool
	}{
		{"a4bc2d5e", "aaaabccddddde", false},
		{"abcd", "abcd", false},
		{"45", "", true},
		{"", "", false},
		{"5к4bв2ш5e", "кккккbccddddde", false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Input: %s", test.input), func(t *testing.T) {
			result, err := UnpackString(test.input)
			if test.err && err == nil {
				t.Errorf("Ожидалась ошибка, но получено: %s", result)
			} else if !test.err && err != nil {
				t.Errorf("Ожидался результат, но получена ошибка: %v", err)
			} else if result != test.expected {
				t.Errorf("Ожидалось '%s', но получено '%s'", test.expected, result)
			}
		})
	}
}
