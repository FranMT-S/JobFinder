package helpers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStartsWithRegionalCharacters(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "🇺🇸 must be true",
			input:    "🇺🇸 United States",
			expected: true,
		},
		{
			name:     "United States must be false",
			input:    "United States",
			expected: false,
		},
		{
			name:     "🇧🇷 must be true",
			input:    "🇧🇷 Brazil",
			expected: true,
		},
		{
			name:     "🇪🇸 must be true",
			input:    "🇪🇸 Spain ",
			expected: true,
		},
	}

	fmt.Println(len("🇪"))
	fmt.Println(len("l"))
	fmt.Println(len("a"))

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := StartsWithRegionalCharacters(test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}
