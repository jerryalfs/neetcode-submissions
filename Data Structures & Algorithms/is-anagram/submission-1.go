func isAnagram(s string, t string) bool {
	// check the length same or not, if not it's not anagram
	if len(s) != len(t) {
		return false
	}
	// alphabet in golang basically is number
    // a is an anchor, all alphabet will subtract with a
    // value of subtraction will put into slot base on alphabet index
    // if value subtraction s and t getting combined not equal to zero it will lead not anagram
	var count [26]int
	for i:=0; i < len(s); i++ {
		count[s[i] - 'a']++
		count[t[i] - 'a']--
	}
	for _,c := range count {
		if c != 0 {
			return false
		}
	}
	return true
}
