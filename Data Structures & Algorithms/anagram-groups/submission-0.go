import "slices"

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 { // O(1) time
		return [][]string{}
	}

	if len(strs) == 1 { // O(1) time
		return [][]string{
			{strs[0]},
		}
	}

	anagramsMap := make(map[string][]string) // O(n) space
	for _, str := range strs { //O(n) time
		chars := []rune(str)
		slices.Sort(chars) // O(nlogn) time

		key := string(chars)
		anagramsMap[key] = append(anagramsMap[key], str)
	}

	var resultAnagrams [][]string // O(n) space
	for _, value := range anagramsMap { // O(n) time
		resultAnagrams = append(resultAnagrams, value)
	}
	return resultAnagrams
}

// TIME COMPLEXITY: O(1) + O(1) + O(n) + O(nlogn) + O(n) = O(nlogn) --- SPACE COMPLEXITY: O(n) 