type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(strconv.Itoa(len(str)))
		sb.WriteByte('|')
		sb.WriteString(str)
	}
	return sb.String()
}

func (s *Solution) Decode(encoded string) []string {
	var result []string
	for i := 0; i < len(encoded); {
		j := i
		for encoded[j] != '|' {
			j ++
		}
		length, _ := strconv.Atoi(encoded[i:j])
		start := j + 1
		result = append(result, encoded[start:start+length])
		i = start + length
	}
	return result
}