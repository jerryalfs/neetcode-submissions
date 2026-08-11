func isPalindrome(s string) bool {
    cs := alphaNumeric(s)
    l := 0
    r := len(cs) - 1
    for l < r {
        if cs[l] != cs[r] {
            return false
        }
        l ++
        r --
    }
    return true
}

func alphaNumeric(s string) string {
    return strings.Map(func(r rune) rune {
        if unicode.IsDigit(r) || unicode.IsLetter(r) {
            return unicode.ToLower(r)
        }
        return -1
    },s)
}
