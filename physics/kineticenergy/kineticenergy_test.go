package kineticenergy

import (
	"testing"
)

func TestKineticEnergy(t *testing.T) {
	tests := getTestCases()
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			actual, err := KineticEnergy(test.mass, test.velocity)
			if err != test.err {
				t.Errorf("description:%v KineticEnergy() = %v, expected err: %v", test.description, err, test.err)
			}
			
			if actual != test.expected {
				t.Errorf("KineticEnergy() = %v, expected %v", actual, test.expected)
			}
		})
	}
}

func getTestCases() []struct {
	description string
	mass float64
	velocity float64
	expected float64
	err error
} {
	testCases := []struct {
		description string
		mass float64
		velocity float64
		expected float64
		err error
	}{
		{
			description: "success",
			mass: 10,
			velocity: 10,
			expected: 500,
		},
		{
			description: "zero mass",
			mass: 0,
			velocity: 10,
			expected: 0,
		},
		{
			description: "zero velocity",
			mass: 10,
			velocity: 0,
			expected: 0,
		},
		{
			description: "negative velocity",
			mass: 10,
			velocity: -20,
			expected: 2000,
		},
		{
			description: "zero mass and velocity",
			mass: 0,
			velocity: 0,
			expected: 0,
		},
		{
			description: "Negative mass",
			mass: -5,
			velocity: 10,
			expected: 0,
			err: ErrNegativeMass,
		},
	}

	return testCases
}