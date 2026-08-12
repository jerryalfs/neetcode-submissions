func twoSum(numbers []int, target int) []int {
    result := make([]int,0)
    mapNumbers := make(map[int]int)
    partner := 0
    for i:=0; i < len(numbers);i++ {
        partner = target - numbers[i]
        if m,ok := mapNumbers[partner];ok {
            result = append(result,i+1)
            result = append(result,m+1)
            sort.Ints(result)
            break
        }
        mapNumbers[numbers[i]] = i
    }
    return result
}
