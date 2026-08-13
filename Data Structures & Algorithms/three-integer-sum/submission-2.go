func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    lenN := len(nums)
    result := make([][]int,0)
    // why subtract with 2 directly, because the minimum result
    // i,j,k it needs at least 2 member
    for i:=0; i < lenN - 2; i++ {
        // skip duplicate if current nums i is same with previous nums i
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        // move the l and r as the rest member going to fulfill j,k
        l := i+1
        r := lenN - 1
        for l < r {
            // sum the entire member to check 
            // already fulfill the target or not
            sum := nums[i] + nums[l] + nums[r]
            switch {
                case sum == 0:
                    // append result because already reach target
                    result = append(result,[]int{nums[i],nums[l],nums[r]})
                    // move the l and r because already got a target
                    l ++
                    r --
                    // check is it current nums l is same with previous nums l
                    for l < r && nums[l] == nums[l-1] {
                        l ++
                    }
                    // check is it current nums r is same with previous nums r
                    for l < r && nums[r] == nums[r+1] {
                        r--
                    }
                case sum < 0 :
                    l ++
                default :
                    r --        
            }
        }
    }
    return result
}
