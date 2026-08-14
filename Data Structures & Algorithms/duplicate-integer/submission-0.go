func hasDuplicate(nums []int) bool {
	mapN := make(map[int]struct{})
	for i:=0; i < len(nums); i ++ {
		if _,ok := mapN[nums[i]]; ok {
			return true
		}
		mapN[nums[i]] = struct{}{}
	}
	return false
}
