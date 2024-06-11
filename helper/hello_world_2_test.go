package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//subtest

func TestSubTest(t *testing.T) {
	t.Run("Test pertama", func(t *testing.T) {
		result := HelloWorld("Kerel")
		expected := "Hello, Kerel"
		if result != expected {
			t.Errorf("Result '%s' not match with expected '%s'", result, expected)
		}
	})

	t.Run("Test kedua", func(t *testing.T) {
		result := HelloWorld("Budi")
		expected := "Hello, Budi"
		assert.Equal(t, expected, result, "return not right")
	})
}

// konsep table test untuk menghindari perulangan kode
func TestHelloWorldTableTest(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Test pertama", "Kerel", "Hello, Kerel"},
		{"Test kedua", "Budi", "Hello, Budi"},
		{"Test ketiga", "", "Hello, "},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.input)
			if result != test.expected {
				t.Errorf("Result '%s' not match with expected '%s'", result, test.expected)
			}
		})
	}

}
