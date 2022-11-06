package potentialenergy

import (
	"testing"
)

func TestPotentialEnergy(t *testing.T) {
	tests := getTestCases()
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			actual, err := PotentialEnergy(test.mass, test.height)
			if err != test.err {
				t.Errorf("description:%v PotentialEnergy() = %v, expected err: %v", test.description, err, test.err)
			}
			
			if actual != test.expected {
				t.Errorf("PotentialEnergy() = %v, expected %v", actual, test.expected)
			}
		})
	}
}

func getTestCases() []struct {
	description string
	mass float64
	height float64
	expected float64
	err error
} {
	testCases := []struct {
		description string
		mass float64
		height float64
		expected float64
		err error
	}{
		{
			description: "success",
			mass: 10,
			height: 10,
			expected: 980.665,
		},
		{
			description: "zero mass",
			mass: 0,
			height: 10,
			expected: 0,
		},
		{
			description: "zero height",
			mass: 10,
			height: 0,
			expected: 0,
		},
		{
			description: "zero mass and height",
			mass: 0,
			height: 0,
			expected: 0,
		},
		{
			description: "Negative mass",
			mass: -5,
			height: 10,
			expected: 0,
			err: ErrNegativeMass,
		},
		{
			description: "negative height",
			mass: 10,
			height: -20,
			expected: 0,
			err: ErrNegativeHeight,
		},
	}

	return testCases
}
