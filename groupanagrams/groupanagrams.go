package groupanagrams

import "sort"

func GroupAnagrams(strs []string) [][]string {
	aMap := make(map[string][]string) // sorted word -> [original words]    

	for _, word := range strs {
		chars := []byte(word)
		sort.Slice(chars, func(i, j int) bool {
			return chars[i] < chars[j]
		})
		sortedstr := string(chars)

		aMap[sortedstr] = append(aMap[sortedstr], word)
	}

	result := make([][]string, 0, len(aMap))
	for _, group := range aMap {
		result = append(result, group)
	}

	return result
}