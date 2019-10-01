// an implementation of the longest increasing subsequence
// Runs in O(nlogn)

package longestincreasing

// import "fmt"

func longest_increasing(arr []int) []int {
	candidates := [][]int{arr[0:1]}

	for i := 1; i < len(arr); i++ {
		num := arr[i]

		back := candidates[len(candidates)-1]
		if num > back[len(back)-1] {
			temp := append(back, num)
			candidates = append(candidates, temp)
		} else if num < candidates[0][0] {
			candidates[0] = []int{num}
		} else {
			start := 0
			end := len(candidates) - 1

			for end > start+1 {
				mid := start + int((end-start)/2)
				curr := candidates[mid]
				if curr[len(curr)-1] < num {
					start = mid
				} else {
					end = mid
				}
			}

			temp := append(candidates[end-1], num)
			candidates[end] = temp
		}
	}
	return candidates[len(candidates)-1]
}

// func main() {
// 	t := []int{4, 6, -1, 4, 0, 12, 8}
// 	fmt.Println(longest_increasing(t))
// }
