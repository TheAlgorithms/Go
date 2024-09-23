package dynamic

import "testing"

func TestHouseRobber(t *testing.T) {
	testCases := map[string]struct {
		nums []int
		want int
	}{
		"one house":             {[]int{5}, 5},
		"two houses":            {[]int{1, 2}, 2},
		"three houses":          {[]int{1, 2, 3}, 4},
		"four houses":           {[]int{2, 7, 9, 3}, 11},
		"five houses":           {[]int{2, 7, 9, 3, 1}, 12},
		"non-adjacent houses":   {[]int{1, 3, 1, 3, 100}, 103},
		"all houses same value": {[]int{10, 10, 10, 10}, 20},
		"large values":          {[]int{100, 1, 1, 100}, 200},
		"stress test":           {[]int{1, 2, 3, 1, 5, 1, 100, 1, 2}, 111},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			if got := HouseRobber(test.nums); got != test.want {
				t.Errorf("rob(%v) = %v, want %v", test.nums, got, test.want)
			}
		})
	}
}
