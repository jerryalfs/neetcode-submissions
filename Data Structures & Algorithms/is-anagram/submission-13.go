func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    // why 26 because alphabet in total 26
    var count [26]int
    for i:=0; i<len(s); i++ {
        // alphabet in golang basically is number
        // a is an anchor, all alphabet will subtract with a
        // value of subtraction will put into slot base on alphabet index
        // if value subtraction s and t getting combined not equal to zero it will lead not anagram
        count[s[i] - 'a'] ++
        count[t[i] - 'a'] --
    }
    for _,c := range count {
	    if c != 0 {
			return false
		}
	}
	return true
}