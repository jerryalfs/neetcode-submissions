func hasDuplicate(nums []int) bool {
	mapNums := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		if _, ok := mapNums[nums[i]]; ok {
			return true
		}
		mapNums[nums[i]] = struct{}{}
	}
	return false
}