package physics_test

import (
	"errors"
	"testing"

	"github.com/TheAlgorithms/Go/physics/kineticenergy"
)

func TestKineticEnergy(t *testing.T) {
	tests := getTests()
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			if res := KineticEnergy(test.mass, test.velocity); res != test.expected {
				t.Errorf("KineticEnergy() = %v, expected %v", res, test.expected)
			}
		})
	}
}

func getTestCases() []struct {
	description string
	mass float64
	velocity float64
	expected float64
} {
	testCases := []struct {
		description string
		mass float64
		velocity float64
		expected float64
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
			expected: errors.New("the mass of an object cannot be less than zero"),
		},
	}

	return testCases
}