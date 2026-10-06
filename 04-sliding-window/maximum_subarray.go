package slidingwindow

func MaxSubarraySum(nums []int, k int) int {
	if k <= 0 || k > len(nums) {
		return 0
	}

	sum := 0

	for i := range k {
		sum += nums[i]
	}

	sumMax := sum

	for i := k; i < len(nums); i++ {
		sum += nums[i]
		sum -= nums[i-k]

		if sum > sumMax {
			sumMax = sum
		}
	}

	return sumMax
}
