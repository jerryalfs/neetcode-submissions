func twoSum(nums []int, target int) []int {
	lenN := len(nums)
	mapP := make(map[int]int)
	pairs := 0
	res := make([]int,0)
	for i:=0 ; i < lenN; i ++ {
		pairs = target - nums[i]
		if v,ok := mapP[pairs]; ok {
			res = append(res,i,v)
			sort.Ints(res)
			return res
		}
		mapP[nums[i]] = i
	}
	return res
}