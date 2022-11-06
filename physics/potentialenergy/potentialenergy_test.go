package physics_test

import (
	"errors"
	"testing"

	"github.com/hypntc/Go/physics/potentialenergy"
)

func TestPotentialEnergy(t *testing.T) {
	tests := getTests()
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			if res := PotentialEnergy(test.mass, test.height); res != test.expected {
				t.Errorf("PotentialEnergy() = %v, expected %v", res, test.expected)
			}
		})
	}
}

func getTestCases() []struct {
	description string
	mass float64
	height float64
	expected float64
} {
	testCases := []struct {
		description string
		mass float64
		height float64
		expected float64
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
			expected: errors.New("the mass of an object cannot be negative"),
		},
		{
			description: "negative height",
			mass: 10,
			height: -20,
			expected: errors.New("the height of an object cannot be negative"),
		},
	}

	return testCases
}