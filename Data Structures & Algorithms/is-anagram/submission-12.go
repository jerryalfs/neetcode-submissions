func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
	runesS := []rune(s)
	runesT := []rune(t)
	sort.Slice(runesS, func(i, j int) bool { return runesS[i] < runesS[j] })
	sort.Slice(runesT, func(i, j int) bool { return runesT[i] < runesT[j] })
	sortedStringS := string(runesS)
	sortedStringT := string(runesT)
	if sortedStringS != sortedStringT {
		return false
	}
	return true
}