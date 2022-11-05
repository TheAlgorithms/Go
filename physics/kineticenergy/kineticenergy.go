// kineticenergy.go
// description: Kinetic Energy
// details:
// In physics, the kinetic energy of an object is the energy that it possesses due to its motion. [Kinetic Energy](https://en.wikipedia.org/wiki/Kinetic_energy) is defined as:
// author(s) [hypntc](https://github.com/hypntc)
// see kineticenergy_test.go

package kineticenergy

import (
	"errors"
	"math"
)

func KineticEnergy(mass, velocity float64) (float64, error) {
	if mass < 0 {
		return 0, errors.New("the mass of an object cannot be less than zero")
	}

	return 0.5 * mass * math.Abs(velocity) * math.Abs(velocity), nil
}