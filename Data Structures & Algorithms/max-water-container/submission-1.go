func maxArea(heights []int) int {
	l := 0
	r := len(heights) - 1
	maxArea := 0
	for l < r {
		w := r-l
		h := min(heights[l],heights[r])
		area := w * h
		if area > maxArea {
			maxArea = area
		}
		if heights[l] < heights[r] {
			l ++
		} else {
			r --
		}
	}
	return maxArea
}
