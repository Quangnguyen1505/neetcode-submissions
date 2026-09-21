import "slices"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false 
	}

	s_rune := []rune(s) // O(n) space 
	slices.Sort(s_rune) // O(nlogn) time

	t_rune := []rune(t) // O(1) space 
	slices.Sort(t_rune) // O(n)time

	return slices.Equal(s_rune, t_rune)
}

// TIME COMPLEXITY: O(nlogn) + O(nlogn) = O(nlogn) --- SPACE COMPLEXITY: O(n) + O(n) = O(n)