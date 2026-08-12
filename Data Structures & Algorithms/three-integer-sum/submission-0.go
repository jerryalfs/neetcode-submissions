func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    lenN := len(nums)
    result := make([][]int,0)
    // subtract len with 2 directly, because we need 2 other item
    // to conclude result, minimum 3 item i,j,k
    for i:=0; i < lenN - 2; i++ {
        // continue the process if duplicate 
        // in-case same with previous index
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        // move l and r to find other 2 member, i will stay there
        l := i+1
        r := lenN - 1
        for l < r {
            // summarize i,l,r to check zero or not
            sum := nums[i] + nums[l] + nums[r]
            switch {
                case sum == 0 :
                    // append result because summary 0
                    result = append(result,[]int{nums[i],nums[l],nums[r]})
                    // move l and r
                    l++
                    r--
                    // check if l duplicate with previous index
                    for l < r && nums[l] == nums[l-1] {
                        l++
                    }
                    for l < r && nums[r] == nums[r+1] {
                        r--
                    }
                // move the l because less than target    
                case sum < 0 :
                    l++ 
                // in default would be the one move    
                default :
                    r--       
            }
        }
    }
    return result
}
