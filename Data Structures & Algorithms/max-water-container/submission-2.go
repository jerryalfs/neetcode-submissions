func maxArea(heights []int) int {
	l := 0
	r := len(heights) - 1
	// maximum amount water can store
	maxArea := 0
	for l < r {
		// identify distance index as width
		w := r-l
		// find height by checking which min value 
		// between heights l and r, why find shortest
		// because water wills spoil on shorter wall from left
		h := min(heights[l],heights[r])
		// area from width x height
		a := w*h
		if a > maxArea {
			maxArea = a
		}
		// move index l because still shortest than index r
		if heights[l] < heights[r] {
			l++
		} else {
			// move index r to check next index since l index taller than r
			r--
		}
	}
	return maxArea
}
