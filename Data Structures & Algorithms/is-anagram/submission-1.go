func isAnagram(s string, t string) bool {
	if len(s) != len(t) { // O(1) time
		return false
	}

	countMapS, countMapT := make(map[rune]int), make(map[rune]int) // O(n) space 
	for inx, value := range s { // O(n) time
		countMapS[value]++
		countMapT[rune(t[inx])]++
	}

	for inx, value := range countMapS { // O(n) time
		if countMapT[inx] != value { 
			return false
		}
	}
	return true
}

// TIME COMPLEXITY: O(1) + O(n) + O(n) = O(n) --- SPACE COMPLEXITY: O(n)
