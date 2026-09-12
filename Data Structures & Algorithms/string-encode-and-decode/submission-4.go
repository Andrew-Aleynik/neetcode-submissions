type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	encoded := ""
	for _, str := range strs {
		length := strconv.Itoa(len(str))
		encoded += length
		encoded += "#"
		encoded += str
	}
	return encoded
}

func (s *Solution) Decode(encoded string) []string {
	result := make([]string, 0)

	i := 0
	for i < len(encoded) {
		j := i
		for j < len(encoded) && encoded[j] != '#' {
			j++
		}

		length, _ := strconv.Atoi(encoded[i:j])
		start := j + 1

		result = append(result, encoded[start:start+length])
		i = start + length
	}
	return result
}