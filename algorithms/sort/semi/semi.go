package semi

import "math"

func Sort(nums []int, low int) []int {
	out := make([]int, len(nums))

	for i := 0; i < len(out); i++ {
		out[i] = math.MinInt
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] >= low && nums[i] < low+len(nums) {
			out[nums[i]-low] = nums[i]
		}
	}

	return out
}
