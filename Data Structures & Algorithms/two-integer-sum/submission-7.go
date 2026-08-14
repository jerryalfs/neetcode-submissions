func twoSum(nums []int, target int) []int {
	mapN := make(map[int]int)
	res := make([]int,0)
	pairs := 0
	for i:=0; i < len(nums); i++ {
		pairs = target - nums[i]
		if v,ok := mapN[pairs]; ok {
			res = append(res,i,v)
			sort.Ints(res)
			return res
		}
		mapN[nums[i]] = i
	}
	return res
}