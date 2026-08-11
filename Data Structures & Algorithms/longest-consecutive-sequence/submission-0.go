func longestConsecutive(nums []int) int {
    mapNums := make(map[int]bool)
    longest := 0
    for _,n := range nums {
        mapNums[n] = true
    }
    for n := range mapNums {
        length := 1
        if mapNums[n-1] {
            continue
        }
        for mapNums[n+length] {
            length ++
        }
        if length > longest {
            longest = length
        }
    }
    return longest
}
