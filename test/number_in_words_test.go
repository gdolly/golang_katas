package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMany(t *testing.T) {
	tests := []struct {
		number        int
		expectedWords string
	}{
		{1, "one"},
		{2, "two"},
		{3, "three"},
		{4, "four"},
		{5, "five"},
		{6, "six"},
		{7, "seven"},
		{8, "eight"},
		{9, "nine"},
		{10, "ten"},
	}

	for _, testCase := range tests {
		t.Run(testCase.expectedWords, func(t *testing.T) {
			actual := InWords(testCase.number)
			if actual != testCase.expectedWords {
				t.Errorf("expected %s, got %s", testCase.expectedWords, actual)
			}
		})
	}
}

func TestShouldReturnZeroFor0(t *testing.T) {
	assert.Equal(t, "zero", InWords(0))
}

func TestShouldReturnOneToTenFor1ToTen(t *testing.T) {
	assert.Equal(t, "zero", InWords(0))
}

func InWords(number int) string {
	conversion := map[int]string{
		0:  "zero",
		1:  "one",
		2:  "two",
		3:  "three",
		4:  "four",
		5:  "five",
		6:  "six",
		7:  "seven",
		8:  "eight",
		9:  "nine",
		10: "ten",
	}
	return conversion[number]
}
