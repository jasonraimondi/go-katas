package jason_kata

func LengthOfLongestCommonPrefix(strs []string) string {
	longestPrefix := ""

	for k, str := range strs {
		newPrefix := ""

		if k == 0 {
			longestPrefix = str
			continue
		}

		strLen := len(str)

		for ik, c := range longestPrefix {
			if ik < strLen && string(c) == string(str[ik]) {
				newPrefix += string(c)
			}
			longestPrefix = newPrefix
			if longestPrefix == "" {
				break
			}
		}
	}

	return longestPrefix
}
