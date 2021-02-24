package recursion

import "fmt"

// https://en.wikipedia.org/wiki/Tower_of_Hanoi

func towerOfHanoi(n int, towerSource, towerDestination string, towerAuxiliary string) {
	if n == 0 {
		return
	}
	towerOfHanoi(n-1, towerSource, towerAuxiliary, towerDestination)
	fmt.Println("Moving Disc from", towerSource, "->", towerDestination)
	towerOfHanoi(n-1, towerAuxiliary, towerDestination, towerSource)
}

// func main() {
// 	//A --> Source
// 	//B --> Target
// 	//C --> towerAuxiliary
// 	towerOfHanoi(3, "TowerA", "TowerB", "TowerC")
// }
