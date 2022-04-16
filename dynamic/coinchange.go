package dynamic
/* 
@params 
coins: combination of coin 
amount: target coins

@return
num of combination
 */


func coinChange(coins []int32, amount int32) int32 {
	combination := make([]int32, amount) 
	combination[0] = 1

	//Time Complexity O(mn)
	//Space Complexity O(n)
	//m is length coins
	//n is amount

	for _,c := range coins { // iterate combination of coin
		for i:= c; i < amount; i++{ 
			// find previous value
			combination[i] += combination[i - c]
		}
	}

	return combination[amount-1]
}