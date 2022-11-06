// potentialenergy.go
// description: In physics, potential energy is the energy held by an object because of its position relative to other objects, stresses within itself, its electric charge, or other factors. [Potential Energy](https://en.wikipedia.org/wiki/Potential_energy)
// [hypntc](https://github.com/hypntc)

package potentialenergy

import (
	"errors"
)

const (
	g = 9.80665
)

func PotentialEnergy(mass, height float64) (float64, error) {
	if mass < 0 {
		return 0, errors.New("the mass of an object cannot be negative")
	}

	if height < 0 {
		return 0, errors.New("the height of an object cannot be negative")
	}

	return mass * g * height, nil
}