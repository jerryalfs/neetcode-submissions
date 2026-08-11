func topKFrequent(nums []int, k int) []int {
	mapCounter := make(map[int]int)
	// count frequency
	for _, n := range nums {
		mapCounter[n]++
	}
	keys := make([]int, 0, len(mapCounter))
	for m := range mapCounter {
		keys = append(keys, m)
	}
	sort.Slice(keys, func(i, j int) bool {
		return mapCounter[keys[i]] > mapCounter[keys[j]]
	})
	
	return keys[:k]
}