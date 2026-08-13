func trap(height []int) int {
	lenH := len(height)
	res := 0
	// return in-case len was zero
	if lenH == 0 {
		return res
	}
	l := 0 // left pointer to check from left to right
	r := lenH - 1 // right pointer to check from right to left
	// predefine current maxLeft and maxRight 
	// to compare with next taller water level
	maxLeft := height[l] 
	maxRight := height[r]
	for l < r {
		// it means maxLeft was shorter wall for water trap
		if maxLeft < maxRight {
			// move index l to check next index was higher or not with current maxLeft
			l++
			// check which one taller
			maxLeft = max(maxLeft,height[l])
			// addition res with subtraction value from taller wall, 
			// if the one taller current l index, it will return 0
			res += maxLeft - height[l]
		} else {
			// check in-case maxLeft was taller than maxRight
			// move index r to compare with current maxRight
			r--
			maxRight = max(maxRight,height[r])
			// addition max res subtraction with height[r] 
			// to find how many water able to trap
			// if current height taller than current maxRight
			// it will return 0 because subtract with himself
			res += maxRight - height[r]
		}
	}
	return res
}