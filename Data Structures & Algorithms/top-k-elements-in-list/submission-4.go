func topKFrequent(nums []int, k int) []int {
	// create map to store and counter of each value from nums
	mapCounter := make(map[int]int)
	for _,n := range nums {
		mapCounter[n]++
	}
	// create slice of int that store 
	// every keys with size len(mapCounter)
	keys := make([]int,0,len(mapCounter))
	for mC := range mapCounter {
		keys = append(keys,mC)
	}
	// sort mapCounter base on keys that we already stored
	// compare the value from each keys
	sort.Slice(keys,func(i,j int) bool {
		return mapCounter[keys[i]] > mapCounter[keys[j]]
	})
	// return the result from the keys until the target
	// because we only need the key
	return keys[:k]
}