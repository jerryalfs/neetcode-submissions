func twoSum(nums []int, target int) []int {
	var result []int
	mapPairs := make(map[int]int)
	for i := 0; i < len(nums); i++ {
        if len(result) >= 2 {
            return result
        }
		pairs := target - nums[i]
		if v, ok := mapPairs[pairs]; ok {
			result = append(result, v)
			result = append(result, i)
		}
		mapPairs[nums[i]] = i
	}
	return result
}