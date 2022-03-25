package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Square(a, b int) int {
	return a * b
}

func TestSquere(t *testing.T) {
	asserts := assert.New(t)
	tests := []struct {
		title  string
		input  []int
		output int
	}{
		{
			title:  "2x3の面積は6になる",
			input:  []int{2, 3},
			output: 6,
		},
		{
			title:  "0x1の面積は0になる",
			input:  []int{0, 1},
			output: 0,
		},
	}

	for _, td := range tests {
		td := td
		t.Run("Square:"+td.title, func(t *testing.T) {
			result := Square(td.input[0], td.input[1])
			asserts.Equal(td.output, result)
		})
	}
}
