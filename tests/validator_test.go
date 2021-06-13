package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidator(t *testing.T) {
	testSuite := []struct {
		name          string
		expectedError bool
		input         interface{}
	}{
		{
			"Less than 3",
			true,
			"ab",
		},
		{
			"More than 3",
			false,
			"abcd",
		},
		{
			"Not a string",
			true,
			10,
		},
		{
			"Is a string",
			false,
			"simple string",
		},
	}

	for _, ts := range testSuite {
		t.Run(ts.name, func(t *testing.T) {
			err := ValidateString(ts.input)
			if ts.expectedError {

				assert.NotNil(t, err)
				return
			}

			assert.Nil(t, err)
		})
	}
}

func TestSum(t *testing.T) {
	assert.Equal(t, 3, sum(1, 2))

	assert.Equal(t, 5, sum(1, 2))

	assert.Equal(t, 10, sum(1, 2))

	assert.Equal(t, 13, sum(1, 2))
}
