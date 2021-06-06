package montecarlo

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	INTERVAL = 1_00_000
)

func estimate_pi() {

	fmt.Printf("Estimating value if π with %d experiments\n", INTERVAL)

	rand.Seed(time.Now().UTC().UnixNano())
	var pi, randX, randY, disFromOrigin float64
	circlePoints, squarePoints := 0, 0

	for i := 0; i < INTERVAL; i++ {

		randX = float64(rand.Int()%(INTERVAL+1)) / float64(INTERVAL)
		randY = float64(rand.Int()%(INTERVAL+1)) / float64(INTERVAL)
		disFromOrigin = math.Sqrt(randX*randX + randY*randY)

		//check if distance from origin is less than or equal to 1, if true then poin lies inside the circle
		if disFromOrigin <= 1 {
			circlePoints++
		}
		squarePoints++

		fmt.Println(randX, randY, circlePoints, squarePoints, " - ", pi)

		pi = float64(4*circlePoints) / float64(squarePoints)
	}

	fmt.Printf("π := %f\n", pi)
}

// func main() {
// 	estimate_pi()
// }
