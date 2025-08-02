package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleanEmojiFromString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "test clean emoji",
			input:    "Hello 👋",
			expected: "Hello",
		},
		{
			name:     "Test clean space and emoji",
			input:    "  Hello 🌍",
			expected: "Hello",
		},
		{
			name:     "test clean money with emoji",
			input:    " 💰 $10k - $30k ",
			expected: "$10k - $30k",
		},
		{
			name:     "Country test",
			input:    " 🇺🇸 United States ",
			expected: "United States",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := CleanEmojiAndTrimString(test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestCleanOnlyNumbersFromString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "test clean only numbers",
			input:    "Hello 123",
			expected: "123",
		},
		{
			name:     "test clean only numbers with $",
			input:    "$10k",
			expected: "10",
		},
		{
			name:     "test clean only numbers with space",
			input:    "$30k ",
			expected: "30",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := CleanOnlyNumbersFromString(test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}
