// kineticenergy.go
// description: In physics, the kinetic energy of an object is the energy that it possesses due to its motion. [Kinetic Energy](https://en.wikipedia.org/wiki/Kinetic_energy) is defined as:
// [hypntc](https://github.com/hypntc)

package kineticenergy

import (
	"errors"
	"math"
)

var (
	ErrNegativeMass = errors.New("the mass of an object cannot be negative")
)

// calculate the kinetic energy of an object
func KineticEnergy(mass, velocity float64) (float64, error) {
	// return an error if the mass is negative
	if mass < 0 {
		return 0, ErrNegativeMass
	}

	return 0.5 * mass * math.Abs(velocity) * math.Abs(velocity), nil
}
