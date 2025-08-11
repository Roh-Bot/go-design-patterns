package medium

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 1 {
		return [][]string{{strs[0]}}
	}
	var result [][]string

	m := make(map[[26]int]int)
	for _, str := range strs {
		var stats [26]int

		for _, v := range str {
			stats[v-'a']++
		}

		if index, ok := m[stats]; ok {
			result[index] = append(result[index], str)
			continue
		}
		result = append(result, []string{str})
		m[stats] = len(result) - 1
	}
	return result
}
