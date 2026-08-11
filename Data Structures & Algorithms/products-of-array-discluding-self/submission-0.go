func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	// Pass 1: result[i] = product of everything to the LEFT of i
	prefix := 1
	for i := 0; i < n; i++ {
		result[i] = prefix
		prefix *= nums[i]
	}

	// Pass 2: multiply in the product of everything to the RIGHT of i
	suffix := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}

	return result
}