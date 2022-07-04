package math
import "fmt"

func decimal_to_binary(num int) {
	if num >= 1 {
		decimal_to_binary(num / 2)
	}
	fmt.Print(num % 2 , " ")
}


