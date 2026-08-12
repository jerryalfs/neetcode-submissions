func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    lenN := len(nums)
    res := make([][]int,0)
    // subtract with 2 directly because need to find other
    // 2 team member, to conclude result need i,j,k
    for i:=0; i < lenN -2; i++ {
        // check duplication for current i with previous index
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        // move l and r to find other member, let i stay there
        l := i + 1
        r := lenN - 1
        for l < r {
            // summarize to check already reach target or not
            sum := nums[i] + nums[l] + nums[r]
            switch {
                case sum == 0 :
                    // append res
                    res = append(res,[]int{nums[i],nums[l],nums[r]})
                    // move l and r
                    l++
                    r--
                    // check duplication l with previous index
                    // if yes move l
                    for l < r && nums[l] == nums[l-1] {
                        l++
                    }
                    // check duplication r with previous index
                    // if yes move r
                    for l < r && nums[r] == nums[r+1] {
                        r--
                    }
                // if sum lower than target, move l    
                case sum < 0 :
                    l++
                default :
                    r--        
            }
        }
    }
    return res
}
