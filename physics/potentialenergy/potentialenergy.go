// potentialenergy.go
// description: In physics, potential energy is the energy held by an object because of its position relative to other objects, stresses within itself, its electric charge, or other factors. [Potential Energy](https://en.wikipedia.org/wiki/Potential_energy)
// [hypntc](https://github.com/hypntc)

package potentialenergy

import (
	"errors"
)

const (
	// gravitational acceleration
	g = 9.80665
)

var (
	ErrNegativeMass = errors.New("the mass of an object cannot be negative")
	ErrNegativeHeight = errors.New("the height of an object cannot be negative")
)

// calculate the potential energy of an object
func PotentialEnergy(mass, height float64) (float64, error) {
	// return an error if the mass is negative
	if mass < 0 {
		return 0, ErrNegativeMass
	}

	// return an error if the height is negative
	if height < 0 {
		return 0, ErrNegativeHeight
	}

	return mass * g * height, nil
}
